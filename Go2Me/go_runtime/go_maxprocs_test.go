package go_runtime

import (
	"runtime"
	"testing"
)

func BenchmarkRunDiffMaxProcessWithIODefault(b *testing.B) {
	b.Logf("now CPU Num is: %v", runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		runDiffMaxProcessWithIO(0, 1000)
	}
}

func BenchmarkRunDiffMaxProcessWithIO(b *testing.B) {
	b.Logf("now CPU Num is: %v", runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		runDiffMaxProcessWithIO(1, 1000)
	}
}
