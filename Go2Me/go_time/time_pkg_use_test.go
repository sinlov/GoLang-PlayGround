package go_time

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNowUnix(t *testing.T) {
	unix := time.Now().Unix()
	t.Logf("unix %v", unix)
	micro := time.Now().UnixMicro()
	t.Logf("micro %v", micro)
	milli := time.Now().UnixMilli()
	t.Logf("milli %v", milli)
}

func TestNowTimeStr(t *testing.T) {
	convey.Convey("TestNowTimeStr", t, func() {
		// mock
		nowTimeStr := NowTimeStr()
		// do
		t.Logf("nowTimeStr: %v", nowTimeStr)
		// verify
		convey.So(nowTimeStr, convey.ShouldNotBeNil)
	})
}

func TestNowTimeFormat(t *testing.T) {
	convey.Convey("mock TestNowTimeFormat", t, func() {
		// mock
		type args struct {
			format string
		}
		type wants struct {
			wantFormat string
		}
		now := time.Now()
		tests := []struct {
			name  string
			args  args
			wants wants
		}{
			{
				name: "Unix base",
				args: args{
					format: time.UnixDate,
				},
				wants: wants{
					wantFormat: now.String(),
				},
			},
			{
				name: "base format",
				args: args{
					format: "Mon Jan 02 15:04",
				},
				wants: wants{
					wantFormat: fmt.Sprintf("%v %v %v %v:%v", now.Weekday().String(), now.Month().String(), now.Day(), now.Hour(), now.Minute()),
				},
			},

			{
				name: "custom format",
				args: args{
					format: "Monday-January-02 15:04",
				},
				wants: wants{
					wantFormat: fmt.Sprintf("%v-%v-%v %v:%v", now.Weekday().String(), now.Month().String(), now.Day(), now.Hour(), now.Minute()),
				},
			},
		}
		convey.Convey("do TestNowTimeFormat", func() {
			for _, test := range tests {
				// do
				nowTimeFormat := NowTimeFormat(test.args.format)
				convey.Convey("verify TestNowTimeFormat"+test.name, func() {
					// verify
					t.Logf("format: %v | retrun [ %v ] want [ %v ]", test.args.format, nowTimeFormat, test.wants.wantFormat)
					//convey.So("", convey.ShouldEqual, "")
					convey.So(nowTimeFormat, convey.ShouldNotBeNil)
				})
			}
		})
	})
}
