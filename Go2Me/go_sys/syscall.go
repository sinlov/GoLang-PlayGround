package go_sys

import (
	"fmt"
	"syscall"
)

func rlimit() {
	var rLimit syscall.Rlimit
	fmt.Printf("rLimit.Max: %v\n", rLimit.Max)
}
