package convert

import (
	"strconv"
	"time"
)

var timeFormats = [...]string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02",
	"2006.01.02 15:04:05",
	"2006.01.02 15:04",
	"2006.01.02",
	"2006/01/02 15:04:05",
	"2006/01/02 15:04",
	"2006/01/02",
	"2006-1-2 15:04:05",
	"2006-1-2 15:04",
	"2006-1-2",
	"2006.1.2 15:04:05",
	"2006.1.2 15:04",
	"2006.1.2",
	"2006/1/2 15:04:05",
	"2006/1/2 15:04",
	"2006/1/2",
}

//StrToDate 字符串转换为日期（去除时间部分，使用Local）  转换错误抛出异常
func StrToDate(str string) time.Time {
	t := StrToDateTime(str)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return t
}

//DateToStr 格式2006-01-02
func DateToStr(v time.Time) string {
	s := v.Format("2006-01-02")
	return s
}

//StrToDateTime 字符串转换为日期（不去除时间部分，使用Local）  转换错误抛出异常
func StrToDateTime(str string) time.Time {
	for _, format := range timeFormats {
		//t, err := time.Parse(format, str)                       //Parse默认是UTC时区，用Time.Local()转换成Local后会多8个小时（北京时区）
		t, err := time.ParseInLocation(format, str, time.Local)
		if err == nil {
			return t
		}
	}

	panic("参数的值不是Time:" + str)
}

//StrToDateTime2 字符串转换为日期（不去除时间部分，使用Local）  转换错误返回默认值
func StrToDateTime2(str string, defaultValue time.Time) time.Time {
	for _, format := range timeFormats {
		//t, err := time.Parse(format, str)                       //Parse默认是UTC时区，用Time.Local()转换成Local后会多8个小时（北京时区）
		t, err := time.ParseInLocation(format, str, time.Local)
		if err == nil {
			return t
		}
	}

	return defaultValue
}

//DateTimeToStr 格式2006.01.02 15:04
func DateTimeToStr(v time.Time) string {
	s := v.Format("2006.01.02 15:04")
	return s
}

//DateTimeWithSecToStr 格式2006.01.02 15:04:05
func DateTimeWithSecToStr(v time.Time) string {
	s := v.Format("2006.01.02 15:04:05")
	return s
}

//StrToFloat 失败抛出异常
func StrToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return f
}

//StrToFloat2 失败返回默认值
func StrToFloat2(str string, defaultValue float64) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return defaultValue
	}

	return f
}

//Float64ToStr ..
func Float64ToStr(v float64, prec int, withLastZeroFlag bool) string {
	s := strconv.FormatFloat(v, 'f', prec, 64)

	if prec > 0 && !withLastZeroFlag { //如果小数位数大于0，而且指定了不需要最后的0，则进行处理
		l := len(s)
		for {
			if l <= 1 { //长度<=1，不处理
				break
			}
			if s[l-1:l] == "0" { //最后一位是0，去掉后继续处理
				s = s[0 : l-1]
				l = len(s)
			} else if s[l-1:l] == "." { //最后一位是.去掉后停止处理
				s = s[0 : l-1]
				break
			} else {
				break
			}
		}
	}

	return s
}

//Float32ToStr ..
func Float32ToStr(v float32, prec int, withLastZeroFlag bool) string {
	vv := float64(v)
	return Float64ToStr(vv, prec, withLastZeroFlag)
}

//StrToBool 失败抛出异常
func StrToBool(str string) bool {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true
	case "0", "f", "F", "false", "FALSE", "False":
		return false
	default:
		panic("参数的值不是bool")
	}
}

//StrToBool2 失败返回默认值
func StrToBool2(str string, defaultValue bool) bool {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true
	case "0", "f", "F", "false", "FALSE", "False":
		return false
	default:
		return defaultValue
	}
}

//BoolToStr ..
func BoolToStr(v bool) string {
	if v {
		return "True"
	}
	return "False"
}

//StrToInt 失败抛出异常
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

//StrToInt2 失败返回默认值
func StrToInt2(str string, defaultValue int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}

	return i
}

//IntToStr ..
func IntToStr(v int) string {
	s := strconv.FormatInt(int64(v), 10)
	return s
}
