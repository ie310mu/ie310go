/*

注意，route模块依赖于session模块，且session.GlobalSessions必须要初始化

*/

package route

import (
	"reflect"

	"github.com/ie310mu/ie310go/session"
)

//IService ..
type IService interface {
	RegisterMethod(name string, m reflect.Value)
	HandleRequest(rs RequestResponse) interface{}
}

//RequestResponse ..
type RequestResponse interface {
	GetStringFromForm(paramName string) string //从参数表单中获取一个字符串参数
	BeforeHandle()                             //可以做一些与协议相关的准备或检查工作，如http的跨域检查
	AfterHandle()                              //可以做一些与协议相关的收尾工作
	GetSessionID() string
	GetSession(forceSessionid string) session.Store //http中根据sessionid，其他协议中根据会话id创建
	GetServiceName() string                         //获取本次请求的服务名，比如userService
	GetMethodName() string                          //获取本次请求的方法名称
	PrepareResult(r interface{}) interface{}        //准备结果
	GetRemoteIP() string
	GetRemotePort() string
}
