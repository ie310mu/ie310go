package md5

import (
	"crypto/md5"
	"fmt"
)

//GetStringMD5 获取失败会抛出异常
func GetStringMD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
