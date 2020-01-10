package go_parse

import _ "encoding/json"
import (
	_ "github.com/bitly/go-simplejson"
)

type MyData struct {
	Name  string  `json:"name"`
	Other float32 `json:"number"`
}

var detail MyData

type DefaultData struct {
	Bool   bool     `json:"bool"`
	String string   `json:"string"`
	Number int64    `json:"number"`
	Float  float64  `json:"float"`
	Array  []string `json:"array"`
	Object *MyData  `json:"object"`
}
