// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logs provide a general log interface
// Usage:
//
// import "github.com/ie310mu/ie310go/logs"
//
//	log := NewLogger(10000)
//	log.SetLogger("console", "")
//
//	> the first params stand for how many channel
//
// Use it like this:
//
//	log.Trace("trace")
//	log.Info("info")
//	log.Warn("warning")
//	log.Debug("debug")
//	log.Critical("critical")
//
//  more docs http://beego.me/docs/module/logs.md
package logs

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
)

// RFC5424 log message levels.
// RFC5424 日志消息级别定义，从0---7
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

// levelLogLogger is defined to implement log.Logger
// the real log level will be LevelEmergency
// 真实输出级别为LevelEmergency
const levelLoggerImpl = -1

// Name for adapter with beego official support
// beego默认支持的适配器
const (
	AdapterConsole   = "console"
	AdapterFile      = "file"
	AdapterMultiFile = "multifile"
	AdapterMail      = "smtp"
	AdapterConn      = "conn"
	AdapterEs        = "es"
	AdapterJianLiao  = "jianliao"
	AdapterSlack     = "slack"
	AdapterAliLS     = "alils"
)

// Legacy log level constants to ensure backwards compatibility.
// 确保向后兼容性的旧日志级别常量
const (
	LevelInfo  = LevelInformational
	LevelTrace = LevelDebug
	LevelWarn  = LevelWarning
)

// Logger生成器定义（一个函数）
type newLoggerFunc func() Logger

// Logger defines the behavior of a log provider.
// Logger接口定义
type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

// 适配器map
var adapters = make(map[string]newLoggerFunc)

//不同级别的日志信息的标识字符
var levelPrefix = [LevelDebug + 1]string{"[M] ", "[A] ", "[C] ", "[E] ", "[W] ", "[N] ", "[I] ", "[D] "}

// Register makes a log provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
// 注册一个Logger生成器
func Register(name string, log newLoggerFunc) {
	if log == nil {
		panic("logs: Register provide is nil")
	}
	if _, dup := adapters[name]; dup {
		panic("logs: Register called twice for provider " + name)
	}
	adapters[name] = log
}

// BeeLogger is default logger in beego application.
// it can contain several providers and log message into all providers.
// BeeLogger结构体定义，日志管理器
type BeeLogger struct {
	lock                sync.Mutex     //用于关键操作时加锁（如初始化、设置异步等）
	level               int            //可输出的日志级别（大于此级别的日志不输出，默认值是7，即输出所有信息）
	init                bool           //标识是否已经初始化
	enableFuncCallDepth bool           //是否显示函数调用堆栈
	loggerFuncCallDepth int            //显示函数调用堆栈时，显示几层
	asynchronous        bool           //是否异步输出日志
	prefix              string         //日志消息的前缀
	msgChanLen          int64          //异步输出消息时，缓冲chan中可存储的数量，默认值1000
	msgChan             chan *logMsg   //异步输出消息的缓冲chan
	signalChan          chan string    //信号chan，用于flush或close
	wg                  sync.WaitGroup //等待组，用于在flush或close时正确结束？？？？
	outputs             []*nameLogger  //真实的日志输出模块
	enableErrorStack    bool           //是否输出error的堆栈信息
}

const defaultAsyncMsgLen = 1e3 //默认的msgChanLen=1000

// name--Logger结构体
type nameLogger struct {
	Logger
	name string
}

//日志内容结构体，包含级别、消息、时间3要素
type logMsg struct {
	level int
	msg   string
	when  time.Time
}

//logMsg结构池，因为logMsg结构会大量地生成，用Pool可加快速度
var logMsgPool *sync.Pool

