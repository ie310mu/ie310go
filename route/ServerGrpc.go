package route

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"strings"
	"sync/atomic"

	"github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/common/throw"
	"github.com/ie310mu/ie310go/session"

	pb "github.com/ie310mu/ie310go/grpc"
	"github.com/ie310mu/ie310go/forks/google.golang.org/grpc"
)

//ServerGrpcConfig ..
type ServerGrpcConfig struct {
	Port string `json:"port,omitempty"`
}

var defaultServeGrpcConfig = ServerGrpcConfig{Port: "8033"}

//NewServerGrpcDefault ..
func NewServerGrpcDefault() *ServerGrpc {
	return &ServerGrpc{services: make(map[string]IService), config: defaultServeGrpcConfig, name: "defaultGrpcServer"}
}

//NewServerGrpc ..
func NewServerGrpc(cfg ServerGrpcConfig, n string) *ServerGrpc {
	return &ServerGrpc{services: make(map[string]IService), config: cfg, name: n}
}

//ServerGrpc ..
type ServerGrpc struct {
	name       string
	config     ServerGrpcConfig
	run        int32 //==1表示正在运行中
	stopch     chan bool
	services   map[string]IService
	grpcserver *grpc.Server
}

//Name ..
func (s *ServerGrpc) Name() string {
	return s.name
}

//Run ..
func (s *ServerGrpc) Run() {
	if atomic.LoadInt32(&s.run) == 1 {
		logsagent.Warn("server " + s.Name() + " is already in running ")
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("server " + s.Name() + " run error： ")
			logsagent.Error(err)
			atomic.StoreInt32(&s.run, 0)
			s.stopch <- false
		}
	}()

	atomic.StoreInt32(&s.run, 1)
	logsagent.Info("server " + s.Name() + " where listen at " + s.config.Port)

	s.grpcserver = s.newGrpcServer()
	defer s.grpcserver.Stop()
	listener, err := s.genListener()
	throw.CheckErr(err)
	defer listener.Close()

	var c chan int
	go s.startListen(listener, c)
	<-c
}

func (s *ServerGrpc) genListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", ":"+s.config.Port)
	return lis, err
}

func (s *ServerGrpc) startListen(listener net.Listener, c chan int) {
	defer func() {
		c <- 1
	}()

	err := s.grpcserver.Serve(listener)
	throw.CheckErr(err)
}

func (s *ServerGrpc) newGrpcServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterIe310GoServcieServer(server, &grpcHandler{srv: s})
	return server
}

type grpcHandler struct {
	srv *ServerGrpc
}

func (h *grpcHandler) Invoke(ctx context.Context, request *pb.GrpcRequest) (r *pb.GrpcResponse, e error) {
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught in server "+h.srv.Name(), err)
			logsagent.Error(err)
			rr := JSONServiceResult{State: 1, Message: fmt.Sprintf("%v", err)}
			data := json.ToJSON(rr)
			r = &pb.GrpcResponse{Data: data}
		}
	}()

	{
		kv := make(map[string]interface{})
		json.FromJSON(request.Data, &kv) //注意传指针
		rr := h.invoke(ctx, kv)
		data := json.ToJSON(rr)
		r = &pb.GrpcResponse{Data: data}
		return
	}
}

func (h *grpcHandler) invoke(ctx context.Context, kv map[string]interface{}) interface{} {
	rs := genRequestResponseGrpc(ctx, h.srv.config, kv)
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
func (h *grpcHandler) getService(key string) IService {
	serviceMap := h.srv.services

	s, ok := serviceMap[key]
	if ok {
		return s
	}

	return nil
}

//Stop ..
func (s *ServerGrpc) Stop() {
	if atomic.CompareAndSwapInt32(&s.run, 1, 0) {
		s.grpcserver.Stop()
		s.stopch <- false
	}
}

//Stopch ..
func (s *ServerGrpc) Stopch() <-chan bool {
	return s.stopch
}

//RegisterService ..
func (s *ServerGrpc) RegisterService(sv IService) {
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

func genRequestResponseGrpc(ctx context.Context, cfg ServerGrpcConfig, form map[string]interface{}) *RequestResponseGrpc {
	rs := &RequestResponseGrpc{}
	rs.ctx = ctx
	rs.cfg = cfg
	rs.form = form

	rs.packageResult = rs.getPackageResultParam(true)

	return rs
}

//RequestResponseGrpc ..
type RequestResponseGrpc struct {
	cfg           ServerGrpcConfig
	ctx           context.Context
	form          map[string]interface{}
	packageResult bool //是否将结果用JsonServiceResult封装
}

//GetStringFromForm 从Form中获取一个字符串
func (rs *RequestResponseGrpc) GetStringFromForm(paramName string) string {
	v, ok := rs.form[paramName]
	if !ok {
		return ""
	}

	return fmt.Sprintf("%v", v)
}

//BeforeHandle ..
func (rs *RequestResponseGrpc) BeforeHandle() {
}

//AfterHandle ..
func (rs *RequestResponseGrpc) AfterHandle() {
}

//GetSession ..
func (rs *RequestResponseGrpc) GetSession(forceSessionid string) session.Store {
	sessionid := forceSessionid
	if sessionid == "" {
		sessionid = rs.GetSessionID()
	}
	sess, err := session.GlobalSessions.GetSessionStore(sessionid)
	throw.CheckErr(err)
	return sess
}

//GetServiceName ..
func (rs *RequestResponseGrpc) GetServiceName() string {
	v := rs.GetStringFromForm("service")
	if v == "" {
		panic("no parameter service")
	}
	return v
}

//GetMethodName ..
func (rs *RequestResponseGrpc) GetMethodName() string {
	v := rs.GetStringFromForm("m")
	if v == "" {
		panic("no parameter m")
	}

	f1 := v[:1]
	f2 := v[1:]
	return strings.ToUpper(f1) + f2 //方法名第1个字母要大写
}

//PrepareResult ..
func (rs *RequestResponseGrpc) PrepareResult(r interface{}) interface{} {
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
func (rs *RequestResponseGrpc) prepareResult(object interface{}) interface{} {
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
func (rs *RequestResponseGrpc) getPackageResultParam(defaultValue bool) bool {
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
func (rs *RequestResponseGrpc) GetSessionID() string {
	sessionid := rs.GetStringFromForm("sessionId")
	if sessionid == "" {
		panic("no sessionId")
	}

	return sessionid
}

//GetRemoteIP ..
func (rs *RequestResponseGrpc) GetRemoteIP() string {
	return ""
}

//GetRemotePort ..
func (rs *RequestResponseGrpc) GetRemotePort() string {
	return ""
}
