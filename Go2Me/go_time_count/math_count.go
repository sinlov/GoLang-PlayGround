package go_time_count

import (
	"fmt"
	"time"
)

func add(x, y int) {
	z := x + y
	fmt.Printf("add number %v \n", z)
}

func MultiAdd(start, count int) {
	fmt.Printf("start at number %v, end count %v\n",start, count)
	start_time := time.Now()
	ch := make(chan int, count)
	for i := start; i < start + count; i++ {
		go add(i, i)
		select {
		case ch <- i:
		default:
			fmt.Printf("channl full at numer %v \n",i)
		}
	}
	<-ch
	end_time := time.Since(start_time).Nanoseconds()
	fmt.Printf("MultiAdd time: %v \n", end_time)
}
