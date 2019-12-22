package SingletonMode

import "testing"

func TestGetSloth(t *testing.T) {
	s := GetLazy()
	s.name = "First assignment singleton pattern"
	t.Logf("first lazy.name: %v\n", s.name)
	s2 := GetLazy()
	t.Logf("second lazy.name: %v\n", s2.name)
}

func TestGetSloth_Empty(t *testing.T) {
	a := []int{7, 2, 8, -7, 4, 0}
	c := make(chan int)
	go moreSloth(a, c)
	select {
	default:
		t.Logf("lazy is empty\n")
	case <-c:
		x := <-c
		if x < 0 {
			t.Logf("lazy is empty, value x: %v\n", x)
		} else {
			t.Logf("value x: %v\n", x)
		}
	}
}

func moreSloth(a []int, c chan int) *lazy {
	sloth := GetLazy()
	if sloth == nil {
		c <- -1
	}
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
	return sloth
}
