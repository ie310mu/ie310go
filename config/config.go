package config

import (
	"io/ioutil"

	"github.com/ie310mu/ie310go/common/json"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/throw"
)

//ScanFromFile ..
func ScanFromFile(fileName string, dest interface{}) {
	b, err := ioutil.ReadFile(fileName)
	throw.CheckErr(err)

	jsonStr := string(b)
	logsagent.Info("the value in config file is:")
	logsagent.Info(jsonStr)

	json.FromJSON(jsonStr, dest)
	throw.CheckErr(err)

	data := json.ToJSON(dest)
	newStr := string(data[:])
	logsagent.Info("the value after unmarshal is:")
	logsagent.Info(newStr)
}
