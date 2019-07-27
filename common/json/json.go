package json

import (
	"encoding/json"

	"github.com/ie310mu/ie310go/common/throw"
)

//ToJSON ..
func ToJSON(obj interface{}) string {
	data, err := json.Marshal(obj)
	throw.CheckErr(err)
	json := string(data[:])

	return json
}

//FromJSON ..
/*

对象
jsonStr := "{\"id\":\"111\",\"name\":\"222\"}"
user := model.User{}
json.FromJSON(jsonStr, &user)  //注意传指针

数组：
var items []*model.GwYqsb
json.FromJSON(jsonStr, &items)

yqsbIdsJSON := args.GetStringParamWithCheck("yqsbIdsJson", false)
var yqsbIds []string
json.FromJSON(yqsbIdsJSON, &yqsbIds)

*/
func FromJSON(jsonStr string, obj interface{}) {
	err := json.Unmarshal([]byte(jsonStr), obj)
	throw.CheckErr(err)
}
