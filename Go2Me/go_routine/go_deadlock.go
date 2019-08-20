package go_routine

import "fmt"

const MaxBizCount int = 5

// for SolveDeadlockBySelect
func doBizCount(ch chan int, quit chan bool) {
	for i := 0; i < MaxBizCount; i++ {
		ch <- i
	}
	quit <- true
}

// most of deadlock use this way to solve
//	first: use go doBizCount(biz, quit)
//	second: use select
//	third: use goto to let main goroutines free
func SolveDeadlockBySelect(biz chan int, quit chan bool) {
	go doBizCount(biz, quit)

	// select chanel
	for {
		select {
		case i := <-biz:
			fmt.Printf("do biz at: %d\n", i)
		case <-quit:
			// can not use break, to solve deadlock
			goto end
		}
	}
end:
	fmt.Println("deadlock has solve")
}

// for SolveDeadlockByClose function
func doBizCountByClose(biz chan int) {
	for i := 0; i < MaxBizCount; i++ {
		biz <- i
	}
	// must use this for panic: all goroutines are asleep - deadlock!
	close(biz)
}

// Not recommended
//	will use more goroutine
func SolveDeadlockByClose(biz chan int, maxCheck int) {
	var count int
	go doBizCountByClose(biz)
	for {
		if count > maxCheck {
			fmt.Printf("max count biz %d, exit function\n", count)
			return
		}
		if i, ok := <-biz; ok {
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		} else {
			// print at this will show why this method will use more goroutine
			fmt.Printf("not ok at this biz, counter is at %d\n", i)
		}
		count++
	}
}
