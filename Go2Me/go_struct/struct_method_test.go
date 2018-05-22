package go_struct

import (
	"testing"
	"fmt"
)

type structA struct {
	Name string
}

type structB struct {
	Name string
}

func (a *structA) print() *structA {
	fmt.Printf("%s, %v\n", "method print A", a.Name)
	// 这里写不写都没有返回，且无法接受放回
	return a
}

func (b *structB) print() string {
	fmt.Printf("%s, %v\n", "print B", b.Name)
	return b.Name
}

func Test_method_return(t *testing.T) {
	a := structA{
		Name: "a name",
	}
	b := structB{
		Name: "b name",
	}
	t.Logf("a print return %s\n", a.print())
	t.Logf("b print return %s\n", b.print())
}
