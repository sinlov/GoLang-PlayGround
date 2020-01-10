package go_parse

import (
	"encoding/json"
	"testing"
)

func Test_json_parse(t *testing.T) {
	//detail.Name = "Our"
	detail.Other = 2.3
	body, err := json.Marshal(detail)
	if err != nil {
		t.Errorf("Marshal error %v", err)
	}
	t.Logf("detail name %v", detail.Name)
	var jsStr MyData
	err = json.Unmarshal(body, &jsStr)
	if err != nil {
		t.Errorf("Unmarshal error %v", err)
	}
	t.Logf("jsStr %v", jsStr)
	t.Logf("jsStr name %v", jsStr.Name)
}

func TestMarshalDefaultValue(t *testing.T) {
	var defaultData DefaultData
	defaultData.Array = make([]string, 0)
	defaultData.Object = &detail
	marshal, err := json.Marshal(defaultData)
	if err != nil {
		t.Errorf("json.Marshal error: %v", err)
	}
	t.Logf("defaultData %v", string(marshal))
}
