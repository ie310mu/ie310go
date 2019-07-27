package route

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	jsonutil "github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/common/throw"
)

func newClientHTTP(rawurl string) *ClientHTTP {
	r := &ClientHTTP{url: rawurl}
	return r
}

//ClientHTTP ..
type ClientHTTP struct {
	ClientBase
	url           string
	readDeadline  time.Time
	writeDeadline time.Time
}

//Close ..
func (c *ClientHTTP) Close() {
}

//Call ..
func (c *ClientHTTP) Call(service string, method string, sessionid string, arg interface{}, r interface{}) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	//组装参数
	data := make(url.Values)
	c.appendValues(data, arg)
	if service != "" {
		data["service"] = []string{service}
	}
	if method != "" {
		data["m"] = []string{method}
	}
	if sessionid != "" {
		data["sessionId"] = []string{sessionid}
	}

	//logsagent.Info(jsonutil.ToJSON(data))
	resp, err := http.PostForm(c.url, data)
	throw.CheckErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	throw.CheckErr(err)
	jsonStr := string(body)

	result := &JSONServiceResult{Data: r}
	jsonutil.FromJSON(jsonStr, result)
	if result.State != 0 {
		errMsg := result.Message
		if errMsg == "" {
			errMsg = "调用错误"
		}
		panic(errMsg)
	}

	return nil
}

//SetDeadline ..
func (c *ClientHTTP) SetDeadline(t time.Time) {
	c.readDeadline = t
	c.writeDeadline = t
}

//SetReadDeadline ..
func (c *ClientHTTP) SetReadDeadline(t time.Time) {
	c.readDeadline = t
}

//SetWriteDeadline ..
func (c *ClientHTTP) SetWriteDeadline(t time.Time) {
	c.writeDeadline = t
}

func (c ClientHTTP) appendValues(data url.Values, arg interface{}) {
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
			data[fmt.Sprintf("%v", k)] = []string{fmt.Sprintf("%v", vv)}
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
			c.appendValues(data, vv)
		} else {
			name := c.getFieldName(f)
			data[name] = []string{fmt.Sprintf("%v", vv)}
		}
	}
}

func (c ClientHTTP) getFieldName(f reflect.StructField) string {
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

//Subscribe http不支持主订阅
func (c *ClientHTTP) Subscribe(service string, method string, arg interface{}, ch interface{}) (csc ClientSubscription, e error) {
	return nil, nil
}
