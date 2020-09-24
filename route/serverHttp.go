package route

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync/atomic"

	"github.com/ie310mu/ie310go/common/convert"
	"github.com/ie310mu/ie310go/common/guid"
	"github.com/ie310mu/ie310go/common/json"
	jsonutil "github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/common/throw"
	"github.com/ie310mu/ie310go/session"
)

var defaultServeHTTPConfig = ServerHTTPConfig{
	Port:          "8003",
	ServiceSuffix: "goss",
	Jsonp:         true,
}

//ServerHTTPConfig ..
//StaticDirs、DefaultStaticDir建议用绝对目录，如  dir.GetCurrentPath() + "static"
type ServerHTTPConfig struct {
	Port             string `json:"port,omitempty"`
	ServiceSuffix    string `json:"serviceSuffix,omitempty"`
	Jsonp            bool   `json:"jsonp,omitempty"`
	LogArgs          bool   `json:"logArgs,omitempty"`
	ServeStaticFunc  ServeStaticFunc
	StaticDirs       map[string]string `json:"staticDirs,omitempty"`       //以第1个path定位
	DefaultStaticDir string            `json:"defaultStaticDir,omitempty"` //不以path定位，只要不带Service.goss就认为是静态资源
}

//NewServerHTTPDefault ..
func NewServerHTTPDefault() *ServerHTTP {
	return &ServerHTTP{services: make(map[string]IService), config: defaultServeHTTPConfig, name: "defaultHttpServer", stopch: make(chan bool)}
}

//NewServerHTTP ..
func NewServerHTTP(cfg ServerHTTPConfig, n string) *ServerHTTP {
	return &ServerHTTP{services: make(map[string]IService), config: cfg, name: n, stopch: make(chan bool)}
}

//ServerHTTP ..
type ServerHTTP struct {
	name     string
	config   ServerHTTPConfig
	run      int32 //==1表示正在运行中
	stopch   chan bool
	services map[string]IService
}

//Name ..
func (s *ServerHTTP) Name() string {
	return s.name
}

//Run ..
func (s *ServerHTTP) Run() {
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
	err := http.ListenAndServe(":"+s.config.Port, &httpHandler{srv: s}) //这句会阻塞，其实执行不到下一句，why?
	throw.CheckErr(err)
}

//Stop ..
func (s *ServerHTTP) Stop() {
	if atomic.CompareAndSwapInt32(&s.run, 1, 0) {
		//todo.............如何关闭http.ListenAndServe???
		logsagent.Info("server " + s.Name() + " set stopch to true")
		s.stopch <- true
		close(s.stopch)
	}
}

//Stopch ..
func (s *ServerHTTP) Stopch() <-chan bool {
	return s.stopch
}

//RegisterService ..
func (s *ServerHTTP) RegisterService(sv IService) {
	reflectVal := reflect.ValueOf(sv)
	className := reflect.Indirect(reflectVal).Type().Name()

	//方法
	rt := reflectVal.Type()
	for i := 0; i < rt.NumMethod(); i++ {
		m := rt.Method(i)
		method := reflectVal.MethodByName(m.Name)
		sv.RegisterMethod(m.Name, method)
	}

	//路径
	f1 := className[:1]
	f2 := className[1:]
	key := strings.ToLower(f1) + f2
	key = "/" + key + "." + s.config.ServiceSuffix

	//注册
	s.services[key] = sv
	logsagent.Info("service register success in " + s.Name() + " : " + key)
}