// NewLogger returns a new BeeLogger.
// channelLen means the number of messages in chan(used where asynchronous is true).
// if the buffering chan is full, logger adapters write to file or other way.
// 返回一个新的NewLogger对象，默认只开启控制台日志
func NewLogger(channelLens ...int64) *BeeLogger {
	bl := new(BeeLogger)
	bl.level = LevelDebug
	bl.loggerFuncCallDepth = 2
	bl.msgChanLen = append(channelLens, 0)[0]
	if bl.msgChanLen <= 0 {
		bl.msgChanLen = defaultAsyncMsgLen
	}
	bl.signalChan = make(chan string, 1)
	bl.setLogger(AdapterConsole)
	return bl
}

// Async set the log to asynchronous and start the goroutine
// 异步输出日志，msgLen可指定缓冲的大小，不指定时默认值为1000
func (bl *BeeLogger) Async(msgLen ...int64) *BeeLogger {
	bl.lock.Lock()         //先加锁
	defer bl.lock.Unlock() //处理完后解锁
	if bl.asynchronous {   //如果已经是异步，不处理
		return bl
	}
	bl.asynchronous = true
	if len(msgLen) > 0 && msgLen[0] > 0 {
		bl.msgChanLen = msgLen[0]
	}
	bl.msgChan = make(chan *logMsg, bl.msgChanLen) //初始化消息chan
	logMsgPool = &sync.Pool{                       //初始化消息池
		New: func() interface{} {
			return &logMsg{}
		},
	}
	bl.wg.Add(1)        //waitGroup加1，需要在日志输出协程读取到flush或close信号时减1，这样flush或close方法才能退出
	go bl.startLogger() //启动日志输出协程
	return bl
}

// SetLogger provides a given logger adapter into BeeLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
//设置日志输出的方式（此为内部方法）
func (bl *BeeLogger) setLogger(adapterName string, configs ...string) error {
	config := append(configs, "{}")[0]
	for _, l := range bl.outputs {
		if l.name == adapterName {
			return fmt.Errorf("logs: duplicate adaptername %q (you have set this logger before)", adapterName)
		}
	}

	log, ok := adapters[adapterName]
	if !ok {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}

	lg := log()
	err := lg.Init(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs.BeeLogger.SetLogger: "+err.Error())
		return err
	}
	bl.outputs = append(bl.outputs, &nameLogger{name: adapterName, Logger: lg})
	return nil
}

// SetLogger provides a given logger adapter into BeeLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
//设置日志输出的方式，可多次调用
func (bl *BeeLogger) SetLogger(adapterName string, configs ...string) error {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	if !bl.init {
		bl.outputs = []*nameLogger{}
		bl.init = true
	}
	return bl.setLogger(adapterName, configs...)
}

// DelLogger remove a logger adapter in BeeLogger.
//删除某个输出方式
func (bl *BeeLogger) DelLogger(adapterName string) error {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	outputs := []*nameLogger{}
	for _, lg := range bl.outputs {
		if lg.name == adapterName {
			lg.Destroy()
		} else {
			outputs = append(outputs, lg)
		}
	}
	if len(outputs) == len(bl.outputs) {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}
	bl.outputs = outputs
	return nil
}

//将某个日志输出到outputs
func (bl *BeeLogger) writeToLoggers(when time.Time, msg string, level int) {
	for _, l := range bl.outputs {
		err := l.WriteMsg(when, msg, level)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to WriteMsg to adapter:%v,error:%v\n", l.name, err)
		}
	}
}

//输出一个消息，不指定级别使用默认的-1作为级别，肯定会被输出
func (bl *BeeLogger) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	// writeMsg will always add a '\n' character
	if p[len(p)-1] == '\n' {
		p = p[0 : len(p)-1]
	}
	// set levelLoggerImpl to ensure all log message will be write out
	err = bl.writeMsg(levelLoggerImpl, string(p))
	if err == nil {
		return len(p), err
	}
	return 0, err
}

