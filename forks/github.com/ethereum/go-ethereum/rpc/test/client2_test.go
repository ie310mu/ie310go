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
	"bufio"
	"context"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/ie310mu/ie310go/forks/github.com/ethereum/go-ethereum/rpc"
	jsonutil "github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/logs"
)

func TestTcpClient(t *testing.T) {
	initLogs()

	var (
		request  = `{"jsonrpc":"2.0","id":1,"method":"test_echo"}` + "\n"
		deadline = time.Now().Add(10 * time.Second)
	)

	conn, err := net.Dial("tcp", "127.0.0.1:10001")
	if err != nil {
		logs.Error(err)
		return
	}
	defer conn.Close()

	conn.SetDeadline(deadline)
	conn.Write([]byte(request))
	conn.(*net.TCPConn).CloseWrite()
	buf := make([]byte, 2000)
	n, err := conn.Read(buf)
	if err != nil {
		logs.Error(err)
	}
	str := string(buf[:n])
	logs.Info(str) //tcp中此种调用方法输出完整结果
}

func TestIpcClientPIPEWindow(t *testing.T) {
	initLogs()

	client, err := rpc.Dial(`\\.\pipe\ie310`)
	if err != nil {
		logs.Error(err)
		return
	}
	defer client.Close()

	var resp string
	if err := client.Call(&resp, "test_echo"); err != nil {
		logs.Error(err)
	}
	jsonstr := jsonutil.ToJSON(resp)
	logs.Info(jsonstr) //不包含外层的jsonMessage包
}

func TestHttpClient(t *testing.T) {
	initLogs()

	client, err := rpc.Dial("http://127.0.0.1:10001")
	if err != nil {
		logs.Error(err)
		return
	}
	defer client.Close()

	var resp string
	if err := client.Call(&resp, "test_echo"); err != nil {
		logs.Error(err)
	}
	jsonstr := jsonutil.ToJSON(resp)
	logs.Info(jsonstr)

	if err := client.Call(&resp, "test_echo2"); err != nil {
		logs.Error(err)
	}
	jsonstr = jsonutil.ToJSON(resp)
	logs.Info(jsonstr)
}

func TestWebsocketClient(t *testing.T) {
	initLogs()

	client, err := rpc.Dial("ws://127.0.0.1:10001")
	if err != nil {
		logs.Error(err)
		return
	}
	defer client.Close()

	var resp string
	if err := client.Call(&resp, "test_echo"); err != nil {
		logs.Error(err)
	}
	jsonstr := jsonutil.ToJSON(resp)
	logs.Info(jsonstr)

	if err := client.Call(&resp, "test_echo2"); err != nil {
		logs.Error(err)
	}
	jsonstr = jsonutil.ToJSON(resp)
	logs.Info(jsonstr)
}

//go client 订阅 ipc
func TestClientSubscribeipc(t *testing.T) {
	initLogs()

	client, err := rpc.Dial(`\\.\pipe\ie310`)
	if err != nil {
		logs.Error(err)
		return
	}
	defer client.Close()

	nc := make(chan SystemState, 10) //用于获取通知结果
	sub, err := client.Subscribe(context.Background(), "test", nc, "systemStateSubscription")
	if err != nil {
		logs.Info("can't subscribe:", err)
	}
	go func() {
		for {
			val := <-nc
			logs.Info(val)
		}
	}()

	go func() {
		time.Sleep(20 * time.Second)
		sub.Unsubscribe()
		select {
		case v := <-nc:
			logs.Info("received value after unsubscribe:", v)
		case err := <-sub.Err():
			if err != nil {
				logs.Info("Err returned a non-nil error after explicit unsubscribe: %q", err)
			}
		case <-time.After(1 * time.Second):
			logs.Info("subscription not closed within 1s after unsubscribe")
		}
	}()

	time.Sleep(30 * time.Second)
}

//go client 订阅 ws
func TestClientSubscribews(t *testing.T) {
	initLogs()

	client, err := rpc.Dial("ws://127.0.0.1:10001")
	if err != nil {
		logs.Error(err)
		return
	}
	defer client.Close()

	nc := make(chan SystemState, 10) //用于获取通知结果
	sub, err := client.Subscribe(context.Background(), "test", nc, "systemStateSubscription")
	if err != nil {
		t.Fatal("can't subscribe:", err)
	}
	go func() {
		for {
			val := <-nc
			logs.Info(val)
		}
	}()

	go func() {
		time.Sleep(30 * time.Second)
		sub.Unsubscribe()
	}()

	time.Sleep(50 * time.Second)
}

//go client 订阅 tcp
func TestClientSubscribetcp(t *testing.T) {
	initLogs()
	var (
		request = `{"jsonrpc":"2.0","id":1,"method":"test_subscribe", "params" : ["systemStateSubscription"]}` + "\n"
		//deadline = time.Now().Add(10 * time.Second)
	)

	conn, err := net.Dial("tcp", "127.0.0.1:10001")
	if err != nil {
		logs.Error(err)
		return
	}
	defer conn.Close()

	//conn.SetDeadline(deadline)
	conn.Write([]byte(request))
	//conn.(*net.TCPConn).CloseWrite()  //不能关，关了之后，得到第1次结果后就在两端报EOF错误
	readbuf := bufio.NewReader(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		sent, err := readbuf.ReadString('\n')
		if err != nil {
			logs.Error(err)
			return
		}
		sent = strings.TrimRight(sent, "\r\n")
		logs.Info(sent)
	}

	time.Sleep(50 * time.Second)
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
	Client rpc.ID
	Time   string
	State  string
}
