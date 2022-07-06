package sliceutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasUint64(t *testing.T) {
	// mock HasUint64
	v := []uint64{12, 13, 14, 17, 20, 31}
	t.Logf("~> mock HasUint64")
	// do HasUint64
	t.Logf("~> do HasUint64")
	// verify HasUint64
	assert.True(t, HasUint64(v, 13))
	assert.False(t, HasUint64(v, 15))
}

func BenchmarkHasUint64Ordinary(b *testing.B) {
	v := []uint64{51, 52, 33, 14, 19, 16, 12, 13, 14, 17, 20, 31, 351, 352, 333, 314, 319, 316, 312, 313, 314, 317, 320, 312}
	for i := 0; i < b.N; i++ {
		assert.True(b, HasUint64(v, 13))
		assert.False(b, HasUint64(v, 15))
	}
}

func BenchmarkHasUint64BS(b *testing.B) {
	v := []uint64{51, 52, 33, 14, 19, 16, 12, 13, 14, 17, 20, 31, 351, 352, 333, 314, 319, 316, 312, 313, 314, 317, 320, 312}
	for i := 0; i < b.N; i++ {
		assert.True(b, HasUint64BS(v, 13))
		assert.False(b, HasUint64BS(v, 15))
	}
}