type httpHandler struct {
	srv *ServerHTTP
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught in server "+h.srv.Name(), err)
		}
	}()

	if h.srv.config.ServeStaticFunc != nil {
		if h.srv.config.ServeStaticFunc(w, r, h.srv.config) {
			return
		}
	}

	rs := genRequestResponseHTTP(w, r, h.srv.config)
	serviceName := rs.GetServiceName()
	if serviceName == "" {
		h.outputResult(rs, JSONServiceResult{State: 1, Message: "serviceName error"})
		return
	}
	servicePath := "/" + serviceName + "." + h.srv.config.ServiceSuffix

	s := h.getService(servicePath)
	if s == nil {
		h.outputResult(rs, JSONServiceResult{State: 1, Message: "servicePath not exists"})
		return
	}

	m := rs.GetStringFromForm("m")
	m = strings.ToUpper(m)
	if m == "SUB" { //订阅请求
		subid, err := h.Sub(rs, s)
		if err == nil {
			result := JSONServiceResult{State: 0, Data: subid}
			h.outputResult(rs, result)
		} else {
			result := JSONServiceResult{State: 1, Message: fmt.Sprintf("%v", err)}
			h.outputResult(rs, result)
		}
	} else if m == "UNSUB" { //取消订阅
		err := h.UnSub(rs)
		if err == nil {
			result := JSONServiceResult{State: 0, Data: true}
			h.outputResult(rs, result)
		} else {
			result := JSONServiceResult{State: 1, Message: fmt.Sprintf("%v", err)}
			h.outputResult(rs, result)
		}
	} else { //正常请求
		result := s.HandleRequest(rs)
		h.outputResult(rs, result)
	}
}

func (h httpHandler) outputResult(rs *RequestResponseHTTP, result interface{}) {
	//application/json（默认是text/plain会被自动压缩，jquery ajax解析会失败导致timeout）
	if rs.CustomContentType == "" {
		rs.w.Header().Set("Content-Type", "application/json;charset=utf-8")
	} else {
		rs.w.Header().Set("Content-Type", rs.CustomContentType)
	}
	if rs.useJsonp {
		rs.w.Header().Set("Access-Control-Allow-Origin", "*")
		rs.w.Header().Set("Access-Control-Allow-Credentials", "true")
		rs.w.Header().Set("Access-Control-Allow-Methods", "*")
		rs.w.Header().Set("Access-Control-Allow-Headers", "*")
		rs.w.Header().Set("Access-Control-Expose-Headers", "*")
	}

	str := ""
	if result != nil {
		str = rs.resultToJSONStr(result)
	}

	//fmt.Fprintf(args.W, result)  //不能用这句输出，%会变成%MISSING
	rs.w.Write([]byte(str))
}

//getService 查找服务，未找到时返回nil
func (h httpHandler) getService(key string) IService {
	serviceMap := h.srv.services

	s, ok := serviceMap[key]
	if ok {
		return s
	}

	return nil
}

func genRequestResponseHTTP(w http.ResponseWriter, r *http.Request, srvc ServerHTTPConfig) *RequestResponseHTTP {
	rs := &RequestResponseHTTP{}
	rs.cfg = srvc
	rs.w = w
	rs.r = r
	rs.r.ParseForm()

	if rs.cfg.LogArgs {
		logsagent.Info(json.ToJSON(rs.r.Form))
	}

	//args.SecondEncoding = args.GetStringParamWithCheck("secondEncoding", true)
	rs.useJsonp = rs.getUseJsonpParam()
	if rs.useJsonp {
		rs.jsonpCallback = rs.getJsonpCallbackParam()
	}
	rs.convertResult = rs.getConvertResultParam(true)
	rs.packageResult = rs.convertResult && rs.getPackageResultParam(true) //如果不转换为json，肯定是不封装结果的

	return rs
}

//RequestResponseHTTP ..
type RequestResponseHTTP struct {
	cfg               ServerHTTPConfig
	w                 http.ResponseWriter
	r                 *http.Request
	useJsonp          bool   //请求方指定的是否使用jsonp  //主要用于http协议  //只适用于单次请求
	jsonpCallback     string //使用jsonp时jsonpCallback函数名称
	convertResult     bool   //是否转换最终结果为json字符串，为否时直接toString()  默认true，主要用于需要接口直接返回字符串而不加引号时设置为false
	packageResult     bool   //是否将结果用JsonServiceResult封装
	CustomContentType string //由上层调用者自定义接口的返回数据类型
	form              map[string]interface{}
}

func (rs *RequestResponseHTTP) SetArgValue(key string, value string) {
	if rs.form == nil {
		rs.form = make(map[string]interface{})
	}
	rs.form[key] = value //不写到http.Request的Form中，内部机制不明确
}

