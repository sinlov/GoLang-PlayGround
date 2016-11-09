package go_routine

import (
	"testing"
	"time"
)

func TestReadToken(t *testing.T) {
	NewTokenInfo()
	for i := 0; i < 20; i++ {
		ReadToken("12332")
		time.Sleep(time.Second * 2)
	}
}
