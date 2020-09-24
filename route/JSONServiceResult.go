package route

import "github.com/ie310mu/ie310go/common/json"

//JSONServiceResult ..
type JSONServiceResult struct {
	State      int         `json:"state"` //不能加omitempty，否则为0时不输出，js中解析为undefined
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	DataJSON   string      `json:"dataJSON,omitempty" `
	ExtandData interface{} `json:"extandData,omitempty"`
}

func (jsr *JSONServiceResult) updateDataJSON() {
	json := json.ToJSON(jsr.Data)
	jsr.DataJSON = json
}
