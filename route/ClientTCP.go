package route

import (
	"bufio"
	"fmt"
	"net"
	"reflect"
	"strings"
	"time"

	jsonutil "github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/throw"
)

func newClientTCP(rawurl string) *ClientTCP {
	conn, err := net.Dial("tcp", rawurl[6:]) //tcp:\\127.0.0.1:8013
	throw.CheckErr(err)
	readbuf := bufio.NewReader(conn)
	r := &ClientTCP{c: conn, readbuf: readbuf}
	return r
}

//ClientTCP ..
type ClientTCP struct {
	ClientBase
	c       net.Conn
	readbuf *bufio.Reader
	closed  bool
}

//Close ..
func (c *ClientTCP) Close() {
	c.closed = true
	c.c.Close()
}

//Call ..
func (c *ClientTCP) Call(service string, method string, sessionid string, arg interface{}, r interface{}) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	requestJSON := jsonutil.ToJSON(genEthrpcRequest(service, method, sessionid, arg)) + "\n"

	c.c.Write([]byte(requestJSON))
	sent, err := c.readbuf.ReadString('\n')
	throw.CheckErr(err)
	sent = strings.TrimRight(sent, "\r\n")

	result := &JSONServiceResult{Data: r}
	er := &ethrpcResult{Result: result}
	jsonutil.FromJSON(sent, er)

	return nil
}

//SetDeadline ..
func (c *ClientTCP) SetDeadline(t time.Time) {
	c.c.SetDeadline(t)
}

//SetReadDeadline ..
func (c *ClientTCP) SetReadDeadline(t time.Time) {
	c.c.SetReadDeadline(t)
}

//SetWriteDeadline ..
func (c *ClientTCP) SetWriteDeadline(t time.Time) {
	c.c.SetWriteDeadline(t)
}

//Subscribe ..
func (c *ClientTCP) Subscribe(service string, method string, arg interface{}, ch interface{}) (csc ClientSubscription, e error) {
	defer func() {
		if err := recover(); err != nil {
			csc = nil
			e = fmt.Errorf("%v", err)
		}
	}()

	chanVal := reflect.ValueOf(ch)
	if chanVal.Kind() != reflect.Chan || chanVal.Type().ChanDir()&reflect.SendDir == 0 {
		panic("ch must be a writable channel")
	}
	if chanVal.IsNil() {
		panic("channel given to Subscribe must not be nil")
	}

	requestJSON := jsonutil.ToJSON(genEthrpcSubscribeRequest(service, method, arg)) + "\n"
	c.c.Write([]byte(requestJSON))
	result, err := c.readbuf.ReadString('\n')
	throw.CheckErr(err)
	idstr := ""
	er := &ethrpcResult{Result: &idstr}
	jsonutil.FromJSON(result, er)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logsagent.Error(err)
			}
		}()

		for {
			if c.closed {
				break
			}

			sent, err := c.readbuf.ReadString('\n')
			throw.CheckErr(err)
			sent = strings.TrimRight(sent, "\r\n")
			//logsagent.Info(sent)

			func() {
				defer func() {
					if err := recover(); err != nil {
						logsagent.Info(sent)
						logsagent.Error(err)
					}
				}()

				elem := reflect.New(chanVal.Type().Elem()) //这是指针
				r := &ethrpcSubscriptionResultParams{Result: elem.Interface()}
				er := &ethrpcSubscriptionResult{Params: r}
				jsonutil.FromJSON(sent, er)
				if er.Method != "" { //method不为空才是正常的消息（取消订阅也会在这里读取到，但method为“”）
					//logsagent.Info(jsonutil.ToJSON(r.Result))
					//sc := reflect.SelectCase{Dir: reflect.SelectSend, Chan: chanVal, Send: reflect.ValueOf(r.Result)}
					chdata := reflect.ValueOf(r.Result)
					for chdata.Kind() == reflect.Ptr { //如果是指针，遍历取到对应的对象
						chdata = chdata.Elem()
					}
					chanVal.Send(chdata)
				}
			}()
		}
	}()

	return &ClientSubscriptionTCP{c: c, subid: idstr}, nil
}

//ClientSubscriptionTCP ..
type ClientSubscriptionTCP struct {
	c       *ClientTCP
	service string
	method  string
	subid   string
}

//Unsubscribe ..
func (s *ClientSubscriptionTCP) Unsubscribe() {
	s.c.closed = true

	requestJSON := jsonutil.ToJSON(genEthrpcUnsubscribeRequest(s.service, s.method, s.subid)) + "\n"
	s.c.c.Write([]byte(requestJSON))
}

type ethrpcSubscriptionResult struct {
	ID     string                          `json:"jsonrpc"`
	Method string                          `json:"method"`
	Params *ethrpcSubscriptionResultParams `json:"params"`
}

type ethrpcSubscriptionResultParams struct {
	Subscription string      `json:"subscription"`
	Result       interface{} `json:"result,omitempty"`
}
