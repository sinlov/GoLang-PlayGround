package go_parse

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/donnie4w/json4g"
)

var json_data = `{"status":200,"msg":"\u6210\u529f","data":[{"RemotePort":"","Subdomain":"sunny","Hostname":"","HttpAuth":"","Proto":{"http":"192.168.2.192:80","tcp":"192.168.2.192:80"}}]}`

type Proto struct {
	Http string `json:"http"`
	Tcp  string `json:"tcp"`
}

type Info struct {
	RemotePort string           `json:"RemotePort"`
	Subdomain  string           `json:"Subdomain"`
	Hostname   string           `json:"Hostname"`
	HttpAuth   string           `json:"HttpAuth"`
	Proto      map[string]Proto `json:"Proto"`
}

type Body struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []Info `json:"data"`
}

func Test_mutli_json_double() {
	js, err := simplejson.NewJson([]byte(json_data))
	first, err := json.Marshal(js)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println("json.Marshal", string(first))
	js.Set("status", nil)
	fmt.Println(js.Get("status"))
	fmt.Println(js)
	sec, err := json.Marshal(js)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println("Set and Get", string(sec))
	js.Del("status")
	thr, err := json.Marshal(js)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println("Del", string(thr))
	res, err := js.MarshalJSON()
	fmt.Println("MarshalJSON", string(res))

	js4g, err := json4g.LoadByString(json_data)
	fmt.Println("LoadByString", js4g.ToString())
	err = js4g.GetNodeByPath("status").SetValue(403)
	if err != nil {
		return
	}
	fmt.Println(js4g.ToString())
	err = js4g.DelNode("status")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("DelNode", js4g.ToString())

}
