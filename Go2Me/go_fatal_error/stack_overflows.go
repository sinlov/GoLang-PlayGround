package go_fatal_error

import (
	"fmt"
	"time"
)

func recursiveMethod() {
	for ; ; {
		recursiveMethod()
	}
}

func recursiveOverFlows() {
	fmt.Printf("start recursive %s\n", time.Now().String())
	recursiveMethod()
	fmt.Printf("end recursive %s\n", time.Now().String())
}
