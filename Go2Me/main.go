package main

import (
	"GoLang-PlayGround/Go2Me/go_parse"
	"fmt"
)

func main() {
	//go_parse.Test_json_parse()
	go_parse.Test_mutli_json_double()
	//res,err := go_parse.AES_Http_Get("http://10.8.230.246:8000/apis/", "c6*#e2&(g*UjX!h*")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//go_parse.Unicode2utf8()
	//go_parse.Unicode2utf8()
	str4En := "123456qwert"
	enRes := go_parse.MByteEncode([]byte(str4En))
	fmt.Println(enRes)
}