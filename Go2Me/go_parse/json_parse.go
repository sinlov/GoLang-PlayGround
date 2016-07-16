package go_parse

import _ "encoding/json"
import (
	_ "github.com/bitly/go-simplejson"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type MyData struct {
	Name string `json:"name"`
	Other float32 `json:"number"`
}

var detail MyData

func Test_json_parse()  {
	detail.Name = "Our"
	detail.Other = 2.3
	body, err :=json.Marshal(detail)
	jsStr, err := simplejson.NewJson(body)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(jsStr)
	fmt.Println(jsStr.Get("name"))
}