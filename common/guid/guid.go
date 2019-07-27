package guid

import (
	"github.com/ie310mu/ie310go/common/throw"
	uuid "github.com/ie310mu/ie310go/forks/github.com/satori/go.uuid"
)

//Get 获取失败会抛出异常
func Get() string {
	uid, err := uuid.NewV4()
	throw.CheckErr(err)
	return uid.String()
}
