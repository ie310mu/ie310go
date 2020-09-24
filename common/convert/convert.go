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

func GetToday() time.Time {
	//下面3种方法都可以获取今天凌晨
	// timeStr := time.Now().Format("2006-01-02") //取今天的日期部分字符串
	// t := StrToDateTime(timeStr)
	//fmt.Println(t)

	//t, _ = time.ParseInLocation("2006-01-02", timeStr, time.Local)
	//fmt.Println(t)

	t := time.Now()
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	//fmt.Println(t)
	return t

	//验证代码：
	// lastTime := convert.StrToDateTime("2019-12-26 01:00")
	// if lastTime.Unix() > convert.GetToday().Unix() {
	// 	logsagent.Info("11111111111111111111")
	// }
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
		//t, err := time.Parse(format, str)                       //Parse默认是UTC时区，用t.Local()转换成Local后会多8个小时（北京时区）
		t, err := time.ParseInLocation(format, str, time.Local)
		if err == nil {
			return t
		}
	}

	panic("参数的值不是Time:" + str)

	/*
		https://studygolang.com/articles/14933?fr=sidebar

		t := time.Now()
		fmt.Println("now : ", t)
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local)
		fmt.Println("convert to local : ", t)
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
		fmt.Println("convert to UTC : ", t)

		timeStr := "2019.12.26 09:08:42"
		t = convert.StrToDateTime(timeStr)
		fmt.Println("local convert : ", t)
		t = convert.StrToDateTimeNotLocal(timeStr)
		fmt.Println("utc : ", t)
		fmt.Println("utc to local : ", t.Local())
		fmt.Println("utc to local : ", t.Local().Local())

		now :  2019-12-26 09:22:14.7338648 +0800 CST m=+0.008020701
		convert to local :  2019-12-26 09:22:14 +0800 CST
		convert to UTC :  2019-12-26 09:22:14 +0000 UTC
		local convert :  2019-12-26 09:08:42 +0800 CST
		utc :  2019-12-26 09:08:42 +0000 UTC
		utc to local :  2019-12-26 17:08:42 +0800 CST
		utc to local :  2019-12-26 17:08:42 +0800 CST

		1.DateTimeToStr函数是抛弃了后面的时区信息的，fmt.Println不会
		2.time.Parse将字符串转换为时间后，会将时区设置为UTC，用t.Local()得到本地时区信息后，时间会加8个小时，后面时区显示为+0800 CST，代表CST比UTC快8小时
		3.所以前面用time.ParseInLocation(format, str, time.Local)，确保时间信息是对的，时区也是对的，因为我们肯定是用Local时区
		4.其实，上面用time.Parse时，和前端、其他系统传输时，都会使用DateTimeToStr抛弃时区部分，问题不大；但如果要存入数据库，就不对了
		总结：将字符串转换为日期时间时，要用time.ParseInLocation明确指定时区，不要用time.Parse，time.Parse使用的时区是UTC和我们的Local北京时间是不一样的
				  time.Now()获取的是本地时区
				  所有时间，统一用Lcoal时区不要用UTC，这样就不会出现问题 ************************************


		timeStr := time.Now().Format("2006-01-02") //取今天的日期部分字符串
		t, _ := time.Parse("2006-01-02", timeStr)  //转换成日期时间
		fmt.Println(": ", t)
		t = time.Unix(t.Unix(), 0)   //这个返回本地时区
		fmt.Println(": ", t)

		:  2019-12-26 00:00:00 +0000 UTC
		:  2019-12-26 08:00:00 +0800 CST  //UTC转为CST就是要多8个小时

		https://blog.csdn.net/wangluotong00/article/details/52074226   验证如下：
		//timeStr := time.Now().Format("2006-01-02")
		timeStr := "2016-07-30"
		fmt.Println("timeStr:", timeStr)
		t, _ := time.Parse("2006-01-02", timeStr)
		timeNumber := t.Unix()
		fmt.Println("timeNumber:", timeNumber)
		t, _ = time.ParseInLocation("2006-01-02", timeStr, time.Local)
		timeNumber = t.Unix()
		fmt.Println("timeNumber:", timeNumber)

		timeStr: 2016-07-30
		timeNumber: 1469836800
		timeNumber: 1469808000   //都是同样的字面时间，但Local时区早8个小时，t.Unix()是以UTC为起点，所以这个时间戳比前面的要小8个小时

	*/

}

func StrToDateTimeNotLocal(str string) time.Time {
	for _, format := range timeFormats {
		t, err := time.Parse(format, str) //Parse默认是UTC时区，用t.Local()转换成Local后会多8个小时（北京时区）
		if err == nil {
			return t
		}
	}

	panic("参数的值不是Time:" + str)
}

//StrToDateTime2 字符串转换为日期（不去除时间部分，使用Local）  转换错误返回默认值
func StrToDateTime2(str string, defaultValue time.Time) time.Time {
	for _, format := range timeFormats {
		//t, err := time.Parse(format, str)                       //Parse默认是UTC时区，用t.Local()转换成Local后会多8个小时（北京时区）
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
