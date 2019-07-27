package ie310go

import (
	"fmt"
	"testing"

	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/logs"
	"github.com/ie310mu/ie310go/route"
)

func TestRun(t *testing.T) {
	//异常处理
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught. ", err)
			logs.Warn("Press return to exit.")
			var empty string
			fmt.Scanln(&empty)
		}
	}()

	//cfg := route.InitTestConfig()
	//SetDefaultLogsInfo(`{"filename":"test.log"}`)
	//AppLogsInitFunc = xxxx
	//SetDbInfo(cfg.DataBase, cfg.DataBaseRO) //这句代码没有放到testservice.go中，因为route模块是独立的，不与db模块耦合
	//SetDefaultSessionInfo("JSESSIONID", false)
	//AppSessionInitFunc = xxxx

	cfg := route.InitTestConfig()
	SetDbInfo(cfg.DataBase, cfg.DataBaseRO)
	Run(nil)
}

/*                        客户端调用说明

//=========================http:
可以用route.Client对象进行调用，参考：ie310go_client_test.go

也可以使用js或浏览器访问url：
http://localhost:8003/testService.goss?m=getTime&caller=jack
{"state":0,"data":"hello, jack, the server time now is 2019.02.20 21:53:47"}


http://localhost:8003/testService.goss?m=getSessionID

http://127.0.0.1:8003/testService.goss?m=echo&caller=jack
{"state":0}

http://localhost:8004/testAdminService.goss?m=getServerStatus
{"state":0,"data":"normal"}

http://localhost:8003/testService.goss?m=add
{"state":0,"data":"4891dbd8-b038-4ccb-a4d0-3c57d17f4602"}

http://localhost:8003/testService.goss?m=get&id=386bd96a-b0e6-4373-9cef-4f1da0abc945
{"state":0,"data":{"id":"386bd96a-b0e6-4373-9cef-4f1da0abc945","dataId":"8222c553-0e88-4cb9-a169-fd2efe37bc85","sort":371965,"createTime":"2019-02-03T14:20:41+08:00","createUserId":"73eec28c-a182-4624-baa6-b34258d9eb53","type":"test","description":"test des"}}


=========================grpc:
可以用route.Client对象进行调用，参考：ie310go_client_test.go
也可以使用任何语言自行编码调用、解析（数据格式待补充）
&pb.GrpcRequest{Data: requestJSON}  requestJSON<----kv map

=========================thrift:
可以用route.Client对象进行调用，参考：ie310go_client_test.go
也可以使用任何语言自行编码调用、解析（数据格式待补充）
data<----kv map json


=========================tcp:
可以用route.Client对象进行调用，参考：ie310go_client_test.go
也可以使用任何语言自行编码调用、解析（数据格式待补充）




=========================ipc:
可以用route.Client对象进行调用，参考：ie310go_client_test.go
也可以使用任何语言自行编码调用、解析（数据格式待补充）





=========================ws:
可以用route.Client对象进行调用，参考：ie310go_client_test.go
也可以使用任何语言自行编码调用、解析（数据格式待补充）


function WebSocketTest()
		         {
		               var ws = new WebSocket("ws://127.0.0.1:8023");
		               ws.onopen = function()
		               {
		                  setTimeout(function (){ws.send('{"id":"1","method":"ethrpc_invoke","params":[{"m":"getList","sessionId":"CADD3638BF8C95718BEAD1852DE62804","service":"gwYqsbService"}]}');},1000);
		               };
		               ws.onmessage = function (evt)
		               {
		                  var received_msg = evt.data;
				          console.info(evt.data);
		               };
		               ws.onclose = function()
		               {
		               };
		               ws.onerror = function(evt)
		               {
				          console.info(evt);
		               };
		         }
				WebSocketTest();



//===================订阅==================
tpc/ipc/ws：
可以用route.Client对象进行调用，参考：ie310go_client_test.go
也可以使用任何语言自行编码调用、解析（数据格式待补充）


http:
不能使用route.Client调用，只能指定一个backurl让服务器回调
http://localhost:8003/testService.goss?m=sub&subMethod=getSystemState&backurl=http://localhost:8003/testService.goss?m=reciveSystemState
8856200d-d112-410b-8b09-a08cb67a57e4
http://localhost:8003/testService.goss?m=unSub&subId=5954e63d-e2f1-4d8a-8246-03baae1e97c9





*/
