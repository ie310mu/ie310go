// Copyright 2018 The go-ethereum Authors
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
	"encoding/binary"
	"errors"
	"sync"
	"time"
)

//用testService、notificationTestService创建一个rpc服务器
func newTestServer() *Server {
	server := NewServer()
	server.idgen = sequentialIDGenerator()
	if err := server.RegisterName("test", new(testService)); err != nil {
		panic(err)
	}
	if err := server.RegisterName("nftest", new(notificationTestService)); err != nil {
		panic(err)
	}
	return server
}

//序列ID生成器
func sequentialIDGenerator() func() ID {
	var (
		mu      sync.Mutex
		counter uint64
	)
	return func() ID {
		mu.Lock()
		defer mu.Unlock()
		counter++
		id := make([]byte, 8)
		binary.BigEndian.PutUint64(id, counter)
		return encodeID(id)
	}
}

type testService struct{}

//参数
type Args struct {
	S string
}

//结果
type Result struct {
	String string
	Int    int
	Args   *Args
}

//没有参数、返回值的导出示例
func (s *testService) NoArgsRets() {}

func (s *testService) Echo(str string, i int, args *Args) Result {
	return Result{str, i, args}
}

func (s *testService) EchoWithCtx(ctx context.Context, str string, i int, args *Args) Result {
	return Result{str, i, args}
}

func (s *testService) Sleep(ctx context.Context, duration time.Duration) {
	time.Sleep(duration)
}

func (s *testService) Rets() (string, error) {
	return "", nil
}

func (s *testService) InvalidRets1() (error, string) {
	return nil, ""
}

func (s *testService) InvalidRets2() (string, string) {
	return "", ""
}

func (s *testService) InvalidRets3() (string, string, error) {
	return "", "", nil
}

//用于客户端通知服务端调用自己的某个方法，成功后返回结果
func (s *testService) CallMeBack(ctx context.Context, method string, args []interface{}) (interface{}, error) {
	c, ok := ClientFromContext(ctx)
	if !ok {
		return nil, errors.New("no client")
	}
	var result interface{}
	err := c.Call(&result, method, args...)
	return result, err
}

//用于客户端通知服务端稍后调用自己的某个方法，会马上返回空结果
func (s *testService) CallMeBackLater(ctx context.Context, method string, args []interface{}) error {
	c, ok := ClientFromContext(ctx)
	if !ok {
		return errors.New("no client")
	}
	go func() {
		<-ctx.Done()
		var result interface{}
		c.Call(&result, method, args...)
	}()
	return nil
}

func (s *testService) Subscription(ctx context.Context) (*Subscription, error) {
	return nil, nil
}

type notificationTestService struct {
	unsubscribed            chan string
	gotHangSubscriptionReq  chan struct{}
	unblockHangSubscription chan struct{}
}

func (s *notificationTestService) Echo(i int) int {
	return i
}

func (s *notificationTestService) Unsubscribe(subid string) {
	if s.unsubscribed != nil {
		s.unsubscribed <- subid
	}
}

func (s *notificationTestService) SomeSubscription(ctx context.Context, n, val int) (*Subscription, error) {
	notifier, supported := NotifierFromContext(ctx)
	if !supported {
		return nil, ErrNotificationsUnsupported
	}

	// By explicitly creating an subscription we make sure that the subscription id is send
	// back to the client before the first subscription.Notify is called. Otherwise the
	// events might be send before the response for the *_subscribe method.
	subscription := notifier.CreateSubscription()
	go func() {
		for i := 0; i < n; i++ {
			if err := notifier.Notify(subscription.ID, val+i); err != nil {
				return
			}
		}
		select {
		case <-notifier.Closed():
		case <-subscription.Err():
		}
		if s.unsubscribed != nil {
			s.unsubscribed <- string(subscription.ID)
		}
	}()
	return subscription, nil
}

// HangSubscription blocks on s.unblockHangSubscription before sending anything.
func (s *notificationTestService) HangSubscription(ctx context.Context, val int) (*Subscription, error) {
	notifier, supported := NotifierFromContext(ctx)
	if !supported {
		return nil, ErrNotificationsUnsupported
	}
	s.gotHangSubscriptionReq <- struct{}{}
	<-s.unblockHangSubscription
	subscription := notifier.CreateSubscription()

	go func() {
		notifier.Notify(subscription.ID, val)
	}()
	return subscription, nil
}

func (s *notificationTestService) Abc(ctx context.Context, val string) (*Subscription, error) {
	notifier, supported := NotifierFromContext(ctx)
	if !supported {
		return nil, ErrNotificationsUnsupported
	}
	subscription := notifier.CreateSubscription()

	go func() {
		notifier.Notify(subscription.ID, val)
	}()
	return subscription, nil
}
