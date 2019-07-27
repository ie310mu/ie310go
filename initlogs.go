package ie310go

import (
	"github.com/ie310mu/ie310go/logs"

	"github.com/ie310mu/ie310go/common/logsagent"
)

var defaultFileLogsConfig = `{"filename":"ie310go.log"}`

//SetDefaultLogsInfo ..
func SetDefaultLogsInfo(c string) {
	defaultFileLogsConfig = c
}

//AppLogsInitFunc 日志初始化***********不指定时使用默认的文件、Console输出
var AppLogsInitFunc appInitFunc = initAppLogsAsDefault

//日志初始化
func initAppLogsAsDefault() {
	beeLogger := logs.GetBeeLogger()
	beeLogger.EnableErrorStack(true)
	beeLogger.EnableFuncCallDepth(true)
	beeLogger.SetLogFuncCallDepth(3)
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, defaultFileLogsConfig)

	//设置ie310go中db等模块的日志代理
	logsagent.SetLogger(beeLogger)
}

//InitLogs 日志初始化
func InitLogs(config string) {
	beeLogger := logs.GetBeeLogger()
	beeLogger.EnableErrorStack(true)
	beeLogger.EnableFuncCallDepth(true)
	beeLogger.SetLogFuncCallDepth(3)
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, config)

	//设置ie310go中db等模块的日志代理
	logsagent.SetLogger(beeLogger)
}
