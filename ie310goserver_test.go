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


//下面这个是访问普通方法，其实用http端口是一样的，在ws订阅的示例见后续
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




//===================订阅----服务端实现==================
package service

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/ie310mu/ie310go/common/convert"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/route"
)

func init() {
	route.Add(&LocalService{})
}

type LocalService struct {
	route.BaseService
}

// func (s LocalService) Login(args *route.ServiceArgs) {
// 	dataId := args.GetStringParamWithCheck("dataId", false)
// 	dataType := args.GetStringParamWithCheck("dataType", false)
// }

//要用全局变量，避免创建多个实例
var systemStateNip = genSystemStateNotiferInProcess()

//GetSystemState ..
func (s LocalService) GetSystemState(args *route.ServiceArgs) route.NotifierInProcess {
	return systemStateNip
}

//SystemState ..
type SystemState struct {
	Time  string `json:"time"`
	State string `json:"state"`
}

func genSystemStateNotiferInProcess() *SystemStateNotiferInProcess {
	r := &SystemStateNotiferInProcess{}
	r.ChsLock = new(sync.Mutex)
	r.Chs = make(map[string]chan interface{})
	return r
}

//SystemStateNotiferInProcess ..
type SystemStateNotiferInProcess struct {
	route.BaseNotiferInProcess
	run int32
}

//Name ..
func (s *SystemStateNotiferInProcess) Name() string {
	return "SystemStateNotiferInProcess"
}

//TryRun ..
func (s *SystemStateNotiferInProcess) TryRun() {
	if atomic.LoadInt32(&s.run) == 1 {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("SystemStateNotiferInProcess run error： ")
			logsagent.Error(err)
			atomic.StoreInt32(&s.run, 0)
		}
	}()

	atomic.StoreInt32(&s.run, 1)

	go func() {
		for {
			if atomic.LoadInt32(&s.run) == 0 {
				break
			}
			time.Sleep(3 * time.Second)
			data := &SystemState{Time: convert.DateTimeWithSecToStr(time.Now()), State: "normal"}
			s.NotiferChs(data)
		}
	}()
}

//Stop ..
func (s *SystemStateNotiferInProcess) Stop() {
	atomic.CompareAndSwapInt32(&s.run, 1, 0)
}


	wsConfig := route.ServerEthrpcConfig{PortOrName: "8023", Type: route.WS}
	wssrv := route.NewServerEthrpc(wsConfig, "wsServer")
	exportServices(route.Services, wssrv)
	route.RegisterServer(wssrv)



*/

//===================订阅-----js实现==================
/*

js订阅：
var ws = new WebSocket("ws://127.0.0.1:22223");
	ws.onopen = function()
	{
		setTimeout(function (){ws.send('{"id":"1","method":"ethrpc_subscribe","params":["sub", {"m":"dataChanged","service":"localService"}]}');},1000);
	};
	ws.onmessage = function (evt)
	{
		var data = JSON.parse(evt.data);  //注意evt.data是字符串
		if(data.id == "1"){
			console.info("订阅成功通知：");
			console.info(data);
		}else{
			console.info("消息：");
			//console.info(data);
			var result = data.params.result;
			console.info(result);
		}
	};
	ws.onclose = function()
	{
	};
	ws.onerror = function(evt)
	{
		console.info(evt);
	};

返回：
{"jsonrpc":"2.0","id":"1","result":"0xd788ab3f27fe0170838ab204f288b466"}
{"jsonrpc":"2.0","method":"ethrpc_subscription","params":{"subscription":"0xd788ab3f27fe0170838ab204f288b466","result":{"time":"2020.04.07 11:51:47","state":"normal"}}}
.....
......

取消订阅：
//注意，这里的ws要用上面的订阅时的ws对象，不然不是同一个会话
ws.send('{"id":"1","method":"ethrpc__unsubscribe","params":["0xf615a59bd59ec4ad5d3728d537bdf7"]}');

//网页直接关闭的话服务端会不会取消订阅？？会：
2020/04/07 12:14:08.227 [I] [BaseService.go:83]  begin invoke method  localService.GetSystemState
2020/04/07 12:14:09.339 [I] [ServerEthrpc.go:232]  &{2020.04.07 12:14:09 normal}  -->  0x511ff7635bb0485486ab6d678b804765
2020/04/07 12:14:12.339 [I] [ServerEthrpc.go:232]  &{2020.04.07 12:14:12 normal}  -->  0x511ff7635bb0485486ab6d678b804765
2020/04/07 12:14:15.340 [I] [ServerEthrpc.go:232]  &{2020.04.07 12:14:15 normal}  -->  0x511ff7635bb0485486ab6d678b804765
2020/04/07 12:14:18.340 [I] [ServerEthrpc.go:232]  &{2020.04.07 12:14:18 normal}  -->  0x511ff7635bb0485486ab6d678b804765
2020/04/07 12:14:18.473 [I] [ServerEthrpc.go:229]  client 0x511ff7635bb0485486ab6d678b804765 err in SystemStateNotiferInProcess：EOF




主动通知示例：
package service

import (
	"sync"
	"sync/atomic"

	"github.com/ie310mu/ie310go/route"
)

func init() {
	route.Add(&LocalService{})
}

type LocalService struct {
	route.BaseService
}

func (s LocalService) NotifyDataChanged(args *route.ServiceArgs) {
	dataId := args.GetStringParamWithCheck("dataId", false)
	dataType := args.GetStringParamWithCheck("dataType", false)

	v := DataChangedArgs{DataId: dataId, DataType: dataType}
	dataChangedNip.notify(v)
}

//要用全局变量，避免创建多个实例
var dataChangedNip = genDataChangedNotiferInProcess()

func (s LocalService) DataChanged(args *route.ServiceArgs) route.NotifierInProcess {
	return dataChangedNip
}

func genDataChangedNotiferInProcess() *DataChangedNotiferInProcess {
	r := &DataChangedNotiferInProcess{}
	r.ChsLock = new(sync.Mutex)
	r.Chs = make(map[string]chan interface{})
	return r
}

type DataChangedNotiferInProcess struct {
	route.BaseNotiferInProcess
	run int32
}

//Name ..
func (s *DataChangedNotiferInProcess) Name() string {
	return "DataChangedNotiferInProcess"
}

//TryRun ..
func (s *DataChangedNotiferInProcess) TryRun() {
}

func (s *DataChangedNotiferInProcess) notify(data DataChangedArgs) {
	s.NotiferChs(data)
}

//Stop ..
func (s *DataChangedNotiferInProcess) Stop() {
	atomic.CompareAndSwapInt32(&s.run, 1, 0)
}

type DataChangedArgs struct {
	DataId   string `json:"dataId"`
	DataType string `json:"dataType"`
}

*/
