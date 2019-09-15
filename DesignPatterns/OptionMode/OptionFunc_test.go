package OptionMode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewOption(t *testing.T) {
	convey.Convey("mock TestNewOption", t, func() {
		// mock
		absOptions := newOptionABS("A", "B", 100)
		optionFast := NewOption()
		changeOption := NewOption(
			WithA("B"),
			WithB("A"),
		)
		convey.Convey("do TestNewOption", func() {
			// do
			t.Logf("absOptions %v", absOptions)
			t.Logf("optionFast %v", optionFast)
			t.Logf("changeOption %v", changeOption)
			convey.Convey("verify TestNewOption", func() {
				// verify
				convey.So(optionFast.A, convey.ShouldEqual, absOptions.A)
				convey.So(optionFast.B, convey.ShouldEqual, absOptions.B)
				convey.So(optionFast.C, convey.ShouldEqual, absOptions.C)
				convey.So(changeOption.A, convey.ShouldEqual, "B")
				convey.So(changeOption.B, convey.ShouldEqual, "A")
				convey.So(changeOption.C, convey.ShouldEqual, 100)
			})
		})
	})
}
