package go_runtime

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loopIO(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("%d\n", i)
	}
	quit <- 0
}

func runDiffMaxProcessWithIO(n, count int) {
	if n != 0 {
		// default 0 not set
		runtime.GOMAXPROCS(n)
	}

	for i := 1; i <= count; i++ {
		go loopIO(count)
	}

	for i := 0; i < count; i++ {
		<-quit
	}

}
