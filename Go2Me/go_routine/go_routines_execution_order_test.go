package go_routine

import "testing"

func TestRoutinesOrderWithChan(t *testing.T) {
	ch := make(chan int)
	go func() {
		t.Log("start goroutine")
		ch <- 0
		t.Log("done goroutine")
	}()
	t.Log("before goroutine")
	<-ch
	//close(ch)
	t.Log("function done")
}
