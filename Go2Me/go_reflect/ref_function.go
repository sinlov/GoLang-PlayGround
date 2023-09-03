package go_reflect

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

type st struct {
}

func (this *st) Echo() {
	fmt.Println("echo()")
}

func (this *st) Echo2() {
	fmt.Println("echo--------------------()")
}

var xmlstr string = `<root>
<func>Echo</func>
<func>Echo2</func>
</root>`

type st2 struct {
	E []string `xml:"func"`
}

func UseMethodByName() {
	s2 := st2{}
	err := xml.Unmarshal([]byte(xmlstr), &s2)
	if err != nil {
		return
	}

	s := &st{}
	v := reflect.ValueOf(s)

	v.MethodByName(s2.E[1]).Call(nil)
}
