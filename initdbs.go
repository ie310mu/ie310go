package ie310go

import (
	"github.com/ilibs/gosql"
)

//SetDbInfo ..
func SetDbInfo(w gosql.Config, r gosql.Config) {
	useDb = true
	writableDbConfig = w
	readonlyDbConfig = r
}

//数据库初始化
func initAppDb() {
	if !useDb {
		return
	}

	configs := make(map[string]*gosql.Config)
	configs["default"] = &writableDbConfig
	configs["ro"] = &readonlyDbConfig
	gosql.Connect(configs)
}

//UseDb 标识是否使用数据库***********不指定时，不会设置数据库连接，如果设置为true，则必须设置WritableDbConfig、ReadonlyDbConfig
var useDb = false

//WritableDbConfig 可写数据库配置
var writableDbConfig gosql.Config

//ReadonlyDbConfig 只读数据库配置
var readonlyDbConfig gosql.Config
