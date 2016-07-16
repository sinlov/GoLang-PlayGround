package go_parse

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"fmt"
)

var json_data = `{"status":200,"msg":"\u6210\u529f","data":[{"RemotePort":"","Subdomain":"sunny","Hostname":"","HttpAuth":"","Proto":{"http":"192.168.2.192:80","tcp":"192.168.2.192:80"}}]}`

type Proto struct {
	Http string `json:"http"`
	Tcp string `json:"tcp"`
}

type Info struct {
	RemotePort string `json:"RemotePort"`
	Subdomain string `json:"Subdomain"`
	Hostname string `json:"Hostname"`
	HttpAuth string `json:"HttpAuth"`
	Proto map[string]Proto `json:"Proto"`
}

type Body struct{
	Status int `json:"status"`
	Msg string `json:"msg"`
	Data []Info `json:"data"`
}

func Test_mutli_json_double()  {
	body, err := json.Marshal(json_data)
	if err != nil {
		panic(err.Error())
	}
	jsObj, err := simplejson.NewJson(body)
	fmt.Println(jsObj)
}