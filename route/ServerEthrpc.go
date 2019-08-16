package route

import (
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"
	"strings"
	"sync/atomic"

	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/common/throw"
	"github.com/ie310mu/ie310go/forks/github.com/ethereum/go-ethereum/rpc"
	"github.com/ie310mu/ie310go/logs"
	"github.com/ie310mu/ie310go/session"
)

const (
	//TCP ..
	TCP = "TCP"

	//IPC 命令管道
	IPC = "IPC"

	//WS WebSocket
	WS = "WS"
)

var defaultServeTCPConfig = ServerEthrpcConfig{PortOrName: "8013", Type: TCP}
var defaultServeIpcConfig = ServerEthrpcConfig{PortOrName: `\\.\pipe\ie310`, Type: IPC}
var defaultServeWsConfig = ServerEthrpcConfig{PortOrName: "8023", Type: WS}

//ServerEthrpcConfig ..
type ServerEthrpcConfig struct {
	PortOrName                 string `json:"portOrName,omitempty"`
	Type                       string `json:"type,omitempty"`
	SessionidWithoutRemoteAddr bool   //获取sessionid时，不在前面附加RemoteAddr
}

//NewServerEthrpcDefaultTCP ..
func NewServerEthrpcDefaultTCP() *ServerEthrpc {
	return &ServerEthrpc{services: make(map[string]IService), config: defaultServeTCPConfig, name: "defaultTcpServer", stopch: make(chan bool)}
}

//NewServerEthrpcDefaultIpc ..
func NewServerEthrpcDefaultIpc() *ServerEthrpc {
	return &ServerEthrpc{services: make(map[string]IService), config: defaultServeIpcConfig, name: "defaultIpcServer", stopch: make(chan bool)}
}

//NewServerEthrpcDefaultWs ..
func NewServerEthrpcDefaultWs() *ServerEthrpc {
	return &ServerEthrpc{services: make(map[string]IService), config: defaultServeWsConfig, name: "defaultWsServer", stopch: make(chan bool)}
}

//NewServerEthrpc ..
func NewServerEthrpc(cfg ServerEthrpcConfig, n string) *ServerEthrpc {
	return &ServerEthrpc{services: make(map[string]IService), config: cfg, name: n, stopch: make(chan bool)}
}

//ServerEthrpc ..
type ServerEthrpc struct {
	name      string
	config    ServerEthrpcConfig
	run       int32 //==1表示正在运行中
	stopch    chan bool
	services  map[string]IService
	ethserver *rpc.Server
}

//Name ..
func (s *ServerEthrpc) Name() string {
	return s.name
}

//Run ..
func (s *ServerEthrpc) Run() {
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
	logsagent.Info("server " + s.Name() + " where listen at " + s.config.PortOrName)

	s.ethserver = s.newEthEthrpcServer()
	defer s.ethserver.Stop()
	listener, err := s.genListener()
	throw.CheckErr(err)
	defer listener.Close()

	var c chan int = make(chan int)
	go s.startListen(listener, c)
	<-c
}

func (s *ServerEthrpc) genListener() (net.Listener, error) {
	switch s.config.Type {
	case TCP:
		return net.Listen("tcp", ":"+s.config.PortOrName)
	case IPC:
		return rpc.IpcListen(s.config.PortOrName)
	case WS:
		return net.Listen("tcp", ":"+s.config.PortOrName)
	default:
		return nil, errors.New("error type")
	}
}

func (s *ServerEthrpc) startListen(listener net.Listener, c chan int) {
	defer func() {
		c <- 1
	}()

	switch s.config.Type {
	case TCP:
		s.ethserver.ServeListener(listener) //这个其实一直没有信号返回，因为listener.Accept在没有连接时会一直阻塞
	case IPC:
		s.ethserver.ServeListener(listener)
	case WS:
		rpc.NewWSServer([]string{"*"}, s.ethserver).Serve(listener)
	default:
	}
}

func (s *ServerEthrpc) newEthEthrpcServer() *rpc.Server {
	server := rpc.NewServer()
	//server.idgen = rpc.SequentialIDGenerator()
	err := server.RegisterName("ethrpc", &ethrpcHandler{srv: s})
	throw.CheckErr(err)
	return server
}

