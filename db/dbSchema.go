package db

import (
	"strings"
	"time"

	"github.com/ie310mu/ie310go/common/convert"
	"github.com/ie310mu/ie310go/common/obj"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//ConvertToTrMapper 转换成mapper对象（试验室分表处理） 只能调用一次
func (m *BaseMapper) ConvertToSuffixMapper(defaultDatabaseName, suffix string) {
	if strings.HasSuffix(m.TableName, "_"+suffix) {
		return
	}

	schemaKey := m.TableName
	m.TableName = "at_" + m.TableName + "_" + suffix //分表处理  加个固定前缀避免代码生成器多生成内容
	checkTableSchema(schemaKey, defaultDatabaseName, m.TableName)
}

var schemaVersionMap = map[string]int{} //每个表的最新版本，key是原始表名
//设置最新版本，上一个版本的检查，必须每次加1
func updateSchemaVersion(schemaKey string, version int) {
	lastVersion, exists := schemaVersionMap[schemaKey]
	if !exists {
		lastVersion = 0
	}
	if lastVersion+1 != version {
		panic("error : last version is " + convert.IntToStr(lastVersion) + "，now version is " + convert.IntToStr(version))
	}
	schemaVersionMap[schemaKey] = version
}

//设置每个版本的sql，会进行版本重复检查，尽量避免脚本设置错误
func SetSchema(schemaKey string, schemas map[int]string, version int, schema string) {
	if strings.Index(schema, "{TableName}") < 0 {
		panic("no string {TableName} in schema: " + schema)
	}

	_, exists := schemas[version]
	if exists {
		panic("version " + convert.IntToStr(version) + " already exists")
	}
	schemas[version] = schema
	updateSchemaVersion(schemaKey, version)
}

var schemaMap = map[string]map[int]string{} //每个表的架构历史，key是原始表名
//设置sqls，会进行主键重复检查，尽量避免脚本设置错误
func SetSchemas(schemaKey string, schemas map[int]string) {
	_, exists := schemaMap[schemaKey]
	if exists {
		panic("schema of " + schemaKey + " already exists")
	}
	schemaMap[schemaKey] = schemas
}

var schemaCheckedLastVersionMap = map[string]int{} //记录每个表最后检查过的版本，避免重复检查
func checkTableSchema(schemaKey string, databaseName string, tableName string) {
	//获取架构数据
	lastVersion, exists := schemaVersionMap[schemaKey]
	if !exists {
		panic("no lastVersion of " + schemaKey + " in schemaVersionMap")
	}
	//获取是否已检查过版本
	lastCheckedVersion, exists := schemaCheckedLastVersionMap[tableName]
	if exists && lastCheckedVersion >= lastVersion {
		return
	}
	schemas, exists := schemaMap[schemaKey]
	if !exists {
		panic("no schemas of " + schemaKey + " in schemaMap")
	}

	//DDL DCL语句，会自动commit
	//因此，下面处理时，需要将对tableschema_version的DML操作放到前面，schema sql放到后面
	//而且，每1个版本的shcema都要单独启动一个事务，且先DDL再DCL最后DML
	mp := GetTableschemaVersionMapper("")
	var flag = true
	for flag {
		mp.Tx2(func(tx *sqlx.Tx) {
			sql := "select version from " + databaseName + ".tableschema_version where tableName=?"
			nowVersion := mp.GetCountss(tx, sql, tableName) //没有的话返回0
			if nowVersion >= lastVersion {
				flag = false
				schemaCheckedLastVersionMap[tableName] = nowVersion
				return
			}

			nowVersion = nowVersion + 1
			sql2, exists := schemas[nowVersion]
			if !exists {
				panic(schemaKey + "的架构中没有版本号为" + convert.IntToStr(nowVersion) + "的数据")
			}

			//执行一次use database，在正常的数据库中操作
			mp.Updatess(tx, "USE "+databaseName+";")

			//先写入版本信息，避免DDL导致事务自动提交后万一出错版本信息无法回滚（即最后执行DDL）
			if nowVersion == 1 {
				sql3 := "insert into " + databaseName + ".tableschema_version(tableName,version,updateTime) values (?,?,?)"
				mp.Updatess(tx, sql3, tableName, nowVersion, time.Now())
			} else {
				sql3 := "update " + databaseName + ".tableschema_version set version=? where tableName=?"
				mp.Updatess(tx, sql3, nowVersion, tableName)
			}

			//执行sql
			sql2 = strings.Replace(sql2, "{DatabaseName}", databaseName, -1)
			sql2 = strings.Replace(sql2, "{TableName}", tableName, -1)
			mp.Updatess(tx, sql2)
		})
	}
}

//TableschemaVersion ..
type TableschemaVersion struct {
	ID         int       `json:"id"  db:"id"  `
	TableName  string    `json:"tableName"  db:"tableName"  `
	Version    int       `json:"version"  db:"version"  `
	UpdateTime time.Time `json:"updateTime"  db:"updateTime"  `
}

//TableschemaVersionMapper ..
type TableschemaVersionMapper struct {
	BaseMapper
}

//GetTableschemaVersionMapper 获取mapper对象
func GetTableschemaVersionMapper(db string) TableschemaVersionMapper {
	var mp TableschemaVersionMapper
	mp.TableName = "tableschema_version"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//将interface转换为*TableschemaVersion
func (m TableschemaVersionMapper) convertItem(v interface{}) *TableschemaVersion {
	if obj.InterfaceIsNil(v) {
		return nil
	}
	return v.(*TableschemaVersion)
}

//Get 根据id获取
func (m TableschemaVersionMapper) Get(tx *sqlx.Tx, id interface{}) *TableschemaVersion {
	sqlStr := "select * from " + m.TableName + " where id = ? "
	var args = []interface{}{id}

	item := &TableschemaVersion{}
	r := m.ScanItem(tx, item, sqlStr, args...)
	return m.convertItem(r)
}

//Exists 根据id判断数据是否存在
func (m TableschemaVersionMapper) Exists(tx *sqlx.Tx, id interface{}) bool {
	sqlStr := "select count(id) from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	count := m.ScanInt(tx, sqlStr, args...)
	return count > 0
}

//GetByCode 根据code获取，需要有code字段
func (m TableschemaVersionMapper) GetByCode(tx *sqlx.Tx, code interface{}) *TableschemaVersion {
	sqlStr := "select * from " + m.TableName + " where code = ? "
	var args = []interface{}{code}

	item := &TableschemaVersion{}
	r := m.ScanItem(tx, item, sqlStr, args...)
	return m.convertItem(r)
}

//Gets 获取多条数据、pageIndex从1开始、whereStr、orderStr可以为空，不要where、order by部分
func (m TableschemaVersionMapper) Gets(tx *sqlx.Tx, pageIndex int, rowsInPage int, whereStr string, orderStr string, args ...interface{}) []*TableschemaVersion {
	sqlStr := "select * from " + m.TableName
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}
	if orderStr != "" {
		sqlStr = sqlStr + " order by " + orderStr
	}

	return m.Getss(tx, pageIndex, rowsInPage, sqlStr, args...)
}

//Getss 获取多条数据、需要完整的sql语句、pageIndex从1开始
func (m TableschemaVersionMapper) Getss(tx *sqlx.Tx, pageIndex int, rowsInPage int, sqlStr string, args ...interface{}) []*TableschemaVersion {
	items := make([]*TableschemaVersion, 0)
	m.ScanItems(tx, pageIndex, rowsInPage, &items, sqlStr, args...)
	return items
}

//GetTops 获取前几条数据
func (m TableschemaVersionMapper) GetTops(tx *sqlx.Tx, top int, whereStr string, orderStr string, args ...interface{}) []*TableschemaVersion {
	return m.Gets(tx, 1, top, whereStr, orderStr, args...)
}

//GetTopss 获取前几条数据，需要完整的sql语句
func (m TableschemaVersionMapper) GetTopss(tx *sqlx.Tx, top int, sqlStr string, args ...interface{}) []*TableschemaVersion {
	return m.Getss(tx, 1, top, sqlStr, args...)
}

//GetFirst 获取第1条数据
func (m TableschemaVersionMapper) GetFirst(tx *sqlx.Tx, whereStr string, orderStr string, args ...interface{}) *TableschemaVersion {
	items := m.GetTops(tx, 1, whereStr, orderStr, args...)
	if len(items) == 0 {
		return nil
	}
	return items[0]
}

//GetCount 获取数量
func (m TableschemaVersionMapper) GetCount(tx *sqlx.Tx, whereStr string, args ...interface{}) int {
	sqlStr := "select count(*) from " + m.TableName
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}

	count := m.ScanInt(tx, sqlStr, args...)
	return count
}

