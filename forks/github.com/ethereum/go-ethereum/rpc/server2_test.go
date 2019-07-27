// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ie310mu/ie310go/common/convert"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/throw"
	"github.com/ie310mu/ie310go/logs"
)

func TestTcpServer(t *testing.T) {
	initLogs()

	server := newTestServer()
	defer server.Stop()

	service := new(ie310testService)
	if err := server.RegisterName("test", service); err != nil {
		logs.Error(err)
		return
	}

	listener, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		logs.Error(err)
		return
	}
	defer listener.Close()

	var c chan int
	go startListen(server, listener, c)
	<-c //这个其实一直没有信号返回，server.ServeListener需要改造？

	//客户端使用net.Dial("tcp", "127.0.0.1:10001")创建连接
}

func TestIpcServerPIPEWindow(t *testing.T) {
	initLogs()

	server := newTestServer()
	defer server.Stop()

	service := new(ie310testService)
	if err := server.RegisterName("test", service); err != nil {
		logs.Error(err)
		return
	}

	listener, err := ipcListen(`\\.\pipe\ie310`)
	if err != nil {
		logs.Error(err)
		return
	}
	defer listener.Close()

	var c chan int
	go startListen(server, listener, c)
	<-c //这个其实一直没有信号返回，server.ServeListener需要改造？

	//客户端使用Dial(`\\.\pipe\ie310`)创建连接
}
func TestHttpServer(t *testing.T) {
	initLogs()

	server := newTestServer()
	defer server.Stop()

	service := new(ie310testService)
	if err := server.RegisterName("test", service); err != nil {
		logs.Error(err)
		return
	}

	err := http.ListenAndServe(":10001", &httpHandler{server: server}) //这句会阻塞
	throw.CheckErr(err)

	//可以使用fiddler的composer模拟访问     也可在客户端使用Dial("http://127.0.0.1:10001")创建连接
	/*
		POST http://127.0.0.1:10001

		Headers:
		User-Agent: Fiddler
		Host: 127.0.0.1:10001
		Content-Type: application/json

		RequestBody:
		{"jsonrpc":"2.0","id":1,"method":"test_echo2"}

		返回：
		{"jsonrpc":"2.0","id":1,"result":"hello world2019.02.19 17:36:11"}

		注意：
		不能调用echo方法，因为没有client
	*/
}
func TestWebsocketServer(t *testing.T) {
	initLogs()

	server := newTestServer()
	defer server.Stop()

	service := new(ie310testService)
	if err := server.RegisterName("test", service); err != nil {
		logs.Error(err)
		return
	}

	var (
		listener net.Listener
		err      error
	)
	if listener, err = net.Listen("tcp", "127.0.0.1:10001"); err != nil {
		logs.Error(err)
		return
	}
	go NewWSServer([]string{"*"}, server).Serve(listener) //第1个参数是允许的域名

	var c chan int
	c <- 1

	//可以在chrome控制台中模拟访问     也可在客户端使用Dial("ws://127.0.0.1:10001")创建连接
	/*
				function WebSocketTest()
		         {
		               var ws = new WebSocket("ws://127.0.0.1:10001");
		               ws.onopen = function()
		               {
		                  setTimeout(function (){ws.send('{"jsonrpc":"2.0","id":1,"method":"test_echo"}');},1000);
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

				返回：
				{"jsonrpc":"2.0","id":1,"result":"hello world2019.02.19 18:11:59"}

				注意：
				不能调用echo方法，因为没有client
	*/
}

func startListen(server *Server, listener net.Listener, c chan int) {
	server.ServeListener(listener)
	c <- 1
}

type httpHandler struct {
	server *Server
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught. ")
			logsagent.Error(err)
		}
	}()

	h.server.ServeHTTP(w, r)
}

func createServer() {
	server := newTestServer()
	service := new(ie310testService)
	if err := server.RegisterName("test", service); err != nil {
		logs.Error(err)
	}
}

type ie310testService struct {
}