type ethrpcHandler struct {
	srv *ServerEthrpc
}

func (h *ethrpcHandler) Invoke(ctx context.Context, kv map[string]interface{}) (r interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught in server "+h.srv.Name(), err)
			logsagent.Error(err)
			r = JSONServiceResult{State: 1, Message: fmt.Sprintf("%v", err)}
		}
	}()

	r = h.invoke(ctx, kv)
	return
}

func (h *ethrpcHandler) invoke(ctx context.Context, kv map[string]interface{}) interface{} {
	rs := genRequestResponseEthrpc(ctx, h.srv.config, kv)
	serviceName := rs.GetServiceName()
	ss := h.getService(serviceName)
	if ss == nil {
		return JSONServiceResult{State: 1, Message: "servicePath not exists"}
	}

	r := ss.HandleRequest(rs)
	return r
}

//Sub ..
//有客户端订阅一次这个函数就会执行一次
func (h *ethrpcHandler) Sub(ctx context.Context, kv map[string]interface{}) (sub *rpc.Subscription, e error) {
	//订阅
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return nil, rpc.ErrNotificationsUnsupported
	}
	subscription := notifier.CreateSubscription()
	idstr := string(subscription.ID)

	//异常处理
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught in server "+h.srv.Name(), err)
			sub = nil
			e = fmt.Errorf("%v", err)
		}
	}()

	//解析要订阅哪个方法
	nip, err := h.sub(ctx, kv)
	throw.CheckErr(err)

	//启动NotifierInProcess
	ss := make(chan interface{}, 10)
	nip.AddChan(idstr, ss)
	go func(nip NotifierInProcess) {
		defer func() {
			if err := recover(); err != nil {
				logsagent.Error("Exception has been caught when run "+nip.Name(), err)
			}
		}()
		nip.TryRun()
	}(nip)

	//处理消息
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logsagent.Error("Exception has been caught in server "+h.srv.Name(), err)
			}
		}()

		for {
			select {
			case <-notifier.Closed():
				nip.RemoveChan(idstr, ss)
				close(ss)
				logs.Info("client " + idstr + " closed in " + nip.Name())
				goto outfor
			case err := <-subscription.Err():
				nip.RemoveChan(idstr, ss)
				close(ss)
				logs.Info("client " + idstr + " err in " + nip.Name() + "：" + fmt.Sprintf("%v", err))
				goto outfor
			case msg := <-ss:
				logs.Info(fmt.Sprintf("%v", msg) + "  -->  " + idstr)
				notifier.Notify(subscription.ID, msg)
			}
		}
	outfor:
	}()

	//返回
	return subscription, nil
}

func (h *ethrpcHandler) sub(ctx context.Context, kv map[string]interface{}) (nip NotifierInProcess, e error) {
	r := h.invoke(ctx, kv)
	if obj.InterfaceIsNil(r) {
		nip = nil
		e = errors.New("method canot be subscribe:  result is nil")
		return
	}

	msg, ok := r.(JSONServiceResult)
	if !ok {
		nip = nil
		e = errors.New("method canot be subscribe:  result is not JSONServiceResult")
		return
	}

	if msg.State != 0 {
		nip = nil
		e = errors.New("method canot be subscribe: " + msg.Message)
		return
	}

	nip, ok2 := msg.Data.(NotifierInProcess)
	if !ok2 {
		nip = nil
		e = errors.New("method canot be subscribe:  result is not NotifierInProcess")
		return
	}

	return nip, nil
}

//getService 查找服务，未找到时返回nil
func (h *ethrpcHandler) getService(key string) IService {
	serviceMap := h.srv.services

	s, ok := serviceMap[key]
	if ok {
		return s
	}

	return nil
}

//Stop ..
func (s *ServerEthrpc) Stop() {
	if atomic.CompareAndSwapInt32(&s.run, 1, 0) {
		s.ethserver.Stop()
		logsagent.Info("server " + s.Name() + " set stopch to true")
		s.stopch <- true
		close(s.stopch)
	}
}