//指定级别，输出消息
//msg中可以包含fmt.PrintLn中的格式化符号%，v...则是对应的值
func (bl *BeeLogger) writeMsg(logLevel int, msg string, v ...interface{}) error {
	if !bl.init {
		bl.lock.Lock()
		bl.setLogger(AdapterConsole)
		bl.lock.Unlock()
	}

	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}

	msg = bl.prefix + " " + msg

	when := time.Now()
	if bl.enableFuncCallDepth {
		_, file, line, ok := runtime.Caller(bl.loggerFuncCallDepth)
		if !ok {
			file = "???"
			line = 0
		}
		_, filename := path.Split(file)
		msg = "[" + filename + ":" + strconv.Itoa(line) + "] " + msg
	}

	//set level info in front of filename info
	if logLevel == levelLoggerImpl {
		// set to emergency to ensure all log will be print out correctly
		logLevel = LevelEmergency
	} else {
		msg = levelPrefix[logLevel] + msg
	}

	if bl.asynchronous {
		lm := logMsgPool.Get().(*logMsg)
		lm.level = logLevel
		lm.msg = msg
		lm.when = when
		bl.msgChan <- lm
	} else {
		bl.writeToLoggers(when, msg, logLevel)
	}
	return nil
}

// SetLevel Set log message level.
// If message level (such as LevelDebug) is higher than logger level (such as LevelWarning),
// log providers will not even be sent the message.
//设置输出级别，高于此级别的日志不会被输出
func (bl *BeeLogger) SetLevel(l int) {
	bl.level = l
}

// GetLevel Get Current log message level.
//获取输出级别
func (bl *BeeLogger) GetLevel() int {
	return bl.level
}

// SetLogFuncCallDepth set log funcCallDepth
//设置函数调用堆栈的显示层级
func (bl *BeeLogger) SetLogFuncCallDepth(d int) {
	bl.loggerFuncCallDepth = d
}

// GetLogFuncCallDepth return log funcCallDepth for wrapper
func (bl *BeeLogger) GetLogFuncCallDepth() int {
	return bl.loggerFuncCallDepth
}

// EnableFuncCallDepth enable log funcCallDepth
//是否显示函数调用堆栈
func (bl *BeeLogger) EnableFuncCallDepth(b bool) {
	bl.enableFuncCallDepth = b
}

// EnableErrorStack ..
func (bl *BeeLogger) EnableErrorStack(b bool) {
	bl.enableErrorStack = b
}

// SetPrefix set prefix
// 设置消息前缀
func (bl *BeeLogger) SetPrefix(s string) {
	bl.prefix = s
}

// start logger chan reading.
// when chan is not empty, write logs.
//日志输出协程
func (bl *BeeLogger) startLogger() {
	gameOver := false //用于标识是否需要退出循环
	for {             //无限循环
		select { //从chan中获取数据
		case bm := <-bl.msgChan: //获取到1个日志数据
			bl.writeToLoggers(bm.when, bm.msg, bm.level)
			logMsgPool.Put(bm)
		case sg := <-bl.signalChan: //获取到一个信号量，可能是flush或close
			// Now should only send "flush" or "close" to bl.signalChan
			bl.flush()         //先缓冲输出所有日志
			if sg == "close" { //如果是close信号，代表BeeLogger被关闭，需要关闭所有的output，并且标识退出循环
				for _, l := range bl.outputs {
					l.Destroy()
				}
				bl.outputs = nil
				gameOver = true
			}
			bl.wg.Done() //表示异步处理结束？？？
		}
		if gameOver {
			break
		}
	}
}

// Emergency Log EMERGENCY level message.
func (bl *BeeLogger) Emergency(format string, v ...interface{}) {
	if LevelEmergency > bl.level {
		return
	}
	bl.writeMsg(LevelEmergency, format, v...)
}

// Alert Log ALERT level message.
func (bl *BeeLogger) Alert(format string, v ...interface{}) {
	if LevelAlert > bl.level {
		return
	}
	bl.writeMsg(LevelAlert, format, v...)
}

// Critical Log CRITICAL level message.
func (bl *BeeLogger) Critical(format string, v ...interface{}) {
	if LevelCritical > bl.level {
		return
	}
	bl.writeMsg(LevelCritical, format, v...)
}