//GetCountss 获取数量，需要完整的sql语句
func (m TableschemaVersionMapper) GetCountss(tx *sqlx.Tx, sqlStr string, args ...interface{}) int {
	count := m.ScanInt(tx, sqlStr, args...)
	return count
}

//GetStringss 获取字符串，需要完整的sql语句
func (m TableschemaVersionMapper) GetStringss(tx *sqlx.Tx, sqlStr string, args ...interface{}) string {
	str := m.ScanString(tx, sqlStr, args...)
	return str
}

//GetFloatss 获取数值，需要完整的sql语句
func (m TableschemaVersionMapper) GetFloatss(tx *sqlx.Tx, sqlStr string, args ...interface{}) float64 {
	v := m.ScanFloat(tx, sqlStr, args...)
	return v
}

//Del 根据id删除
func (m TableschemaVersionMapper) Del(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.DeleteOrUpdateItems(tx, sqlStr, args...)
}

//Dels 根据sql删除
//whereStr可以为空，不要where部分
func (m TableschemaVersionMapper) Dels(tx *sqlx.Tx, whereStr string, args ...interface{}) int {
	sqlStr := "delete from " + m.TableName
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}

	return m.DeleteOrUpdateItems(tx, sqlStr, args...)
}

//Delss 根据sql删除，需要完整的sql语句
func (m TableschemaVersionMapper) Delss(tx *sqlx.Tx, sqlStr string, args ...interface{}) int {
	return m.DeleteOrUpdateItems(tx, sqlStr, args...)
}

