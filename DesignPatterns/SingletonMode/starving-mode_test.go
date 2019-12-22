package SingletonMode

import "testing"

func TestGetHungry(t *testing.T) {
	h1 := GetStarving()
	t.Logf("GetStarving first name: %v\n", h1.name)
	h2 := GetStarving()
	t.Logf("GetStarving second name: %v\n", h2.name)
}

func TestGetHungry_Empty(t *testing.T) {
	a := []int{7, 2, 8, -7, 4, 0}
	c := make(chan int)
	go moreHungry(a, c)
	select {
	default:
		t.Logf("starving is empty\n")
	case <-c:
		x := <-c
		if x < 0 {
			t.Logf("lazy is empty, value x: %v\n", x)
		} else {
			t.Logf("value x: %v\n", x)
		}
	}
}

func moreHungry(a []int, c chan int) *starving {
	hungry := GetStarving()
	if hungry == nil {
		c <- -1
	}
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
	return hungry
}
