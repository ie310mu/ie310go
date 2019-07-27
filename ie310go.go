package ie310go

import (
	"fmt"

	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/logs"
	"github.com/ie310mu/ie310go/route"
)

//Run 启动api服务
func Run(callback func()) {
	//异常处理
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught. ", err)
			logs.Warn("Press return to exit.")
			var empty string
			fmt.Scanln(&empty)
		}
	}()

	//应用程序初始化
	if AppLogsInitFunc != nil {
		AppLogsInitFunc()
	}
	initAppDb()
	if AppSessionInitFunc != nil {
		AppSessionInitFunc()
	}

	if callback != nil {
		callback()
	}
	route.Run()
}

type appInitFunc func()
