// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/ie310mu/ie310go/forks/github.com/ethereum/go-ethereum/log"
)

var (
	contextType      = reflect.TypeOf((*context.Context)(nil)).Elem() //上下方类型定义 context.Context
	errorType        = reflect.TypeOf((*error)(nil)).Elem()           //错误类型定义 error
	subscriptionType = reflect.TypeOf(Subscription{})                 //订阅类型定义，结构体Subscription
	stringType       = reflect.TypeOf("")
)

//记录所有注册的方法和类
type serviceRegistry struct {
	mu       sync.Mutex
	services map[string]service
}

// service represents a registered object.
//注册了的一个服务
type service struct {
	name          string               // name for service  //服务名称
	callbacks     map[string]*callback // registered handlers  //服务的回调
	subscriptions map[string]*callback // available subscriptions/notifications
}

// callback is a method callback which was registered in the server
//回调函数
type callback struct {
	fn          reflect.Value  // the function  方法
	rcvr        reflect.Value  // receiver object of method, set if fn is method  //接收者，方法的第1个参数，对象本身
	argTypes    []reflect.Type // input argument types  //输入参数
	hasCtx      bool           // method's first argument is a context (not included in argTypes)  //是否有上下文参数
	errPos      int            // err return idx, of -1 when method cannot return error  //返回数据中error的序号，-1时表示不返回error
	isSubscribe bool           // true if this is a subscription callback  //是否支持发布订阅
}

//注册服务  rcvr是一个结构体对象
func (r *serviceRegistry) registerName(name string, rcvr interface{}) error {
	rcvrVal := reflect.ValueOf(rcvr)
	if name == "" { //服务名不能为空
		return fmt.Errorf("no service name for type %s", rcvrVal.Type().String())
	}
	callbacks := suitableCallbacks(rcvrVal) //
	if len(callbacks) == 0 {
		return fmt.Errorf("service %T doesn't have any suitable methods/subscriptions to expose", rcvr)
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if r.services == nil {
		r.services = make(map[string]service)
	}
	svc, ok := r.services[name]
	if !ok {
		svc = service{
			name:          name,
			callbacks:     make(map[string]*callback),
			subscriptions: make(map[string]*callback),
		}
		r.services[name] = svc
	}
	for name, cb := range callbacks {
		if cb.isSubscribe {
			svc.subscriptions[name] = cb
		} else {
			svc.callbacks[name] = cb
		}
	}
	return nil
}

// callback returns the callback corresponding to the given RPC method name.
func (r *serviceRegistry) callback(method string) *callback {
	elem := strings.SplitN(method, serviceMethodSeparator, 2)
	if len(elem) != 2 {
		return nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.services[elem[0]].callbacks[elem[1]]
}

// subscription returns a subscription callback in the given service.
func (r *serviceRegistry) subscription(service, name string) *callback {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.services[service].subscriptions[name]
}

// suitableCallbacks iterates over the methods of the given type. It determines if a method
// satisfies the criteria for a RPC callback or a subscription callback and adds it to the
// collection of callbacks. See server documentation for a summary of these criteria.
//查找所有可供注册的函数
func suitableCallbacks(receiver reflect.Value) map[string]*callback {
	typ := receiver.Type()
	callbacks := make(map[string]*callback) //创建callback map
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		if method.PkgPath != "" { //PkgPath为空是大写开关的方法，即导出的方法
			continue // method not exported
		}
		cb := newCallback(receiver, method.Func) //创建callback
		if cb == nil {
			continue // function invalid
		}
		name := formatName(method.Name) //第1个字符转为小写
		callbacks[name] = cb            //注册callback
	}
	return callbacks
}

// newCallback turns fn (a function) into a callback object. It returns nil if the function
// is unsuitable as an RPC callback.
//生成一个回调，不符合规则时返回nil
//规则为：方法必须导出
//             方法有3种返回值的定义：响应、错误、响应+错误
//             方法参数必须导出或是内置类型
//             方法返回值必须导出或是内置类型
func newCallback(receiver, fn reflect.Value) *callback {
	fntype := fn.Type()
	c := &callback{fn: fn, rcvr: receiver, errPos: -1, isSubscribe: isPubSub(fntype)} //生成一个callback
	// Determine parameter types. They must all be exported or builtin types.
	c.makeArgTypes()                       //解析输入参数：第1个对象本身、第2个上下文（如果有）、其他参数
	if !allExportedOrBuiltin(c.argTypes) { //判断所有类型是否是导出的或内建的
		return nil
	}
	// Verify return types. The function must return at most one error
	// and/or one other non-error value.
	//检查返回值
	outs := make([]reflect.Type, fntype.NumOut())
	for i := 0; i < fntype.NumOut(); i++ {
		outs[i] = fntype.Out(i)
	}
	if len(outs) > 2 || !allExportedOrBuiltin(outs) { //返回值数量大于2，或者没有导出、不是内建，不符合要求
		return nil
	}
	// If an error is returned, it must be the last returned value.
	switch {
	case len(outs) == 1 && isErrorType(outs[0]): //1个参数、第1个参数是error    (1个参数不是error的情况不用处理，errPos默认=-1)
		c.errPos = 0
	case len(outs) == 2:
		if isErrorType(outs[0]) || !isErrorType(outs[1]) { //2个参数、第1个参数是error，第2个不是error，不符合要求
			return nil
		}
		c.errPos = 1 //第2个参数是error
	}
	return c
}

// makeArgTypes composes the argTypes list.
//解析输入参数：第1个对象本身、第2个上下文（如果有）、其他参数
func (c *callback) makeArgTypes() {
	fntype := c.fn.Type()
	// Skip receiver and context.Context parameter (if present).
	firstArg := 0
	if c.rcvr.IsValid() {
		firstArg++
	}
	if fntype.NumIn() > firstArg && fntype.In(firstArg) == contextType { //第2个参数是不是上下文参数（第1个是对象本身）
		c.hasCtx = true
		firstArg++
	}
	// Add all remaining parameters.
	c.argTypes = make([]reflect.Type, fntype.NumIn()-firstArg)
	for i := firstArg; i < fntype.NumIn(); i++ {
		c.argTypes[i-firstArg] = fntype.In(i)
	}
}

// call invokes the callback.
func (c *callback) call(ctx context.Context, method string, args []reflect.Value) (res interface{}, errRes error) {
	// Create the argument slice.
	fullargs := make([]reflect.Value, 0, 2+len(args))
	if c.rcvr.IsValid() {
		fullargs = append(fullargs, c.rcvr)
	}
	if c.hasCtx {
		fullargs = append(fullargs, reflect.ValueOf(ctx))
	}
	fullargs = append(fullargs, args...)

	// Catch panic while running the callback.
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Error("RPC method " + method + " crashed: " + fmt.Sprintf("%v\n%s", err, buf))
			errRes = errors.New("method handler crashed")
		}
	}()
	// Run the callback.
	results := c.fn.Call(fullargs)
	if len(results) == 0 {
		return nil, nil
	}
	if c.errPos >= 0 && !results[c.errPos].IsNil() {
		// Method has returned non-nil error value.
		err := results[c.errPos].Interface().(error)
		return reflect.Value{}, err
	}
	return results[0].Interface(), nil
}

