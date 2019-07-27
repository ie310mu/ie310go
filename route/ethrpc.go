package route

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/ie310mu/ie310go/common/convert"
	"github.com/ie310mu/ie310go/common/obj"
)

//ethrpc的请求包
type ethrpcRequest struct {
	ID     string        `json:"id"    `
	Method string        `json:"method"    `
	Params []interface{} `json:"params"    `
}

//将Call的入参转换为ethrpc需要的格式(kv)
func genEthrpcRequestKv(service string, method string, sessionid string, arg interface{}) map[string]interface{} {
	kv := make(map[string]interface{})
	appendEthrpcKv(kv, arg)

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

//将Call的入参转换为ethrpc需要的格式(ethrpcRequest)
func genEthrpcRequest(service string, method string, sessionid string, arg interface{}) *ethrpcRequest {
	kv := genEthrpcRequestKv(service, method, sessionid, arg)
	return &ethrpcRequest{ID: sequentialIDGenerator()(), Method: "ethrpc_invoke", Params: []interface{}{kv}}
}

//订阅请求
func genEthrpcSubscribeRequest(service string, method string, arg interface{}) *ethrpcRequest {
	kv := genEthrpcRequestKv(service, method, "", arg)
	return &ethrpcRequest{ID: sequentialIDGenerator()(), Method: "ethrpc_subscribe", Params: []interface{}{"sub", kv}}
}

//取消订阅请求
func genEthrpcUnsubscribeRequest(service string, method string, subid string) *ethrpcRequest {
	return &ethrpcRequest{ID: sequentialIDGenerator()(), Method: "ethrpc__unsubscribe", Params: []interface{}{subid}}
}

type ethrpcResult struct {
	ID     string      `json:"id,omitempty"`
	Method string      `json:"method,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

func sequentialIDGenerator() func() string {
	var (
		mu      sync.Mutex
		counter uint64
	)
	return func() string {
		mu.Lock()
		defer mu.Unlock()
		counter++
		return convert.IntToStr(int(counter))
	}
}

func appendEthrpcKv(kv map[string]interface{}, arg interface{}) {
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
			appendEthrpcKv(kv, vv)
		} else {
			name := getFieldName(f)
			kv[name] = vv //todo..........
		}
	}
}

func getFieldName(f reflect.StructField) string {
	name := f.Name
	f1 := name[:1]
	f2 := name[1:]
	defaultName := strings.ToLower(f1) + f2

	tag := f.Tag.Get("json")

	if tag == "" || strings.Index(tag, "-") == 0 || tag == "omitempty" {
		return defaultName
	}

	index := strings.Index(tag, ",")
	if index == -1 {
		return defaultName
	}

	return tag[:index]
}
