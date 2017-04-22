package go_routine

import (
	"testing"
	"time"
)

func BenchmarkReadToken(b *testing.B) {
	NewTokenInfo()
	for i := 0; i < 20; i++ {
		ReadToken("12321", 10)
		time.Sleep(time.Second * 2)
	}
}

func BenchmarkReadTokenByGoto(b *testing.B) {
	NewTokenInfo()
	for i := 0; i < 20; i++ {
		ReadTokenByGoto("12321", 10)
		time.Sleep(time.Second * 2)
	}
}