// Is this an exported - upper case - name?
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

// Are all those types exported or built-in?
//判断所有类型是否是导出的或内建的
func allExportedOrBuiltin(types []reflect.Type) bool {
	for _, typ := range types {
		for typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		// PkgPath will be non-empty even for an exported type,
		// so we need to check the type name as well.
		if !isExported(typ.Name()) && typ.PkgPath() != "" {
			return false
		}
	}
	return true
}

// Is t context.Context or *context.Context?
//是否是上下文类型
func isContextType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t == contextType //context.Context
}

// Does t satisfy the error interface?
//判断类型是否是error
func isErrorType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr { //如果是指针，遍历取到对应的对象
		t = t.Elem()
	}
	return t.Implements(errorType) //看是否实现了errorType
}

// Is t Subscription or *Subscription?
//判断是否是订阅类型
func isSubscriptionType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr { //如果是指针，遍历取到对应的对象
		t = t.Elem()
	}
	return t == subscriptionType //看是否是订阅类型  结构体Subscription
}

// isPubSub tests whether the given method has as as first argument a context.Context and
// returns the pair (Subscription, error).
//判断方法是否支持发布订阅
func isPubSub(methodType reflect.Type) bool {
	// numIn(0) is the receiver type
	if methodType.NumIn() < 2 || methodType.NumOut() != 2 { //参数个数<2、返回值个数!=2，都不支持发布订阅
		return false
	}
	return isContextType(methodType.In(1)) && //第2个数参数是否是context  context.Context
		isSubscriptionType(methodType.Out(0)) && //第1个返回值是订阅类型
		isErrorType(methodType.Out(1)) //第2个返回值是error
}

// formatName converts to first character of name to lowercase.
//第1个字符转为小写
func formatName(name string) string {
	ret := []rune(name)
	if len(ret) > 0 {
		ret[0] = unicode.ToLower(ret[0])
	}
	return string(ret)
}
