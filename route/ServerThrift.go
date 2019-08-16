package route

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync/atomic"

	"github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/common/throw"
	"github.com/ie310mu/ie310go/session"

	"github.com/ie310mu/ie310go/forks/github.com/apache/thrift/lib/go/thrift"
	pb "github.com/ie310mu/ie310go/thrift/gen-go/thrift"
)

//ServerThriftConfig ..
type ServerThriftConfig struct {
	Port string `json:"port,omitempty"`
}

var defaultServeThriftConfig = ServerThriftConfig{Port: "8043"}

//NewServerThriftDefault ..
func NewServerThriftDefault() *ServerThrift {
	return &ServerThrift{services: make(map[string]IService), config: defaultServeThriftConfig, name: "defaultThriftServer", stopch: make(chan bool)}
}

//NewServerThrift ..
func NewServerThrift(cfg ServerThriftConfig, n string) *ServerThrift {
	return &ServerThrift{services: make(map[string]IService), config: cfg, name: n, stopch: make(chan bool)}
}

//ServerThrift ..
type ServerThrift struct {
	name         string
	config       ServerThriftConfig
	run          int32 //==1表示正在运行中
	stopch       chan bool
	services     map[string]IService
	thriftserver *thrift.TSimpleServer
}

//Name ..
func (s *ServerThrift) Name() string {
	return s.name
}

//Run ..
func (s *ServerThrift) Run() {
	if atomic.LoadInt32(&s.run) == 1 {
		logsagent.Warn("server " + s.Name() + " is already in running ")
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("server " + s.Name() + " run error： ")
			logsagent.Error(err)
			atomic.StoreInt32(&s.run, 0)
			logsagent.Info("server " + s.Name() + " set stopch to true")
			s.stopch <- true
			close(s.stopch)
		}
	}()

	atomic.StoreInt32(&s.run, 1)
	logsagent.Info("server " + s.Name() + " where listen at " + s.config.Port)

	s.thriftserver = s.newThriftServer()
	defer s.thriftserver.Stop()

	var c chan int = make(chan int)
	go s.startListen(c)
	<-c
}

func (s *ServerThrift) startListen(c chan int) {
	defer func() {
		c <- 1
	}()

	err := s.thriftserver.Serve()
	throw.CheckErr(err)
}

func (s *ServerThrift) newThriftServer() *thrift.TSimpleServer {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(":" + s.config.Port)
	throw.CheckErr(err)

	handler := &thriftHandler{srv: s}
	processor := pb.NewIe310goThirftServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	return server
}

type thriftHandler struct {
	srv *ServerThrift
}

func (h *thriftHandler) Invoke(ctx context.Context, data string) (r string, e error) {
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught in server "+h.srv.Name(), err)
			logsagent.Error(err)
			rr := JSONServiceResult{State: 1, Message: fmt.Sprintf("%v", err)}
			r = json.ToJSON(rr)
		}
	}()

	{
		kv := make(map[string]interface{})
		json.FromJSON(data, &kv) //注意传指针
		rr := h.invoke(ctx, kv)
		r = json.ToJSON(rr)
		return
	}
}

func (h *thriftHandler) invoke(ctx context.Context, kv map[string]interface{}) interface{} {
	rs := genRequestResponseThrift(ctx, h.srv.config, kv)
	serviceName := rs.GetServiceName()
	ss := h.getService(serviceName)
	if ss == nil {
		return JSONServiceResult{State: 1, Message: "servicePath not exists"}
	}

	m := rs.GetStringFromForm("m")
	m = strings.ToUpper(m)
	if m == "SUB" { //订阅请求
		panic("不支持订阅模式")
	} else if m == "UNSUB" { //取消订阅
		panic("不支持订阅模式")
	} else { //正常请求
		r := ss.HandleRequest(rs)
		return r
	}

}

//getService 查找服务，未找到时返回nil
func (h *thriftHandler) getService(key string) IService {
	serviceMap := h.srv.services

	s, ok := serviceMap[key]
	if ok {
		return s
	}

	return nil
}

//Stop ..
func (s *ServerThrift) Stop() {
	if atomic.CompareAndSwapInt32(&s.run, 1, 0) {
		s.thriftserver.Stop()
		logsagent.Info("server " + s.Name() + " set stopch to true")
		s.stopch <- true
		close(s.stopch)
	}
}

