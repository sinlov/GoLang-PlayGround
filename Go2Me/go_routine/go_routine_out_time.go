package go_routine

import (
	"fmt"
	"time"
)

// 默认 5 秒超时
const DEFAULT_TIME_OUT time.Duration = time.Second * 5

func count_sum_out(start, end int, ch chan string, is_time_out bool) int {
	var sum int = 0
	for i := start; i < end; i++ {
		sum += i
	}
	if is_time_out {
		time.Sleep(time.Second * 8)
	} else {
		time.Sleep(time.Second * 3)
	}
	ch <- fmt.Sprintf("Sum from %d to %d is %d\n", start, end, sum)
	return sum
}

func try_count_sum_out(is_time_out bool) {
	var ch = make(chan string, 1)
	var timeout = make(chan bool, 1)

	go count_sum_out(1, 10, ch, is_time_out)
	go func() {
		time.Sleep(DEFAULT_TIME_OUT)
		timeout <- true
	}()

	select {
	case sum := <-ch:
		fmt.Print(sum)
	case <-timeout:
		fmt.Printf("Sorry, TIMEOUT! at %v\n", DEFAULT_TIME_OUT)
	}
}
