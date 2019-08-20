package go_routine

import "testing"

func TestSolve_deadlock(t *testing.T) {
	ch := make(chan int)
	quit := make(chan bool)
	SolveDeadlockBySelect(ch, quit)
}

func TestSolveDeadlockByClose(t *testing.T) {
	ch := make(chan int)
	SolveDeadlockByClose(ch, 10)
}
