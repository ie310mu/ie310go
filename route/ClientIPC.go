package route

import (
	"context"
	"fmt"

	"github.com/ie310mu/ie310go/forks/github.com/ethereum/go-ethereum/rpc"
	"github.com/ie310mu/ie310go/common/throw"
)

func newClientIPC(rawurl string) *ClientIPC {
	client, err := rpc.Dial(rawurl)
	throw.CheckErr(err)
	r := &ClientIPC{c: client}
	return r
}

//ClientIPC ..
type ClientIPC struct {
	ClientBase
	c *rpc.Client
}

//Close ..
func (c *ClientIPC) Close() {
	c.c.Close()
}

//Call ..
func (c *ClientIPC) Call(service string, method string, sessionid string, arg interface{}, r interface{}) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	result := &JSONServiceResult{Data: r}
	err := c.c.Call(result, "ethrpc_invoke", genEthrpcRequestKv(service, method, sessionid, arg))
	throw.CheckErr(err)

	if result.State != 0 {
		errMsg := result.Message
		if errMsg == "" {
			errMsg = "调用错误"
		}
		panic(errMsg)
	}

	return nil
}

//Subscribe ..
func (c *ClientIPC) Subscribe(service string, method string, arg interface{}, ch interface{}) (csc ClientSubscription, e error) {
	defer func() {
		if err := recover(); err != nil {
			csc = nil
			e = fmt.Errorf("%v", err)
		}
	}()

	sub, err := c.c.Subscribe(context.Background(), "ethrpc", ch, "sub", genEthrpcRequestKv(service, method, "", arg))
	throw.CheckErr(err)
	return &ClientSubscriptionIPC{sub: sub}, nil
}

//ClientSubscriptionIPC ..
type ClientSubscriptionIPC struct {
	sub *rpc.ClientSubscription
}

//Unsubscribe ..
func (s *ClientSubscriptionIPC) Unsubscribe() {
	s.sub.Unsubscribe()
}
