package go_reflect

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAnonymousTypeCheck(t *testing.T) {
	convey.Convey("TestAnonymousTypeCheck", t, func() {
		// mock

		// do
		anonymousTypeCheck()
		// verify
		convey.So("", convey.ShouldEqual, "")
	})
}
