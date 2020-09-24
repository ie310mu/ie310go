package route

import (
	"context"
	"fmt"
	"reflect"
	"time"

	jsonutil "github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/common/throw"
	pb "github.com/ie310mu/ie310go/grpc"
	"google.golang.org/grpc"
)

func newClientGrpc(rawurl string) *ClientGrpc {
	conn, err := grpc.Dial(rawurl[7:], grpc.WithInsecure()) //grpc:\\127.0.0.1:8033
	throw.CheckErr(err)
	client := pb.NewIe310GoServcieClient(conn)
	r := &ClientGrpc{c: conn, client: client, duration: 30 * time.Second}
	return r
}

//ClientGrpc ..
type ClientGrpc struct {
	ClientBase
	c        *grpc.ClientConn
	client   pb.Ie310GoServcieClient
	closed   bool
	duration time.Duration
}

//Close ..
func (c *ClientGrpc) Close() {
	c.closed = true
	c.c.Close()
}

//Call ..
func (c *ClientGrpc) Call(service string, method string, sessionid string, arg interface{}, r interface{}) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	requestJSON := jsonutil.ToJSON(genGrpcRequestKv(service, method, sessionid, arg))

	ctx, cancel := context.WithTimeout(context.Background(), c.duration)
	defer cancel()
	result, err := c.client.Invoke(ctx, &pb.GrpcRequest{Data: requestJSON})
	throw.CheckErr(err)

	jsr := &JSONServiceResult{Data: r}
	jsonutil.FromJSON(result.Data, jsr)
	if jsr.State != 0 {
		errMsg := jsr.Message
		if errMsg == "" {
			errMsg = "调用错误"
		}
		panic(errMsg)
	}

	return nil
}

//SetDeadline ..
func (c *ClientGrpc) SetDeadline(t time.Time) {
	c.duration = t.Sub(time.Now())
	//c.c.SetDeadline(t)
}

//SetReadDeadline ..
func (c *ClientGrpc) SetReadDeadline(t time.Time) {
	c.duration = t.Sub(time.Now())
	//c.c.SetReadDeadline(t)
}

//SetWriteDeadline ..
func (c *ClientGrpc) SetWriteDeadline(t time.Time) {
	c.duration = t.Sub(time.Now())
	//c.c.SetWriteDeadline(t)
}

//Subscribe ..
func (c *ClientGrpc) Subscribe(service string, method string, arg interface{}, ch interface{}) (csc ClientSubscription, e error) {
	return nil, nil
}

func genGrpcRequestKv(service string, method string, sessionid string, arg interface{}) map[string]interface{} {
	kv := make(map[string]interface{})
	appendGrpcKv(kv, arg)

	if service != "" {
		kv["service"] = service
	}
	if method != "" {
		kv["m"] = method
	}
	if sessionid != "" {
		kv["sessionId"] = sessionid
	}

	return kv
}

func appendGrpcKv(kv map[string]interface{}, arg interface{}) {
	if obj.InterfaceIsNil(arg) {
		return
	}

	//获取类型和值
	t := reflect.TypeOf(arg)
	for t.Kind() == reflect.Ptr { //如果是指针，遍历取到对应的对象
		t = t.Elem()
	}
	v := reflect.ValueOf(arg)
	for v.Kind() == reflect.Ptr { //如果是指针，遍历取到对应的对象
		v = v.Elem()
	}

	//看是否是map  //如果是map，值是struct时，不会递归处理（map的值只支持基础类型）
	if t.Kind() == reflect.Map {
		keys := v.MapKeys()
		for _, k := range keys {
			vv := v.MapIndex(k)
			kv[fmt.Sprintf("%v", k)] = vv.Interface() //todo..........这样取是否正确？
		}
		return
	}

	//下面只处理struct
	if t.Kind() != reflect.Struct {
		return
	}

	//遍历字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}

		//获取字段
		field := v.FieldByName(f.Name)
		if !field.IsValid() {
			continue
		}

		//取值、类型
		vv := field.Interface()
		tt := reflect.TypeOf(vv)
		for tt.Kind() == reflect.Ptr { //如果是指针，遍历取到对应的对象
			tt = tt.Elem()
		}

		//字段还是结构体时，递归处理，否则可以添加值
		if tt.Kind() == reflect.Struct {
			appendGrpcKv(kv, vv)
		} else {
			name := getFieldName(f)
			kv[name] = vv //todo..........
		}
	}
}
