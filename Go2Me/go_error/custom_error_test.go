package go_error

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDiyEqual(t *testing.T) {
	convey.Convey("mock TestDiyEqual", t, func() {
		// mock
		var testNum = 10
		convey.Convey("do TestDiyEqual", func() {
			// do
			result, err := Equal(testNum)
			convey.Convey("verify TestDiyEqual", func() {
				// verify
				convey.So(result, convey.ShouldEqual, testNum)
				convey.So(err, convey.ShouldEqual, nil)
			})
		})
	})
}