func (ie310testService) Echo(ctx context.Context) string {
	c, ok := ClientFromContext(ctx)
	if !ok {
		return "no client"
	}

	logs.Info("收到请求%v", c)
	time.Sleep(200 * time.Millisecond)
	return "hello world" + convert.DateTimeWithSecToStr(time.Now())
}

func (ie310testService) Echo2() string {
	time.Sleep(200 * time.Millisecond)
	return "hello world" + convert.DateTimeWithSecToStr(time.Now())
}

//有客户端订阅一次这个函数就会执行一次
func (s *ie310testService) SystemStateSubscription(ctx context.Context) (*Subscription, error) {
	notifier, supported := NotifierFromContext(ctx)
	if !supported {
		return nil, ErrNotificationsUnsupported
	}

	subscription := notifier.CreateSubscription()
	idstr := string(subscription.ID)
	logs.Info(subscription)

	nip := GetSystemStateNotiferInProcess()
	ss := make(chan interface{}, 10)
	nip.AddChan(idstr, ss)
	go func(nip NotifierInProcess) {
		defer func() {
			if err := recover(); err != nil {
				logsagent.Error("Exception has been caught when run "+nip.Name(), err)
			}
		}()
		nip.TryRun()
	}(nip)

	go func() {
		for {
			select {
			case <-notifier.Closed():
				nip.RemoveChan(idstr, ss)
				close(ss)
				logs.Info("client " + idstr + " closed")
				goto outfor
			case err := <-subscription.Err():
				nip.RemoveChan(idstr, ss)
				close(ss)
				logs.Info("client " + idstr + " err：" + fmt.Sprintf("%v", err))
				goto outfor
			case msg := <-ss:
				logs.Info(idstr + ":    " + fmt.Sprintf("%v", msg))
				notifier.Notify(subscription.ID, msg)
			}
		}
	outfor:
	}()

	return subscription, nil
}

var ssnfp = genSystemStateNotiferInProcess()

func GetSystemStateNotiferInProcess() NotifierInProcess {
	return ssnfp
}

func initLogs() {
	beeLogger := logs.GetBeeLogger()
	beeLogger.EnableErrorStack(true)
	beeLogger.EnableFuncCallDepth(true)
	beeLogger.SetLogFuncCallDepth(3)
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, `{"filename":"ie310go.log"}`)
	//logs.Async(0)

	//设置ie310go中db等模块的日志代理
	logsagent.SetLogger(beeLogger)
}

type SystemState struct {
	Time  string
	State string
}

func genSystemStateNotiferInProcess() *SystemStateNotiferInProcess {
	r := &SystemStateNotiferInProcess{}
	r.chsLock = new(sync.Mutex)
	r.chs = make(map[string]chan interface{})
	return r
}

type SystemStateNotiferInProcess struct {
	BaseNotiferInProcess
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
			data := &SystemState{Time: time.Now().String(), State: time.Now().String()}
			s.notiferChs(data)
		}
	}()
}

//Stop ..
func (s *SystemStateNotiferInProcess) Stop() {
	atomic.CompareAndSwapInt32(&s.run, 1, 0)
}

type NotifierInProcess interface {
	Name() string
	AddChan(key string, ch chan interface{})
	RemoveChan(key string, ch chan interface{})
	TryRun()
	Stop()
}

type BaseNotiferInProcess struct {
	chsLock *sync.Mutex
	chs     map[string]chan interface{}
}

//AddChan ..
func (nip *BaseNotiferInProcess) AddChan(key string, ch chan interface{}) {
	nip.chsLock.Lock()
	defer nip.chsLock.Unlock()

	nip.chs[key] = ch
}

//RemoveChan ..
func (nip *BaseNotiferInProcess) RemoveChan(key string, ch chan interface{}) {
	nip.chsLock.Lock()
	defer nip.chsLock.Unlock()

	delete(nip.chs, key)
}

//Name ..
func (nip *BaseNotiferInProcess) Name() string {
	return "BaseNotiferInProcess"
}

//notiferChs ..
func (nip *BaseNotiferInProcess) notiferChs(data interface{}) {
	nip.chsLock.Lock()
	defer nip.chsLock.Unlock()

	for _, ch := range nip.chs {
		ch <- data
	}
}
