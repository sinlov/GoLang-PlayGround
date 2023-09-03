package go_byte

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsBigEndian(t *testing.T) {
	convey.Convey("TestIsBigEndian", t, func() {
		// mock
		bigEndian := IsBigEndian()
		t.Logf("is bigEndian: %v", bigEndian)
		// do
		// verify
		convey.So("", convey.ShouldEqual, "")
	})
}

func TestIsLittleEndian(t *testing.T) {
	convey.Convey("TestIsLittleEndian", t, func() {
		// mock
		littleEndian := IsLittleEndian()
		t.Logf("is littleEndian: %v", littleEndian)
		// do
		// verify
		convey.So("", convey.ShouldEqual, "")
	})
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