//Stopch ..
func (s *ServerEthrpc) Stopch() <-chan bool {
	return s.stopch
}

//RegisterService ..
func (s *ServerEthrpc) RegisterService(sv IService) {
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

func genRequestResponseEthrpc(ctx context.Context, cfg ServerEthrpcConfig, form map[string]interface{}) *RequestResponseEthrpc {
	rs := &RequestResponseEthrpc{}
	rs.ctx = ctx
	rs.cfg = cfg
	rs.form = form

	rs.packageResult = rs.getPackageResultParam(true)

	return rs
}

//RequestResponseEthrpc ..
type RequestResponseEthrpc struct {
	cfg           ServerEthrpcConfig
	ctx           context.Context
	form          map[string]interface{}
	packageResult bool //是否将结果用JsonServiceResult封装
}

//GetStringFromForm 从Form中获取一个字符串
func (rs *RequestResponseEthrpc) GetStringFromForm(paramName string) string {
	v, ok := rs.form[paramName]
	if !ok {
		return ""
	}

	return fmt.Sprintf("%v", v)
}

//BeforeHandle ..
func (rs *RequestResponseEthrpc) BeforeHandle() {
}

//AfterHandle ..
func (rs *RequestResponseEthrpc) AfterHandle() {
}

//GetSession ..
func (rs *RequestResponseEthrpc) GetSession(forceSessionid string) session.Store {
	sessionid := forceSessionid
	if sessionid == "" {
		sessionid = rs.GetSessionID()
	}
	sess, err := session.GlobalSessions.GetSessionStore(sessionid)
	throw.CheckErr(err)
	return sess
}

//GetServiceName ..
func (rs *RequestResponseEthrpc) GetServiceName() string {
	v := rs.GetStringFromForm("service")
	if v == "" {
		panic("no parameter service")
	}
	return v
}

//GetMethodName ..
func (rs *RequestResponseEthrpc) GetMethodName() string {
	v := rs.GetStringFromForm("m")
	if v == "" {
		panic("no parameter m")
	}

	f1 := v[:1]
	f2 := v[1:]
	return strings.ToUpper(f1) + f2 //方法名第1个字母要大写
}

//PrepareResult ..
func (rs *RequestResponseEthrpc) PrepareResult(r interface{}) interface{} {
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
func (rs *RequestResponseEthrpc) prepareResult(object interface{}) interface{} {
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
func (rs *RequestResponseEthrpc) getPackageResultParam(defaultValue bool) bool {
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
func (rs *RequestResponseEthrpc) GetSessionID() string {
	sessionid := rs.GetStringFromForm("sessionId")
	if sessionid == "" {
		panic("no sessionId")
	}

	if rs.cfg.SessionidWithoutRemoteAddr { //指定不使用远端地址作为前缀
		return sessionid
	}
	if rs.cfg.Type == IPC { //命名管道没有远程地址
		return sessionid
	}

	c, ok := rpc.ClientFromContext(rs.ctx)
	if !ok {
		panic("no client")
	}
	sessionid = c.RemoteAddr() + "_" + sessionid //主要是tcp附加前缀，ipc没有远程地址，ws如果要与http协议的session同步可能要指定SessionidWithoutRemoteAddr=true
	return sessionid
}

//GetRemoteIP ..
func (rs *RequestResponseEthrpc) GetRemoteIP() string {
	c, ok := rpc.ClientFromContext(rs.ctx)
	if !ok {
		panic("no client")
	}
	ip := c.RemoteAddr()
	if ip == "" {
		panic("get remote ip error")
	}

	i := strings.Index(ip, ":")
	if i < 0 {
		panic("get remote ip error")
	}
	return ip[:i]
}

//GetRemotePort ..
func (rs *RequestResponseEthrpc) GetRemotePort() string {
	c, ok := rpc.ClientFromContext(rs.ctx)
	if !ok {
		panic("no client")
	}
	ip := c.RemoteAddr()
	if ip == "" {
		panic("get remote port error")
	}

	i := strings.Index(ip, ":")
	if i < 0 {
		panic("get remote port error")
	}
	return ip[i+1:]
}
