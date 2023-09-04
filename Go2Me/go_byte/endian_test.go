package go_byte

import (
	"testing"
)

func TestIsBigEndian(t *testing.T) {
	bigEndian := IsBigEndian()
	t.Logf("is bigEndian: %v", bigEndian)
}

func TestIsLittleEndian(t *testing.T) {
	littleEndian := IsLittleEndian()
	t.Logf("is littleEndian: %v", littleEndian)
}

func BenchmarkIsBigEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigEndian := IsBigEndian()
		//b.Logf("is bigEndian: %v", bigEndian)
		if bigEndian {

		}
	}
}

func BenchmarkIsLittleEndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		littleEndian := IsLittleEndian()
		//b.Logf("is littleEndian: %v", littleEndian)
		if littleEndian {

		}
	}
}
