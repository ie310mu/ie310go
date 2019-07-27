package route

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/session"
)

//BaseService ..
type BaseService struct {
	methodMap map[string]reflect.Value
}

var unRegisterMethodNames = []string{"RegisterMethod", "HandleRequest", "GetSessionItem", "SetSessionItem", "DeleteSessionItem", "DeleteAllSessionItems"}

//RegisterMethod ..
func (s *BaseService) RegisterMethod(name string, m reflect.Value) {
	for _, v := range unRegisterMethodNames {
		if v == name {
			return
		}
	}
	if strings.ToUpper(name) == "Sub" || strings.ToUpper(name) == "UnSub" {
		panic("方法名称" + name + "是保留名，无法使用此名称注册")
	}

	if s.methodMap == nil {
		s.methodMap = map[string]reflect.Value{}
	}
	s.methodMap[name] = m
}

//HandleRequest ..
func (s *BaseService) HandleRequest(rs RequestResponse) interface{} {
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught. ", err)
		}
	}()

	args := &ServiceArgs{rs: rs}
	return s.excute(s, args)
}

//处理
func (s *BaseService) excute(object interface{}, args *ServiceArgs) (r interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught. ", err)

			msg := obj.InterfaceToStr(err)
			jsr := JSONServiceResult{State: 1, Message: msg}
			r = jsr
		}
	}()

	args.rs.BeforeHandle()
	s.excuting(args)
	result := s.invokeMethod(args)
	s.excuted(args)
	args.rs.AfterHandle()

	return result
}

//执行方法
func (s *BaseService) invokeMethod(args *ServiceArgs) interface{} {
	m := args.rs.GetMethodName()

	//查找方法
	var method reflect.Value
	if tmp, ok := s.methodMap[m]; ok {
		method = tmp
	} else {
		panic("method " + m + " not exists")
	}

	msg := "begin invoke method  " + args.rs.GetServiceName() + "." + m
	logsagent.Info(msg)

	pars := []reflect.Value{reflect.ValueOf(args)}
	rs := method.Call(pars)

	var result interface{}
	if len(rs) == 0 {
		result = args.rs.PrepareResult(nil)
	} else {
		r := rs[0]
		rr := r.Interface() //v.Interface()是反射的原始对象  v是原始对象的reflect.Value值
		result = args.rs.PrepareResult(rr)
	}
	return result
}

func (s *BaseService) excuting(args *ServiceArgs) {
	var method reflect.Value
	if tmp, ok := s.methodMap["BeforeExcute"]; ok {
		method = tmp
	} else {
		return
	}

	pars := []reflect.Value{reflect.ValueOf(args)}
	method.Call(pars)
}

func (s *BaseService) excuted(args *ServiceArgs) {
	var method reflect.Value
	if tmp, ok := s.methodMap["AfterExcute"]; ok {
		method = tmp
	} else {
		return
	}

	pars := []reflect.Value{reflect.ValueOf(args)}
	method.Call(pars)
}

//GetSessionItem 从Session中获取值
func (s *BaseService) GetSessionItem(args *ServiceArgs, key string, forceSessionid string) interface{} {
	sess := args.rs.GetSession(forceSessionid)
	if sess == nil {
		return nil
	}

	r := sess.Get(key)
	return r
}

//GetSessionItemString ..
func (s *BaseService) GetSessionItemString(args *ServiceArgs, key string, forceSessionid string) string {
	v := s.GetSessionItem(args, key, forceSessionid)
	if obj.InterfaceIsNil(v) {
		return ""
	}

	switch v.(type) {
	case string:
		return v.(string)
	case []byte:
		data := v.([]byte)
		return string(data)
	default:
		return fmt.Sprint(v)
	}
}

//SetSessionItem 设置值到Session
func (s *BaseService) SetSessionItem(args *ServiceArgs, key string, v interface{}, release bool, forceSessionid string) {
	sess := args.rs.GetSession(forceSessionid)
	if sess == nil {
		return
	}

	sess.Set(key, v)
	if release {
		sess.SessionRelease(nil)
	}
}

//SessionRelease 保存值或者读取值
func (s *BaseService) SessionRelease(sess session.Store) {
	sess.SessionRelease(nil)
}

//DeleteSessionItem 在Session中删除一个项目
func (s *BaseService) DeleteSessionItem(args *ServiceArgs, key string, release bool, forceSessionid string) {
	sess := args.rs.GetSession(forceSessionid)
	if sess == nil {
		return
	}

	sess.Delete(key)
	if release {
		sess.SessionRelease(nil)
	}
}

//DeleteAllSessionItems 在Session中删除所有项目
func (s *BaseService) DeleteAllSessionItems(args *ServiceArgs, release bool, forceSessionid string) {
	sess := args.rs.GetSession(forceSessionid)
	if sess == nil {
		logsagent.Info("sess is nil")
		return
	}
	logsagent.Info("the session id is " + sess.SessionID())

	sess.Flush()
	if release {
		sess.SessionRelease(nil)
	}
}
