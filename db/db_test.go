package db

import (
	"testing"
	"time"

	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/logs"
	"github.com/ie310mu/ie310go/forks/github.com/ilibs/gosql"
	"github.com/ie310mu/ie310go/forks/github.com/jmoiron/sqlx"
)

func TestXXX(t *testing.T) {
	logs.Info("111")
	configs := make(map[string]*gosql.Config)
	configs["default"] = &gosql.Config{Enable: true, Driver: "mysql", Dsn: "zkb:password@tcp(127.0.0.1:3307)/zkb?charset=utf8&parseTime=True&loc=Asia%2FShanghai", ShowSql: true, MaxOpenConns: 20}
	gosql.Connect(configs)

	mp := GetDataHistoryMapper("")
	s1 := mp.Get(nil, "00008550-aaf2-48e2-90f7-51c8891170e3")
	logs.Info(s1)
}

//DataHistory ..
type DataHistory struct {
	ID           string    `json:"id"  db:"id"  `
	DataID       string    `json:"dataId"  db:"dataId"  `
	Sort         int       `json:"sort"  db:"sort"  `
	CreateTime   time.Time `json:"createTime"  db:"createTime"  `
	CreateUserID string    `json:"createUserId"  db:"createUserId"  `
	Type         string    `json:"type"  db:"type"  `
	Description  string    `json:"description"  db:"description"  `
}

//DataHistoryMapper 所有只读操作不会抛出异常，只会返回空数据，可写操作则会抛出异常
type DataHistoryMapper struct {
	BaseMapper
}

//GetDataHistoryMapper ..
func GetDataHistoryMapper(db string) DataHistoryMapper {
	var mp DataHistoryMapper
	mp.TableName = "data_history"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//将interface转换为*model.DataHistory
func (m DataHistoryMapper) convertItem(v interface{}) *DataHistory {
	if obj.InterfaceIsNil(v) {
		return nil
	}
	return v.(*DataHistory)
}

//Get 根据id获取
func (m DataHistoryMapper) Get(tx *sqlx.Tx, id interface{}) *DataHistory {
	sqlStr := "select * from " + m.TableName + " where id = ? "
	var args = []interface{}{id}

	item := &DataHistory{}
	r := m.ScanItem(tx, item, sqlStr, args...)
	return m.convertItem(r)
}

//Exists 根据id判断数据是否存在
func (m DataHistoryMapper) Exists(tx *sqlx.Tx, id interface{}) bool {
	sqlStr := "select count(id) from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	count := m.ScanInt(tx, sqlStr, args...)
	return count > 0
}

//GetByCode 根据code获取，需要有code字段
func (m DataHistoryMapper) GetByCode(tx *sqlx.Tx, code interface{}) *DataHistory {
	sqlStr := "select * from " + m.TableName + " where code = ? "
	var args = []interface{}{code}

	item := &DataHistory{}
	r := m.ScanItem(tx, item, sqlStr, args...)
	return m.convertItem(r)
}

//Gets 获取多条数据
//pageIndex从1开始
//whereStr、orderStr可以为空，不要where、order by部分
func (m DataHistoryMapper) Gets(tx *sqlx.Tx, pageIndex int, rowsInPage int, whereStr string, orderStr string, args ...interface{}) []*DataHistory {
	sqlStr := "select * from " + m.TableName
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}
	if orderStr != "" {
		sqlStr = sqlStr + " order by " + orderStr
	}

	return m.Getss(tx, pageIndex, rowsInPage, sqlStr, args...)
}

//Getss 获取多条数据
//pageIndex从1开始
//sqlStr可完全自定义
func (m DataHistoryMapper) Getss(tx *sqlx.Tx, pageIndex int, rowsInPage int, sqlStr string, args ...interface{}) []*DataHistory {
	items := make([]*DataHistory, 0)
	m.ScanItems(tx, pageIndex, rowsInPage, &items, sqlStr, args...)
	return items
}

//GetTops ..
func (m DataHistoryMapper) GetTops(tx *sqlx.Tx, top int, whereStr string, orderStr string, args ...interface{}) []*DataHistory {
	return m.Gets(tx, 1, top, whereStr, orderStr, args...)
}

//GetFirst ..
func (m DataHistoryMapper) GetFirst(tx *sqlx.Tx, whereStr string, orderStr string, args ...interface{}) *DataHistory {
	items := m.GetTops(tx, 1, whereStr, orderStr, args...)
	if len(items) == 0 {
		return nil
	}
	return items[0]
}

//GetCount ..
func (m DataHistoryMapper) GetCount(tx *sqlx.Tx, whereStr string, args ...interface{}) int {
	sqlStr := "select count(*) from " + m.TableName
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}

	count := m.ScanInt(tx, sqlStr, args...)
	return count
}

//GetCountss ..
func (m DataHistoryMapper) GetCountss(tx *sqlx.Tx, sqlStr string, args ...interface{}) int {
	count := m.ScanInt(tx, sqlStr, args...)
	return count
}

//GetFloatss ..
func (m DataHistoryMapper) GetFloatss(tx *sqlx.Tx, sqlStr string, args ...interface{}) float64 {
	v := m.ScanFloat(tx, sqlStr, args...)
	return v
}

//Del 根据id删除
func (m DataHistoryMapper) Del(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.DeleteOrUpdateItems(tx, sqlStr, args...)
}

//Dels 根据sql删除
//whereStr可以为空，不要where部分
func (m DataHistoryMapper) Dels(tx *sqlx.Tx, whereStr string, args ...interface{}) int {
	sqlStr := "delete from " + m.TableName
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}

	return m.DeleteOrUpdateItems(tx, sqlStr, args...)
}

//Inserts 插入数据
func (m DataHistoryMapper) Inserts(tx *sqlx.Tx, items []*DataHistory) int {
	var total int
	for _, item := range items {
		_, count := m.Insert(tx, item)
		total = total + count
	}
	return total
}

//Updates 更新数据
func (m DataHistoryMapper) Updates(tx *sqlx.Tx, updateStr string, whereStr string, args ...interface{}) int {
	sqlStr := " update " + m.TableName + " set " + updateStr
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}
	count := m.DeleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Updatess 更新数据
//sqlStr可完全自定义
func (m DataHistoryMapper) Updatess(tx *sqlx.Tx, sqlStr string, args ...interface{}) int {
	count := m.DeleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Insert 插入数据
func (m DataHistoryMapper) Insert(tx *sqlx.Tx, item *DataHistory) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,dataId,createTime,createUserId,type,description) values (?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.DataID, item.CreateTime, item.CreateUserID, item.Type, item.Description}
	id, count := m.InsertItem(tx, sqlStr, args...)
	return id, count
}

//Update 更新数据
func (m DataHistoryMapper) Update(tx *sqlx.Tx, item *DataHistory) int {
	sqlStr := "update " + m.TableName + " set id=?,dataId=?,createTime=?,createUserId=?,type=?,description=? where id=? "
	var args = []interface{}{item.ID, item.DataID, item.CreateTime, item.CreateUserID, item.Type, item.Description, item.ID}
	count := m.DeleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//customContentBegin
//customContentEnd
