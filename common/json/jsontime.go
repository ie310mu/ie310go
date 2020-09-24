package json

import (
	"fmt"
	"strings"
	"time"

	"github.com/ie310mu/ie310go/common/convert"
)

/*
jst := json.JsonTime(tm)
tm := time.Time(jst)

json.JsonTime(time.Now())

*/
type JsonTime time.Time

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

// 实现它的json序列化方法
func (t *JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	// now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	str := string(data)
	if str == "" || str == "null" {
		return
	}
	str = strings.Replace(str, `"`, ``, -1)
	now := convert.StrToDateTime(str)
	*t = JsonTime(now)
	return
}
