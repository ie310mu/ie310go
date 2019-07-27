package sql

import (
	"strings"
	"time"

	"github.com/ie310mu/ie310go/common/convert"
)

var opers = [...]string{">=", ">", "<=", "<", "=", "!="}

//ExpressionToSQLStr 将>10、<=35、>2018-05-23、<=2019-1-1等类似的语句，转换成sql语句，返回sql语句和sql参数
//valueType  1数值  2日期  其他类型暂未实现
func ExpressionToSQLStr(expression string, field string, valueType int) (sqlStr string, sqlArg interface{}) {
	defer func() {
		if err := recover(); err != nil {
			sqlStr = ""
			sqlArg = nil
		}
	}()

	//检查表达式
	expression = strings.Trim(expression, "")
	if expression == "" {
		return "", nil
	}

	//读取操作符
	oper := ""
	for _, o := range opers {
		index := strings.Index(expression, o)
		if index == 0 {
			oper = o
			expression = expression[len(oper):] //截取操作符
			break
		}
	}
	if oper == "" {
		return "", nil
	}

	//转换值
	switch valueType {
	case 1: //数值
		v := convert.StrToFloat(expression)
		return (field + oper + "?"), v
	case 2: //日期
		v := convert.StrToDate(expression)
		//logs.Info(v.Local().Format("2006-01-02 15:04:05.000"))
		if oper == "<=" { //加上1天，再减去1秒  (目的是将实际值中的时间部分屏蔽)
			v = v.AddDate(0, 0, 1).Add(-1 * time.Second)
		}
		//logs.Info(v.Local().Format("2006-01-02 15:04:05.000"))
		return (field + oper + "?"), v
	default:
		return "", nil
	}
}