//GetStringFromForm 从Form中获取一个字符串
func (rs *RequestResponseHTTP) GetStringFromForm(paramName string) string {
	v, ok := rs.form[paramName]
	if ok {
		return fmt.Sprintf("%v", v)
	}

	//logsagent.Info(jsonutil.ToJSON(rs.r.Form))
	vs := rs.r.Form[paramName]
	if vs == nil || len(vs) == 0 || vs[0] == "" {
		vs = rs.r.PostForm[paramName]
	}
	if vs == nil || len(vs) == 0 {
		return ""
	}

	return vs[0]
}

//BeforeHandle ..
func (rs *RequestResponseHTTP) BeforeHandle() {
	//跨域检查，防止csrf攻击
	if !rs.useJsonp {
		referer := rs.r.Referer()
		host := "https"
		if strings.HasPrefix(rs.r.Proto, "HTTP") {
			host = "http"
		}
		host = host + "://" + rs.r.Host
		if referer != "" && !strings.HasPrefix(referer, host) {
			panic("referer error")
		}
	}
}

//AfterHandle ..
func (rs *RequestResponseHTTP) AfterHandle() {
}

//GetSession ..
func (rs *RequestResponseHTTP) GetSession(forceSessionid string) session.Store {
	sess, err := session.GlobalSessions.SessionStart(rs.w, rs.r, forceSessionid)
	throw.CheckErr(err)
	return sess
}

//GetServiceName ..
func (rs *RequestResponseHTTP) GetServiceName() string {
	path := rs.r.URL.Path[1:]
	path = strings.Replace(path, "."+rs.cfg.ServiceSuffix, "", -1)

	if path == "" {
		path = rs.GetStringFromForm("service")
	}

	return path
}

//GetMethodName ..
func (rs *RequestResponseHTTP) GetMethodName() string {
	v := rs.GetStringFromForm("m")
	if v == "" {
		v = rs.GetStringFromForm("method")
	}
	if v == "" {
		panic("no parameter m or method")
	}

	//检查是否是订阅方法
	if strings.ToUpper(v) == "SUB" {
		v = rs.GetStringFromForm("subMethod") //取真实方法名
	}

	f1 := v[:1]
	f2 := v[1:]
	return strings.ToUpper(f1) + f2 //方法名第1个字母要大写
}

