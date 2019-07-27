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

package ie310go

import (
	"fmt"
	"testing"
	"time"

	jsonutil "github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/throw"
	"github.com/ie310mu/ie310go/logs"
	"github.com/ie310mu/ie310go/route"
)

func TestThriftClient(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("Exception has been caught. ", err)
		}
	}()

	initLogs()

	//getTimeMapParam
	getTimeMapParam := make(map[string]interface{})
	getTimeMapParam["caller"] = "ie310"
	getTimeMapParam["des"] = "abc"
	getTimeMapParam["serial"] = 3334
	getTimeParam := &GetTimeParam{Caller: "ie310", Des: "abc", Serial: 333}
	getTime2Param := &GetTimeParam{Caller: "ie310", Des: "abc", Serial: 444}

	//Thrift StructParam StringResult
	{
		client := route.NewClient(`thrift:\\127.0.0.1:8043`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getTime", "", getTimeParam, &r)
		throw.CheckErr(err)
		showResult("Thrift StructParam StringResult        testService.getTime ", getTimeParam, r)
	}

	//Thrift StructResult
	{
		client := route.NewClient(`thrift:\\127.0.0.1:8043`)
		defer client.Close()
		r := &route.Time2Result{}
		err := client.Call("testService", "getTime2", "", getTime2Param, r)
		throw.CheckErr(err)
		showResult("Thrift StructResult        testService.getTime2 ", getTime2Param, r)
	}
}

func TestGrpcClient(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("Exception has been caught. ", err)
		}
	}()

	initLogs()

	//getTimeMapParam
	getTimeMapParam := make(map[string]interface{})
	getTimeMapParam["caller"] = "ie310"
	getTimeMapParam["des"] = "abc"
	getTimeMapParam["serial"] = 3334
	getTimeParam := &GetTimeParam{Caller: "ie310", Des: "abc", Serial: 333}
	getTime2Param := &GetTimeParam{Caller: "ie310", Des: "abc", Serial: 444}

	//Grpc StructParam StringResult
	{
		client := route.NewClient(`grpc:\\127.0.0.1:8033`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getTime", "", getTimeParam, &r)
		throw.CheckErr(err)
		showResult("Grpc StructParam StringResult        testService.getTime ", getTimeParam, r)
	}

	//Grpc StructResult
	{
		client := route.NewClient(`grpc:\\127.0.0.1:8033`)
		defer client.Close()
		r := &route.Time2Result{}
		err := client.Call("testService", "getTime2", "", getTime2Param, r)
		throw.CheckErr(err)
		showResult("Grpc StructResult        testService.getTime2 ", getTime2Param, r)
	}
}

