//go:build !test

package go_fatal_error

import (
	"fmt"
	"time"
)

//nolint:golint,unused
func recursiveMethod() {
	for {
		recursiveMethod()
	}
}

//nolint:golint,unused
func recursiveOverFlows() {
	fmt.Printf("start recursive %s\n", time.Now().String())
	recursiveMethod()
	fmt.Printf("end recursive %s\n", time.Now().String())
}
