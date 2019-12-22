package go_struct

import (
	"testing"
)

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
