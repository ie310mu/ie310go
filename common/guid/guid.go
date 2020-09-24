package guid

import (
	uuid "github.com/satori/go.uuid"
)

//Get 获取失败会抛出异常
func Get() string {
	uid := uuid.NewV4()
	return uid.String()
}