func TestClient(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("Exception has been caught. ", err)
		}
	}()

	initLogs()

	//getTimeMapParam
	getTimeMapParam := make(map[string]interface{})
	getTimeMapParam["caller"] = "ie310"
	getTimeMapParam["des"] = "abc"
	getTimeMapParam["serial"] = 3334
	getTimeParam := &GetTimeParam{Caller: "ie310", Des: "abc", Serial: 333}
	getTime2Param := &GetTimeParam{Caller: "ie310", Des: "abc", Serial: 444}

	//Tcp MapParam StringResult
	{
		client := route.NewClient(`tcp:\\127.0.0.1:8013`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getTime", "", getTimeMapParam, &r)
		throw.CheckErr(err)
		showResult("Tcp MapParam StringResult        testService.getTime ", getTimeParam, r)
	}

	//Tcp StructParam StringResult
	{
		client := route.NewClient(`tcp:\\127.0.0.1:8013`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getTime", "", getTimeParam, &r)
		throw.CheckErr(err)
		showResult("Tcp StructParam StringResult        testService.getTime ", getTimeParam, r)
	}

	//Tcp GetSessionId
	{
		client := route.NewClient(`tcp:\\127.0.0.1:8013`)
		defer client.Close()
		r := ""
		client.Call("testService", "getSessionID", "111", nil, &r)
		showResult("Tcp GetSessionId        testService.getSessionID ", nil, r)
		r2 := ""
		err := client.Call("testService", "getSessionID", "111", nil, &r2)
		throw.CheckErr(err)
		showResult("Tcp GetSessionId        testService.getSessionID ", nil, r2)
	}

	//Ipc StringResult
	{
		client := route.NewClient(`\\.\pipe\ie310`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getTime", "", getTimeParam, &r)
		throw.CheckErr(err)
		showResult("Ipc StringResult        testService.getTime ", getTimeParam, r)
	}

	//Ipc StructResult
	{
		client := route.NewClient(`\\.\pipe\ie310`)
		defer client.Close()
		r := &route.Time2Result{}
		err := client.Call("testService", "getTime2", "", getTime2Param, r)
		throw.CheckErr(err)
		showResult("Ipc StructResult        testService.getTime2 ", getTime2Param, r)
	}

	//Ipc GetSessionId
	{
		client := route.NewClient(`\\.\pipe\ie310`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getSessionID", "222", nil, &r)
		showResult("Ipc GetSessionId        testService.getSessionID ", nil, r)
		r2 := ""
		err = client.Call("testService", "getSessionID", "222", nil, &r2)
		throw.CheckErr(err)
		showResult("Ipc GetSessionId        testService.getSessionID ", nil, r2)
	}

	//Ws StructResult
	{
		client := route.NewClient(`ws://127.0.0.1:8023`)
		defer client.Close()
		r := &route.Time2Result{}
		err := client.Call("testService", "getTime2", "", getTime2Param, r)
		throw.CheckErr(err)
		showResult("Ws StructResult        testService.getTime2 ", getTime2Param, r)
	}

	//Ws GetSessionId
	{
		client := route.NewClient(`ws://127.0.0.1:8023`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getSessionID", "333", nil, &r)
		showResult("Ws GetSessionId        testService.getSessionID ", nil, r)
		r2 := ""
		err = client.Call("testService", "getSessionID", "333", nil, &r2)
		throw.CheckErr(err)
		showResult("Ws GetSessionId        testService.getSessionID ", nil, r2)
	}

	//Http StructParam StructResult
	{
		client := route.NewClient(`http://127.0.0.1:8003`)
		defer client.Close()
		r := &route.Time2Result{}
		err := client.Call("testService", "getTime2", "", getTime2Param, &r)
		throw.CheckErr(err)
		showResult("Http StructParam StructResult        testService.getTime2 ", getTime2Param, r)
	}

	//Http MapParam
	{
		client := route.NewClient(`http://127.0.0.1:8003`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getTime", "", getTimeMapParam, &r)
		throw.CheckErr(err)
		showResult("Http MapParam        testService.getTime ", getTimeMapParam, r)
	}

	//Http GetSessionId
	{
		client := route.NewClient(`http://127.0.0.1:8003`)
		defer client.Close()
		r := ""
		err := client.Call("testService", "getSessionID", "123", nil, &r)
		showResult("Http  GetSessionId        testService.getSessionID ", nil, r) //http下使用的是cookie中的sessionid不是这里指定的
		r2 := ""
		err = client.Call("testService", "getSessionID", "", nil, &r2)
		throw.CheckErr(err)
		showResult("Http GetSessionId        testService.getSessionID ", nil, r2) //没有存储cookie，应该会得到另一个sessionId
	}
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

func showResult(name string, to interface{}, back interface{}) {
	logs.Info("-----------" + name + "-----------")
	logs.Info("--> " + jsonutil.ToJSON(to))
	logs.Info("<-- " + jsonutil.ToJSON(back))
	logs.Info("")
}

//GetTimeParam
type GetTimeParam struct {
	Caller string `json:"caller,omitempty"    `
	Des    string `json:"des,omitempty"    `
	Serial int    `json:"serial,omitempty"    `
}

type id string

func TestTypeString(t *testing.T) {
	var id1 id
	id1 = "aaa"
	var id2 string
	id2 = fmt.Sprintf("%v", id1)
	logs.Info(id2)
}

func TestChan(t *testing.T) {
	initLogs()

	ch := make(chan int, 10) //存满了才会阻塞，只要有1个就可以被读取

	go func() {
		for {
			select {
			case v := <-ch:
				logs.Info(v)
			case <-time.After(1 * time.Second):
				logs.Info("read timeouut")
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(50 * time.Second)
}

func TestWriteChanAfterClose(t *testing.T) {
	initLogs()

	ch := make(chan int, 10)
	close(ch)
	ch <- 1
}

//go client 订阅 ipc
func TestClientSubscribeipc(t *testing.T) {
	initLogs()

	client := route.NewClient(`\\.\pipe\ie310`)
	defer client.Close()

	ch := make(chan route.SystemState, 10) //用于获取通知结果
	sub, err := client.Subscribe("testService", "getSystemState", nil, ch)
	throw.CheckErr(err)

	//循环等待数据
	go func() {
		for {
			select {
			case val := <-ch:
				logs.Info(val)
			}
		}
	}()

	//10s后取消订阅
	go func() {
		time.Sleep(10 * time.Second)
		sub.Unsubscribe()
	}()

	//15s后关闭主程
	time.Sleep(15 * time.Second)
}

//go client 订阅 ws
func TestClientSubscribews(t *testing.T) {
	initLogs()

	client := route.NewClient(`ws://127.0.0.1:8023`)
	defer client.Close()

	ch := make(chan route.SystemState, 10) //用于获取通知结果
	sub, err := client.Subscribe("testService", "getSystemState", nil, ch)
	throw.CheckErr(err)

	//循环等待数据
	go func() {
		for {
			select {
			case val := <-ch:
				logs.Info(val)
			}
		}
	}()

	//10s后取消订阅
	go func() {
		time.Sleep(10 * time.Second)
		sub.Unsubscribe()
	}()

	//15s后关闭主程
	time.Sleep(15 * time.Second)
}

//go client 订阅 tcp
func TestClientSubscribetcp(t *testing.T) {
	initLogs()

	client := route.NewClient(`tcp:\\127.0.0.1:8013`)
	defer client.Close()

	ch := make(chan route.SystemState, 10) //用于获取通知结果
	sub, err := client.Subscribe("testService", "getSystemState", nil, ch)
	throw.CheckErr(err)

	//循环等待数据
	go func() {
		for {
			select {
			case val := <-ch:
				logs.Info(val)
			}
		}
	}()

	//10s后取消订阅
	go func() {
		time.Sleep(10 * time.Second)
		sub.Unsubscribe()
	}()

	//15s后关闭主程
	time.Sleep(15 * time.Second)
}

func TestJson(t *testing.T) {
	jsonstr := `{"jsonrpc":"2.0","method":"ethrpc_subscription","params":{"subscription":"0xf07c96dc5eb82660704f9e5aaf159332","result":{"time":"2019-03-02 23:49:58.3971408 +0800 CST m=+124.056242201","state":"2019-03-02 23:49:58.3971408 +0800 CST m=+124.056242201"}}}`

	r := &ethrpcSubscriptionResultParams{Result: route.SystemState{}}
	er := &ethrpcSubscriptionResult{Params: r}
	jsonutil.FromJSON(jsonstr, er)
	logsagent.Info(jsonutil.ToJSON(er))
}

type ethrpcSubscriptionResult struct {
	ID     string                          `json:"jsonrpc,omitempty"`
	Method string                          `json:"method,omitempty"`
	Params *ethrpcSubscriptionResultParams `json:"params,omitempty"`
}

type ethrpcSubscriptionResultParams struct {
	Subscription string      `json:"subscription"`
	Result       interface{} `json:"result,omitempty"`
}

func TestTyWsd(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("Exception has been caught. ", err)
		}
	}()

	initLogs()

	//param
	param := make(map[string]interface{})
	param["data"] = "xxxxxx"

	//Tcp StructParam StringResult
	{
		client := route.NewClient(`tcp:\\127.0.0.1:8015`)
		defer client.Close()
		r := ""
		err := client.Call("networkingService", "tjWsdUpload", "", param, &r)
		throw.CheckErr(err)
		showResult("Tcp StructParam StringResult        networkingService.tjWsdUpload ", param, r)
	}

	// 注意，请求参数中必须要有id，不然无法正确返回
	//eth test中没有加id可以正确返回，是因为框架自动添加了此参数
	//不需要结束符，ethrpc框架会自动解析json字符串结束
	//   {"id":"1","method":"ethrpc_invoke","params":[{"service":"networkingService","m":"tjWsdUpload", "packageResult":"false", "data":"xxxxxx"}]}
}
