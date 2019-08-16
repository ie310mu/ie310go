package dir

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ie310mu/ie310go/common/throw"
)

//sep := string(os.PathSeparator)

//GetCurrentPath ..
func GetCurrentPath() string {
	// file, err := exec.LookPath(os.Args[0])  //exec.LookPath:解析环境变量？
	// throw.CheckErr(err)
	file := os.Args[0]
	path, err := filepath.Abs(file)
	throw.CheckErr(err)
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		panic(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1])
}

//PathExists ..
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//CreatePath ..
func CreatePath(path string) {
	err := os.Mkdir(path, os.ModePerm)
	throw.CheckErr(err)
}

//DeletePath ..
func DeletePath(path string) {
	err := os.Remove(path)
	throw.CheckErr(err)
}
