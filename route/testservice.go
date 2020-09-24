package route

import (
	"fmt"
	"time"

	"github.com/ie310mu/ie310go/common/convert"
	"github.com/ie310mu/ie310go/common/guid"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/config"
	"github.com/ie310mu/ie310go/db"
	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

func registerTestServers() {
	httpConfig := ServerHTTPConfig{
		Port:          "8003",
		ServiceSuffix: "goss",
		Jsonp:         true,

		ServeStaticFunc: DefaultServeStatic,
		StaticDirs: map[string]string{
			"/this": ".",
			"/zkb":  "F:/ie310.code/011.ZKB/zkb/zkb",
		},
		DefaultStaticDir: ".",
	}
	srv := NewServerHTTP(httpConfig, "httpServer")
	srv.RegisterService(&TestService{})
	RegisterServer(srv)

	srvAdmin := NewServerHTTP(ServerHTTPConfig{Port: "8004", ServiceSuffix: "goss", Jsonp: true}, "adminServer")
	srvAdmin.RegisterService(&TestAdminService{})
	RegisterServer(srvAdmin)

	srvTCP := NewServerEthrpcDefaultTCP()
	srvTCP.RegisterService(&TestService{})
	RegisterServer(srvTCP)

	srvIpc := NewServerEthrpcDefaultIpc()
	srvIpc.RegisterService(&TestService{})
	RegisterServer(srvIpc)

	srvWs := NewServerEthrpcDefaultWs()
	srvWs.RegisterService(&TestService{})
	RegisterServer(srvWs)

	srvGrpc := NewServerGrpcDefault()
	srvGrpc.RegisterService(&TestService{})
	RegisterServer(srvGrpc)

	srvThrift := NewServerThriftDefault()
	srvThrift.RegisterService(&TestService{})
	RegisterServer(srvThrift)
}

//=========================服务=========================

//TestService ..
type TestService struct {
	BaseService
}

//Echo ..
func (s TestService) Echo(args *ServiceArgs) {
	caller := args.GetStringParam("caller")
	logsagent.Info("caller:" + caller)
}

//GetSessionID ..
func (s TestService) GetSessionID(args *ServiceArgs) string {
	return args.rs.GetSessionID()
}

//GetTime ..
func (s TestService) GetTime(args *ServiceArgs) string {
	caller := args.GetStringParam("caller")
	des := args.GetStringParam("des")

	now := time.Now()
	str := convert.DateTimeWithSecToStr(now)
	return "hello, " + caller + "," + des + ", the server time now is " + str
}

//GetTime2 ..
func (s TestService) GetTime2(args *ServiceArgs) Time2Result {
	caller := args.GetStringParam("caller")
	des := args.GetStringParam("des")

	now := time.Now()
	str := convert.DateTimeWithSecToStr(now)
	return Time2Result{Caller: caller, Des: des, Result: "hello, " + caller + "," + des + ", the server time now is " + str}
}

//Time2Result ..
type Time2Result struct {
	Caller string
	Des    string
	Result string
}

//TestAdminService ..
type TestAdminService struct {
	BaseService
}

//GetServerStatus ..
func (s TestAdminService) GetServerStatus(args *ServiceArgs) string {
	return "normal"
}

//=========================配置定义=========================

var testConfig TestConfig

//InitTestConfig ..
func InitTestConfig() TestConfig {
	config.ScanFromFile("testconfig.json", &testConfig)
	return testConfig
}

//TestConfig ..
type TestConfig struct {
	DataBase   gosql.Config
	DataBaseRO gosql.Config
}

//=========================测试数据库访问=========================
//注意，需要数据库有这张表与对应字段

var createTableStr = "CREATE TABLE `data_history` ( " +
	" `id` char(36) NOT NULL, " +
	" 	`dataId` char(36) DEFAULT NULL, " +
	" `sort` int(10) unsigned NOT NULL AUTO_INCREMENT, " +
	" `createTime` datetime DEFAULT NULL, " +
	" `createUserId` char(36) DEFAULT NULL, " +
	" `type` varchar(200) DEFAULT NULL, " +
	" `description` varchar(500) DEFAULT NULL, " +
	" PRIMARY KEY (`id`), " +
	" KEY `data_history_createTime` (`createTime`), " +
	" KEY `data_history_createType` (`type`), " +
	" KEY `data_history_createUserId` (`createUserId`), " +
	" KEY `data_history_dataId` (`dataId`), " +
	" KEY `data_history_sort` (`sort`) " +
	" ) ENGINE=InnoDB AUTO_INCREMENT=371964 DEFAULT CHARSET=utf8 ;"

//Add ..
func (s TestService) Add(args *ServiceArgs) string {
	defer func() {
		if err := recover(); err != nil {
			msg := "error, try to create table / or set database:           " + createTableStr + "           " + fmt.Sprintf("%v", err)
			panic(msg)
		}
	}()

	var item DataHistory
	item.ID = guid.Get()
	item.DataID = guid.Get()
	item.Sort = 1
	item.CreateTime = time.Now()
	item.CreateUserID = guid.Get()
	item.Type = "test"
	item.Description = "test des"

	mp := GetDataHistoryMapper("") //可写连接
	mp.Insert(nil, &item)
	return item.ID
}

//Get ..
func (s TestService) Get(args *ServiceArgs) *DataHistory {
	defer func() {
		if err := recover(); err != nil {
			msg := "error, try to create table / or set database:           " + createTableStr + "           " + fmt.Sprintf("%v", err)
			panic(msg)
		}
	}()

	id := args.GetStringParamWithCheck("id", true)
	mp := GetDataHistoryMapper("ro") //ro代表只读连接
	item := mp.Get(nil, id)
	return item
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

//DataHistoryMapper ..
type DataHistoryMapper struct {
	db.BaseMapper
}

//GetDataHistoryMapper 获取mapper对象
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

//Gets 获取多条数据、pageIndex从1开始、whereStr、orderStr可以为空，不要where、order by部分
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

//Getss 获取多条数据、需要完整的sql语句、pageIndex从1开始
func (m DataHistoryMapper) Getss(tx *sqlx.Tx, pageIndex int, rowsInPage int, sqlStr string, args ...interface{}) []*DataHistory {
	items := make([]*DataHistory, 0)
	m.ScanItems(tx, pageIndex, rowsInPage, &items, sqlStr, args...)
	return items
}

//GetTops 获取前几条数据
func (m DataHistoryMapper) GetTops(tx *sqlx.Tx, top int, whereStr string, orderStr string, args ...interface{}) []*DataHistory {
	return m.Gets(tx, 1, top, whereStr, orderStr, args...)
}

//GetTopss 获取前几条数据，需要完整的sql语句
func (m DataHistoryMapper) GetTopss(tx *sqlx.Tx, top int, sqlStr string, args ...interface{}) []*DataHistory {
	return m.Getss(tx, 1, top, sqlStr, args...)
}

//GetFirst 获取第1条数据
func (m DataHistoryMapper) GetFirst(tx *sqlx.Tx, whereStr string, orderStr string, args ...interface{}) *DataHistory {
	items := m.GetTops(tx, 1, whereStr, orderStr, args...)
	if len(items) == 0 {
		return nil
	}
	return items[0]
}

//GetCount 获取数量
func (m DataHistoryMapper) GetCount(tx *sqlx.Tx, whereStr string, args ...interface{}) int {
	sqlStr := "select count(*) from " + m.TableName
	if whereStr != "" {
		sqlStr = sqlStr + " where " + whereStr
	}

	count := m.ScanInt(tx, sqlStr, args...)
	return count
}

//GetCountss 获取数量，需要完整的sql语句
func (m DataHistoryMapper) GetCountss(tx *sqlx.Tx, sqlStr string, args ...interface{}) int {
	count := m.ScanInt(tx, sqlStr, args...)
	return count
}

//GetFloatss 获取数值，需要完整的sql语句
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

//Delss 根据sql删除，需要完整的sql语句
func (m DataHistoryMapper) Delss(tx *sqlx.Tx, sqlStr string, args ...interface{}) int {
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

//Updatess 更新数据、需要完整的sql语句
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