// Error Log ERROR level message.
func (bl *BeeLogger) Error(format string, v ...interface{}) {
	if LevelError > bl.level {
		return
	}
	bl.writeMsg(LevelError, format, v...)

	if bl.enableErrorStack {
		stack := debug.Stack()
		stackString := string(stack)
		bl.writeMsg(LevelError, stackString)
	}
}

// Warning Log WARNING level message.
func (bl *BeeLogger) Warning(format string, v ...interface{}) {
	if LevelWarn > bl.level {
		return
	}
	bl.writeMsg(LevelWarn, format, v...)
}

// Notice Log NOTICE level message.
func (bl *BeeLogger) Notice(format string, v ...interface{}) {
	if LevelNotice > bl.level {
		return
	}
	bl.writeMsg(LevelNotice, format, v...)
}

// Informational Log INFORMATIONAL level message.
func (bl *BeeLogger) Informational(format string, v ...interface{}) {
	if LevelInfo > bl.level {
		return
	}
	bl.writeMsg(LevelInfo, format, v...)
}

// Debug Log DEBUG level message.
func (bl *BeeLogger) Debug(format string, v ...interface{}) {
	if LevelDebug > bl.level {
		return
	}
	bl.writeMsg(LevelDebug, format, v...)
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func (bl *BeeLogger) Warn(format string, v ...interface{}) {
	if LevelWarn > bl.level {
		return
	}
	bl.writeMsg(LevelWarn, format, v...)
}

// Info Log INFO level message.
// compatibility alias for Informational()
func (bl *BeeLogger) Info(format string, v ...interface{}) {
	if LevelInfo > bl.level {
		return
	}
	bl.writeMsg(LevelInfo, format, v...)
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func (bl *BeeLogger) Trace(format string, v ...interface{}) {
	if LevelDebug > bl.level {
		return
	}
	bl.writeMsg(LevelDebug, format, v...)
}

// Flush flush all chan data.
// 输出所有在缓冲中的日志
func (bl *BeeLogger) Flush() {
	if bl.asynchronous { //如果是异步，用信号通知startLogger处理
		bl.signalChan <- "flush"
		bl.wg.Wait() //等待startLogger处理完毕
		bl.wg.Add(1) //恢复异步等待的状态
		return
	}

	bl.flush() //不是异步，直接处理
}

// Close close logger, flush all chan data and destroy all adapters in BeeLogger.
// 输出所有在缓冲中的日志，并且关闭BeeLogger
func (bl *BeeLogger) Close() {
	if bl.asynchronous { //如果是异步，用信号通知用信号通知startLogger处理
		bl.signalChan <- "close"
		bl.wg.Wait()      //等待startLogger处理完毕
		close(bl.msgChan) //关闭chan
	} else { //不是异步，直接处理
		bl.flush()
		for _, l := range bl.outputs {
			l.Destroy()
		}
		bl.outputs = nil
	}
	close(bl.signalChan) //关闭chan
}

// Reset close all outputs, and set bl.outputs to nil
//重置：先flush，再清空outputs
func (bl *BeeLogger) Reset() {
	bl.Flush()
	for _, l := range bl.outputs {
		l.Destroy()
	}
	bl.outputs = nil
}

//输出所有日志（缓冲中如果有全部输出），再通知outputs Flush
func (bl *BeeLogger) flush() {
	if bl.asynchronous {
		for {
			if len(bl.msgChan) > 0 {
				bm := <-bl.msgChan
				bl.writeToLoggers(bm.when, bm.msg, bm.level)
				logMsgPool.Put(bm)
				continue
			}
			break
		}
	}
	for _, l := range bl.outputs {
		l.Flush()
	}
}

// beeLogger references the used application logger.
//logs空间中默认使用的BeeLogger对象
var beeLogger = NewLogger()

// GetBeeLogger returns the default BeeLogger
func GetBeeLogger() *BeeLogger {
	return beeLogger
}

var beeLoggerMap = struct {
	sync.RWMutex
	logs map[string]*log.Logger
}{
	logs: map[string]*log.Logger{},
}

// GetLogger returns the default BeeLogger
// 这个方法可通过不用的prefix作为key来获取原生的log.Logger对象，但里面封装的Writer还是beeLogger对象，有何意义？？？？
// 是为了使用不同的前缀来输出日志信息？？？
func GetLogger(prefixes ...string) *log.Logger {
	prefix := append(prefixes, "")[0]
	if prefix != "" {
		prefix = fmt.Sprintf(`[%s] `, strings.ToUpper(prefix))
	}
	beeLoggerMap.RLock()
	l, ok := beeLoggerMap.logs[prefix]
	if ok {
		beeLoggerMap.RUnlock()
		return l
	}

	beeLoggerMap.RUnlock()

	beeLoggerMap.Lock()
	defer beeLoggerMap.Unlock()
	l, ok = beeLoggerMap.logs[prefix]
	if !ok {
		l = log.New(beeLogger, prefix, 0)
		beeLoggerMap.logs[prefix] = l
	}
	return l
}

// Reset will remove all the adapter
func Reset() {
	beeLogger.Reset()
}

// Async set the beelogger with Async mode and hold msglen messages
func Async(msgLen ...int64) *BeeLogger {
	return beeLogger.Async(msgLen...)
}

// SetLevel sets the global log level used by the simple logger.
func SetLevel(l int) {
	beeLogger.SetLevel(l)
}

// SetPrefix sets the prefix
func SetPrefix(s string) {
	beeLogger.SetPrefix(s)
}

// EnableFuncCallDepth enable log funcCallDepth
func EnableFuncCallDepth(b bool) {
	beeLogger.enableFuncCallDepth = b
}

// SetLogFuncCall set the CallDepth, default is 4
func SetLogFuncCall(b bool) {
	beeLogger.EnableFuncCallDepth(b)
	beeLogger.SetLogFuncCallDepth(4)
}

// SetLogFuncCallDepth set log funcCallDepth
func SetLogFuncCallDepth(d int) {
	beeLogger.loggerFuncCallDepth = d
}

// SetLogger sets a new logger.
func SetLogger(adapter string, config ...string) error {
	return beeLogger.SetLogger(adapter, config...)
}

// Emergency logs a message at emergency level.
func Emergency(f interface{}, v ...interface{}) {
	beeLogger.Emergency(formatLog(f, v...))
}

// Alert logs a message at alert level.
func Alert(f interface{}, v ...interface{}) {
	beeLogger.Alert(formatLog(f, v...))
}

// Critical logs a message at critical level.
func Critical(f interface{}, v ...interface{}) {
	beeLogger.Critical(formatLog(f, v...))
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	beeLogger.Error(formatLog(f, v...))
}

// Warning logs a message at warning level.
func Warning(f interface{}, v ...interface{}) {
	beeLogger.Warn(formatLog(f, v...))
}

// Warn compatibility alias for Warning()
func Warn(f interface{}, v ...interface{}) {
	beeLogger.Warn(formatLog(f, v...))
}

// Notice logs a message at notice level.
func Notice(f interface{}, v ...interface{}) {
	beeLogger.Notice(formatLog(f, v...))
}

// Informational logs a message at info level.
func Informational(f interface{}, v ...interface{}) {
	beeLogger.Info(formatLog(f, v...))
}

// Info compatibility alias for Warning()
func Info(f interface{}, v ...interface{}) {
	beeLogger.Info(formatLog(f, v...))
}

// Debug logs a message at debug level.
func Debug(f interface{}, v ...interface{}) {
	beeLogger.Debug(formatLog(f, v...))
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Trace(f interface{}, v ...interface{}) {
	beeLogger.Trace(formatLog(f, v...))
}

//消息格式化
func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
