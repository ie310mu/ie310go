package route

import (
	"net/url"
	"time"

	"github.com/ie310mu/ie310go/common/throw"
)

//Client 客户端
type Client interface {
	Close()

	//sessionid可以为空    //arg可以是struct，也可以是map
	Call(service string, method string, sessionid string, arg interface{}, r interface{}) (e error)

	//arg可以是struct，也可以是map
	Subscribe(service string, method string, arg interface{}, ch interface{}) (csc ClientSubscription, e error)

	SetDeadline(t time.Time)
	SetReadDeadline(t time.Time)
	SetWriteDeadline(t time.Time)
}

//ClientSubscription 订阅对象
type ClientSubscription interface {
	Unsubscribe()
}

//NewClient ..
func NewClient(rawurl string) Client {
	u, err := url.Parse(rawurl)
	throw.CheckErr(err)

	switch u.Scheme {
	case "thrift":
		return newClientThrift(rawurl)
	case "grpc":
		return newClientGrpc(rawurl)
	case "tcp":
		return newClientTCP(rawurl)
	case "ws", "wss":
		return newClientWs(rawurl)
	case "http", "https":
		return newClientHTTP(rawurl)
	// case "stdio":
	// 	return DialStdIO(ctx)
	case "":
		return newClientIPC(rawurl)
	default:
		return nil
	}
}

//ClientBase ..
type ClientBase struct {
}

//SetDeadline ..
func (c *ClientBase) SetDeadline(t time.Time) {
}

//SetReadDeadline ..
func (c *ClientBase) SetReadDeadline(t time.Time) {
}

//SetWriteDeadline ..
func (c *ClientBase) SetWriteDeadline(t time.Time) {
}
