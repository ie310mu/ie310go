package route

import (
	"strconv"

	"time"

	"github.com/ie310mu/ie310go/common/json"

	//"code.google.com/p/mahonia"
	"github.com/ie310mu/ie310go/common/convert"
)

//ServiceArgs ..
type ServiceArgs struct {
	rs RequestResponse
}

//GetRs ..
func (args *ServiceArgs) GetRs() RequestResponse {
	return args.rs
}

//GetStringParamWithCheck 获取字符串参数，canBeNullOrEmpty指定是否可以为空
func (args *ServiceArgs) GetStringParamWithCheck(paramName string, canBeNullOrEmpty bool) string {
	v := args.rs.GetStringFromForm(paramName)
	if canBeNullOrEmpty == false && v == "" {
		panic("参数" + paramName + "的值不能为空")
	}

	//todo..........................url解码
	//2018.07.27暂时没有找到解决办法，通过在js强行关闭secondEncoding来处理
	//目前测试% +-!&等各种符号、中文都可以获取，但不确定是否所有中文字、所以浏览器环境都没有问题

	// if args.secondEncoding != "" && v != "" {
	// 	//这句可以解码，但中文内容好像不对，而且也不能指定编码  (unescape私有，encoding是整数且没有 utf-8)
	// 	//v, _ = url.PathUnescape(v)

	// 	//"code.google.com/p/mahonia"
	// 	// var dec mahonia.Decoderdec = mahonia.NewDecoder(args.secondEncoding)
	// 	// v, _ = dec.ConvertStringOK(v)
	// }

	return v
}

//GetStringParam 获取字符串参数，可以为空
func (args *ServiceArgs) GetStringParam(paramName string) string {
	return args.GetStringParamWithCheck(paramName, true)
}

//GetStringParamWithDefault 获取字符串参数，可指定默认值
func (args *ServiceArgs) GetStringParamWithDefault(paramName string, defaultValue string) string {
	result := args.GetStringParamWithCheck(paramName, true)
	if result == "" {
		result = defaultValue
	}
	return result
}

//GetBoolParam 获取bool参数，不能为空
func (args *ServiceArgs) GetBoolParam(paramName string) bool {
	str := args.GetStringParamWithCheck(paramName, false)
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true
	case "0", "f", "F", "false", "FALSE", "False":
		return false
	default:
		panic("参数" + paramName + "的值不是bool")
	}
}

//GetBoolParamWithDefault 获取bool参数，可指定默认值
func (args *ServiceArgs) GetBoolParamWithDefault(paramName string, defaultValue bool) (r bool) {
	defer func() {
		if err := recover(); err != nil {
			r = defaultValue
		}
	}()

	b := args.GetBoolParam(paramName)
	return b
}

//GetIntParam 获取int参数，不能为空
func (args *ServiceArgs) GetIntParam(paramName string) int {
	str := args.GetStringParamWithCheck(paramName, false)
	i, err := strconv.Atoi(str)
	if err != nil {
		panic("参数" + paramName + "的值不是int")
	}
	return i
}

//GetIntParamWithDefault 获取bool参数，可指定默认值
func (args *ServiceArgs) GetIntParamWithDefault(paramName string, defaultValue int) (r int) {
	defer func() {
		if err := recover(); err != nil {
			r = defaultValue
		}
	}()

	i := args.GetIntParam(paramName)
	return i
}

//GetFloatParam 获取float参数，不能为空
func (args *ServiceArgs) GetFloatParam(paramName string) float64 {
	str := args.GetStringParamWithCheck(paramName, false)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic("参数" + paramName + "的值不是float")
	}
	return f
}

//GetFloatParamWithDefault 获取float参数，可指定默认值
func (args *ServiceArgs) GetFloatParamWithDefault(paramName string, defaultValue float64) (r float64) {
	defer func() {
		if err := recover(); err != nil {
			r = defaultValue
		}
	}()

	f := args.GetFloatParam(paramName)
	return f
}

//GetDateTimeParam 获取DateTime参数，不能为空
func (args *ServiceArgs) GetDateTimeParam(paramName string) time.Time {
	str := args.GetStringParamWithCheck(paramName, false)
	return convert.StrToDateTime(str)
}