//Inserts 插入数据
func (m TableschemaVersionMapper) Inserts(tx *sqlx.Tx, items []*TableschemaVersion) int {
	var total int
	for _, item := range items {
		_, count := m.Insert(tx, item)
		total = total + count
	}
	return total
}

//Updates 更新数据
func (m TableschemaVersionMapper) Updates(tx *sqlx.Tx, updateStr string, whereStr string, args ...interface{}) int {
	sqlStr := " update " + m.TableName + " set " + updateStr
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}
	count := m.DeleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Updatess 更新数据、需要完整的sql语句
func (m TableschemaVersionMapper) Updatess(tx *sqlx.Tx, sqlStr string, args ...interface{}) int {
	count := m.DeleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Insert 插入数据
func (m TableschemaVersionMapper) Insert(tx *sqlx.Tx, item *TableschemaVersion) (int, int) {
	sqlStr := "insert into " + m.TableName + "(tableName,version,updateTime) values (?,?,?)"
	var args = []interface{}{item.TableName, item.Version, item.UpdateTime}
	id, count := m.InsertItem(tx, sqlStr, args...)
	return id, count
}

//Update 更新数据
func (m TableschemaVersionMapper) Update(tx *sqlx.Tx, item *TableschemaVersion) int {
	sqlStr := "update " + m.TableName + " set tableName=?,version=?,updateTime=? where id=? "
	var args = []interface{}{item.TableName, item.Version, item.UpdateTime, item.ID}
	count := m.DeleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//UpdateWithExtWhereStr 更新数据
func (m TableschemaVersionMapper) UpdateWithExtWhereStr(tx *sqlx.Tx, item *TableschemaVersion, extWhereStr string, extArgs ...interface{}) int {
	sqlStr := "update " + m.TableName + " set tableName=?,version=?,updateTime=? where id=? "
	var args = []interface{}{item.TableName, item.Version, item.UpdateTime, item.ID}
	if extWhereStr != "" {
		sqlStr = sqlStr + " " + extWhereStr
		for _, extArg := range extArgs {
			args = append(args, extArg)
		}
	}
	count := m.DeleteOrUpdateItems(tx, sqlStr, args...)
	return count
}