//PrepareResult ..
func (rs *RequestResponseHTTP) PrepareResult(r interface{}) interface{} {
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
func (rs *RequestResponseHTTP) prepareResult(object interface{}) interface{} {
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

func (rs *RequestResponseHTTP) resultToJSONStr(result interface{}) string {
	var json string
	if rs.convertResult {
		json = jsonutil.ToJSON(result)
	} else {
		switch t := result.(type) {
		case string:
			json = t
		default:
			//v := reflect.ValueOf(result)
			//json = v.String()
			json = fmt.Sprintf("%v", result)
		}
	}

	if rs.useJsonp {
		if rs.jsonpCallback != "" {
			json = rs.jsonpCallback + "(" + json + ");"
		}
	}
	return json
}

//GetUseJsonpParam 获取是否允许跨域
func (rs *RequestResponseHTTP) getUseJsonpParam() bool {
	if !rs.cfg.Jsonp {
		return false
	}

	useJsonpStr := rs.GetStringFromForm("useJsonp")
	if useJsonpStr != "" && strings.ToUpper(useJsonpStr) == "TRUE" {
		return true
	}

	return false
}

//GetJsonpCallbackParam 获取跨域的回调方法名
func (rs *RequestResponseHTTP) getJsonpCallbackParam() string {
	result := rs.GetStringFromForm("jsonpCallback") //自定义的
	if result == "" {
		result = rs.GetStringFromForm("callback") //jquery的参数名
	}

	//2019.04.21如果没有指定jsonpCallback或callback，则不会进行封装，直接返回json字符串
	//（适应于需要跨域访问但又不需要封装的情况）
	// if result == "" {
	// 	result = "jsonpCallback"
	// }

	return result
}

//GetConvertResultParam 获取是否将结果转换为json
func (rs *RequestResponseHTTP) getConvertResultParam(defaultValue bool) bool {
	convertResult := rs.GetStringFromForm("convertResult")

	if convertResult == "" {
		return defaultValue
	}

	if convertResult != "" && strings.ToUpper(convertResult) == "TRUE" {
		return true
	}

	return false
}

//GetPackageResultParam 获取是否将结果封装为JsonServiceResult
func (rs *RequestResponseHTTP) getPackageResultParam(defaultValue bool) bool {
	packageResult := rs.GetStringFromForm("packageResult")

	if packageResult == "" {
		return defaultValue
	}

	if packageResult != "" && strings.ToUpper(packageResult) == "TRUE" {
		return true
	}

	return false
}

//GetSessionID ..
func (rs *RequestResponseHTTP) GetSessionID() string {
	sess := rs.GetSession("") //不用getSid，那个不会生成sesssionid的header
	return sess.SessionID()

	// if args.sessionid != "" {
	// 	return args.sessionid
	// }

	// args.sessionid = args.GetStringParamWithDefault("sessionId", "")

	// //从Header获取
	// if args.sessionid == "" {
	// 	//注意header中的key会被转换成Content-Length的格式（以-分隔，首字母大写其他小写）
	// 	vs := args.R.Header["Sessionid"]
	// 	if vs != nil && len(vs) != 0 {
	// 		args.sessionid = vs[0]
	// 	}
	// }

	// //从cookie获取
	// // if sessionid == "" {
	// // 	vs := args.R.Header["Cookie"]
	// // 	if vs != nil || len(vs) > 0 {
	// // 		for _, v := range vs {
	// // 			if strings.Index(v, "JSESSIONID=") == 0 {
	// // 				sessionid = strings.Replace(v, "JSESSIONID=", "", -1)
	// // 				break
	// // 			}
	// // 		}
	// // 	}
	// // }
	// c, err := args.R.Cookie("JSESSIONID")
	// if err == nil {
	// 	args.sessionid = c.Value
	// }

	// return args.sessionid
}

//GetRemoteIP ..
func (rs *RequestResponseHTTP) GetRemoteIP() string {
	ip := rs.r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = rs.r.RemoteAddr
	}
	return ip
}

//GetRemotePort ..
func (rs *RequestResponseHTTP) GetRemotePort() string {
	panic("未实现")
}

//GetW ..
func (rs *RequestResponseHTTP) GetW() http.ResponseWriter {
	return rs.w
}

//GetR ..
func (rs *RequestResponseHTTP) GetR() *http.Request {
	return rs.r
}

//有客户端订阅一次这个函数就会执行一次
func (h *httpHandler) Sub(rs *RequestResponseHTTP, s IService) (subid string, e error) {
	//异常处理
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught in server "+h.srv.Name(), err)
			subid = ""
			e = fmt.Errorf("%v", err)
		}
	}()

	//检查参数
	subMethod := rs.GetStringFromForm("subMethod")
	if subMethod == "" {
		return "", errors.New("no parameter subMethod")
	}
	backurl := rs.GetStringFromForm("backurl")
	if backurl == "" {
		return "", errors.New("no parameter backurl")
	}

	//生成一个notifier
	hn := genHttpnotifier(backurl)

	//解析要订阅哪个方法
	nip, err := h.sub(rs, s)
	throw.CheckErr(err)

	//启动NotifierInProcess
	ss := make(chan interface{}, 10)
	nip.AddChan(hn.id, ss)
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
			case <-hn.closed:
				nip.RemoveChan(hn.id, ss)
				close(ss)
				logsagent.Info("client " + hn.id + " closed in " + nip.Name())
				goto outfor
			case msg := <-ss:
				//logsagent.Info(hn.id + ":    " + fmt.Sprintf("%v", msg))
				hn.Notify(msg)
			}
		}
	outfor:
	}()

	//返回
	addHttpnotifier(hn)
	return hn.id, nil
}

