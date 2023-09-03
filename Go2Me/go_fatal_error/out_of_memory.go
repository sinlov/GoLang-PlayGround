//go:build !test

package go_fatal_error

import (
	"container/list"
	"fmt"
	"time"
)

//nolint:golint,unused
func listOutOfMem() {
	fmt.Printf("start list OOM %s\n", time.Now().String())
	forOOM := list.New()
	for {
		oneVoid := []struct {
			name string
		}{
			{
				name: "OOM data",
			},
		}
		forOOM.PushBack(oneVoid)

	}
	//fmt.Printf("end OOM %s\n", time.Now().String())
}
