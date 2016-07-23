package go_base

import (
	"fmt"
	"testing"
)

func Test_if(t *testing.T) {
	x := 10
	if x > 10 {
		fmt.Println("number is max 10")
	} else {
		fmt.Println("number is min 10")
	}
}