func (h *httpHandler) sub(rs *RequestResponseHTTP, s IService) (nip NotifierInProcess, e error) {
	r := s.HandleRequest(rs)
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

//有客户端订阅一次这个函数就会执行一次
func (h *httpHandler) UnSub(rs *RequestResponseHTTP) error {
	subid := rs.GetStringFromForm("subId")
	if subid == "" {
		return errors.New("no parameter subId")
	}

	return removeHttpnotifier(subid)
}

var httpnotifierMap = make(map[string]*httpnotifier)

func addHttpnotifier(hn *httpnotifier) {
	httpnotifierMap[hn.id] = hn
}

func removeHttpnotifier(subid string) error {
	hn, ok := httpnotifierMap[subid]
	if !ok {
		return errors.New("no sub has id " + subid)
	}

	hn.closed <- true
	return nil
}

//HttpnotifierMaxErrorCount http订阅时，允许的最大的回调失败的次数，达到此次数后，订阅将被关闭
var HttpnotifierMaxErrorCount = 5

func genHttpnotifier(backurl string) *httpnotifier {
	hn := &httpnotifier{id: guid.Get(), backurl: backurl, maxErrorCount: HttpnotifierMaxErrorCount}
	hn.closed = make(chan bool, 1)
	return hn
}

type httpnotifier struct {
	id            string
	backurl       string
	closed        chan bool
	errorCount    int
	maxErrorCount int
}

func (hn *httpnotifier) addErrorCount() {
	hn.errorCount = hn.errorCount + 1      //每失败一次记录一次
	if hn.errorCount >= hn.maxErrorCount { //错误超过一定次数则移除之
		logsagent.Warn("http消息发布失败达到" + convert.IntToStr(hn.maxErrorCount) + "次，subid：" + hn.id + "，关闭此订阅")
		removeHttpnotifier(hn.id)
	}
}

func (hn *httpnotifier) Notify(msg interface{}) {
	logsagent.Info(fmt.Sprintf("%v", msg) + "  -->  " + hn.backurl)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logsagent.Warn("http消息发布失败，subid：" + hn.id + "：" + fmt.Sprintf("%v", err))
				hn.addErrorCount()
			}
		}()

		//组装参数
		data := make(url.Values)
		data["datajson"] = []string{jsonutil.ToJSON(msg)}

		resp, err := http.PostForm(hn.backurl, data)
		throw.CheckErr(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		throw.CheckErr(err)
		resultstr := string(body)

		logsagent.Info("<--  " + resultstr)
	}()
}

//DefaultServeStatic ..
func DefaultServeStatic(w http.ResponseWriter, r *http.Request, cfg ServerHTTPConfig) bool {
	//前缀相符
	if cfg.StaticDirs != nil {
		for prefix, staticDir := range cfg.StaticDirs {
			if strings.HasPrefix(r.URL.Path, prefix) {
				file := staticDir + r.URL.Path[len(prefix):]
				http.ServeFile(w, r, file)
				logsagent.Info("serve static : " + r.URL.Path + "  --> " + file)
				return true
			}
		}
	}

	//path中不带Service.goss字样
	if cfg.DefaultStaticDir != "" {
		if strings.Index(r.URL.Path, "Service."+cfg.ServiceSuffix) < 0 {
			file := cfg.DefaultStaticDir + r.URL.Path
			p, f := filepath.Split(file)
			if f == "" {
				indexFile := p + "index.html"
				if fileIsExist(indexFile) {
					http.ServeFile(w, r, file)
					//logsagent.Info("serve static : " + r.URL.Path + "  --> " + file)
				}
			} else {
				http.ServeFile(w, r, file)
				//logsagent.Info("serve static : " + r.URL.Path + "  --> " + file)
			}
			//logsagent.Info(file + " : " + p + ",   " + f)
			return true
		}
	}

	return false
}

func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//ServeStaticFunc 静态文件处理回调
type ServeStaticFunc func(w http.ResponseWriter, r *http.Request, cfg ServerHTTPConfig) bool
