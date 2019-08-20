package go_routine

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to channel
}

func go_sum() {
	a := []int{7, 2, 8, -7, 4, 0}
	c := make(chan int)
	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Printf("value x: %v,| value y: %v\n", x, y)
}