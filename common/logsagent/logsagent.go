/*
	此工具用于在ie310go的模块中注入真实的日志输出，未注册时，默认使用console输出
*/

package logsagent

import (
	"fmt"
	"strings"
)

func fmtPrintf(f interface{}, v ...interface{}) bool {
	if logger == nil {
		msg := formatLog(f, v)
		fmt.Print(msg)
		return true
	}

	return false
}

//Emergency ..
func Emergency(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Emergency(fmt.Sprintf("%v", f), v...)
	}
}

//Alert ..
func Alert(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Alert(fmt.Sprintf("%v", f), v...)
	}
}

//Critical ..
func Critical(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Critical(fmt.Sprintf("%v", f), v...)
	}
}

//Error ..
func Error(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Error(fmt.Sprintf("%v", f), v...)
	}
}

//Warn ..
func Warn(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Warn(fmt.Sprintf("%v", f), v...)
	}
}

//Notice ..
func Notice(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Notice(fmt.Sprintf("%v", f), v...)
	}
}

//Info ..
func Info(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Info(fmt.Sprintf("%v", f), v...)
	}
}

//Debug ..
func Debug(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Debug(fmt.Sprintf("%v", f), v...)
	}
}

//Trace ..
func Trace(f interface{}, v ...interface{}) {
	if !fmtPrintf(f, v) {
		logger.Trace(fmt.Sprintf("%v", f), v...)
	}
}

var logger Logger

//SetLogger ..
func SetLogger(l Logger) {
	logger = l
}

//Logger ..
type Logger interface {
	Emergency(f string, v ...interface{})
	Alert(f string, v ...interface{})
	Critical(f string, v ...interface{})
	Error(f string, v ...interface{})
	Warn(f string, v ...interface{})
	Notice(f string, v ...interface{})
	Info(f string, v ...interface{})
	Debug(f string, v ...interface{})
	Trace(f string, v ...interface{})
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