//GetDateTimeParamWithDefault 获取DateTime参数，可指定默认值
func (args *ServiceArgs) GetDateTimeParamWithDefault(paramName string, defaultValue time.Time) (r time.Time) {
	defer func() {
		if err := recover(); err != nil {
			r = defaultValue
		}
	}()

	t := args.GetDateTimeParam(paramName)
	return t
}

//GetDateParam 获取Date参数，不能为空
func (args *ServiceArgs) GetDateParam(paramName string) time.Time {
	str := args.GetStringParamWithCheck(paramName, false)
	return convert.StrToDate(str)
}

//GetDateParamWithDefault 获取Date参数，可指定默认值
func (args *ServiceArgs) GetDateParamWithDefault(paramName string, defaultValue time.Time) (r time.Time) {
	defer func() {
		if err := recover(); err != nil {
			r = defaultValue
		}
	}()

	t := args.GetDateParam(paramName)
	return t
}

var zeroTime time.Time

//GetBeginDateParam 获取开始日期（仅日期部分），获取不到d.IsZero()==true，对应参数名beginDate，格式2019-08-12
func (args *ServiceArgs) GetBeginDateParam() time.Time {
	d := args.GetDateParamWithDefault("beginDate", zeroTime)
	return d
}

//GetEndDateParam 获取结束日期（仅日期部分，且加1天），获取不到d.IsZero()==true，对应参数名endDate，格式2019-08-12
func (args *ServiceArgs) GetEndDateParam() time.Time {
	d := args.GetDateParamWithDefault("endDate", zeroTime)
	if !d.IsZero() {
		d = time.Date(d.Year(), d.Month(), d.Day()+1, 0, 0, 0, 0, time.Local)
	}
	return d
}

//GetBeginDateTimeParam 获取开始日期时间
func (args *ServiceArgs) GetBeginDateTimeParam() time.Time {
	d := args.GetDateTimeParamWithDefault("beginDate", zeroTime)
	if !d.IsZero() {
		d = time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute(), d.Second()+1, 0, time.Local)
	}
	return d
}

//GetEndDateTimeParam 获取结束日期时间
func (args *ServiceArgs) GetEndDateTimeParam() time.Time {
	d := args.GetDateTimeParamWithDefault("endDate", zeroTime)
	if !d.IsZero() {
		d = time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute(), d.Second()+1, 0, time.Local)
	}
	return d
}

//GetPageSizeParam 获取参数：每页数据条数，默认20，最大100
func (args *ServiceArgs) GetPageSizeParam() int {
	pageSize := args.GetIntParamWithDefault("pageSize", 20)
	if pageSize < 0 {
		pageSize = 0
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return pageSize
}

//GetCurrentPageParam 获取参数：当前页数，从1开始，默认为1，值小于1时重置为1
func (args *ServiceArgs) GetCurrentPageParam() int {
	currentPage := args.GetIntParamWithDefault("currentPage", 1)
	if currentPage <= 0 {
		currentPage = 1
	}
	return currentPage
}

//GetObjectParam 获取一个对象
func (args *ServiceArgs) GetObjectParam(paramName string, obj interface{}) {
	jsonStr := args.GetStringParamWithCheck(paramName, false)
	json.FromJSON(jsonStr, obj)
}

//GetIsMobile 是否从移动设备访问，对应参数isMobile=true
func (args *ServiceArgs) GetIsMobile() bool {
	isMobile := args.GetBoolParamWithDefault("isMobile", false)
	return isMobile
}

//GetIsApp 是否从app访问，对应参数isApp=true
func (args *ServiceArgs) GetIsApp() bool {
	isApp := args.GetBoolParamWithDefault("isApp", false)
	return isApp
}

//GetRemoteIP ..
func (args *ServiceArgs) GetRemoteIP() string {
	return args.rs.GetRemoteIP()
}

//GetSessionID ..
func (args *ServiceArgs) GetSessionID() string {
	return args.rs.GetSessionID()
}