//Stopch ..
func (s *ServerThrift) Stopch() <-chan bool {
	return s.stopch
}

//RegisterService ..
func (s *ServerThrift) RegisterService(sv IService) {
	reflectVal := reflect.ValueOf(sv)
	className := reflect.Indirect(reflectVal).Type().Name()

	//方法
	rt := reflectVal.Type()
	for i := 0; i < rt.NumMethod(); i++ {
		m := rt.Method(i)
		c := reflect.ValueOf(sv)
		method := c.MethodByName(m.Name)
		sv.RegisterMethod(m.Name, method)
	}

	//路径
	f1 := className[:1]
	f2 := className[1:]
	key := strings.ToLower(f1) + f2

	//注册
	s.services[key] = sv
	logsagent.Info("service register success in " + s.Name() + " : " + key)
}

func genRequestResponseThrift(ctx context.Context, cfg ServerThriftConfig, form map[string]interface{}) *RequestResponseThrift {
	rs := &RequestResponseThrift{}
	rs.ctx = ctx
	rs.cfg = cfg
	rs.form = form

	rs.packageResult = rs.getPackageResultParam(true)

	return rs
}

//RequestResponseThrift ..
type RequestResponseThrift struct {
	cfg           ServerThriftConfig
	ctx           context.Context
	form          map[string]interface{}
	packageResult bool //是否将结果用JsonServiceResult封装
}

//GetStringFromForm 从Form中获取一个字符串
func (rs *RequestResponseThrift) GetStringFromForm(paramName string) string {
	v, ok := rs.form[paramName]
	if !ok {
		return ""
	}

	return fmt.Sprintf("%v", v)
}

//BeforeHandle ..
func (rs *RequestResponseThrift) BeforeHandle() {
}

//AfterHandle ..
func (rs *RequestResponseThrift) AfterHandle() {
}

//GetSession ..
func (rs *RequestResponseThrift) GetSession(forceSessionid string) session.Store {
	sessionid := forceSessionid
	if sessionid == "" {
		sessionid = rs.GetSessionID()
	}
	sess, err := session.GlobalSessions.GetSessionStore(sessionid)
	throw.CheckErr(err)
	return sess
}

//GetServiceName ..
func (rs *RequestResponseThrift) GetServiceName() string {
	v := rs.GetStringFromForm("service")
	if v == "" {
		panic("no parameter service")
	}
	return v
}

//GetMethodName ..
func (rs *RequestResponseThrift) GetMethodName() string {
	v := rs.GetStringFromForm("m")
	if v == "" {
		panic("no parameter m")
	}

	f1 := v[:1]
	f2 := v[1:]
	return strings.ToUpper(f1) + f2 //方法名第1个字母要大写
}

//PrepareResult ..
func (rs *RequestResponseThrift) PrepareResult(r interface{}) interface{} {
	if obj.InterfaceIsNil(r) {
		if rs.packageResult {
			return JSONServiceResult{State: 0}
		}
		return nil
	}

	result := rs.prepareResult(r)
	return result
}

//准备结果
func (rs *RequestResponseThrift) prepareResult(object interface{}) interface{} {
	if !rs.packageResult {
		return object
	}

	switch t := object.(type) {
	case JSONServiceResult:
		return t
	default:
		jsr := JSONServiceResult{State: 0, Data: object}
		return jsr
	}
}

//GetPackageResultParam 获取是否将结果封装为JsonServiceResult
func (rs *RequestResponseThrift) getPackageResultParam(defaultValue bool) bool {
	packageResult := rs.GetStringFromForm("packageResult")

	if packageResult == "" {
		return defaultValue
	}

	if packageResult != "" && strings.ToUpper(packageResult) == "TRUE" {
		return true
	}

	return false
}

//GetSessionID sessionid由客户端指定，参数名sessionId，实际是addr+sessionid组合，目的是避免sessionid泄露被劫持
//不用remoteAddr作为sessionid，可能被重复使用
func (rs *RequestResponseThrift) GetSessionID() string {
	sessionid := rs.GetStringFromForm("sessionId")
	if sessionid == "" {
		panic("no sessionId")
	}

	return sessionid
}

//GetRemoteIP ..
func (rs *RequestResponseThrift) GetRemoteIP() string {
	return ""
}

//GetRemotePort ..
func (rs *RequestResponseThrift) GetRemotePort() string {
	return ""
}
