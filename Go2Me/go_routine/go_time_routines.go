package go_routine

import (
	"fmt"
	"time"
)

func sumStartEnd(start, end int, ch chan string) {
	ch <- fmt.Sprintf("Sum from %d to %d is starting at %s\n", start, end, time.Now().String())
	var sum int = 0
	for i := start; i < end; i++ {
		sum += i
	}
	time.Sleep(time.Second * 10)
	ch <- fmt.Sprintf("Sum from %d to %d is %d at %s\n", start, end, sum, time.Now().String())
	//ch <- fmt.Sprintf("Time %v Sum from %d to %d is %d\n", time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST"), start, end, sum)
}

func routinesSumStartEnd(isBuff bool) {
	var ch chan string
	if isBuff {
		ch = make(chan string, 10)
	} else {
		ch = make(chan string)

	}

	for i := 0; i < 10; i++ {
		go sumStartEnd(i, i+10, ch)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(<-ch)
	}
}
