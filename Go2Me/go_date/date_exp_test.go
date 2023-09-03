package go_date

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
)

func TestUnix(t *testing.T) {
	convey.Convey("mock TestUnix", t, func() {
		// mock
		convey.Convey("do TestUnix", func() {
			// do
			unix := unix()
			convey.Convey("verify TestUnix", func() {
				// verify
				fmt.Printf("unix %v\n", unix)
				convey.So(unix, convey.ShouldNotEqual, 0)
			})
		})
	})
}

func TestUnixStr(t *testing.T) {
	convey.Convey("mock TestUnixStr", t, func() {
		// mock

		convey.Convey("do TestUnixStr", func() {
			// do
			unixStr := unixStr()
			convey.Convey("verify TestUnixStr", func() {
				// verify
				fmt.Printf("unixStr %v\n", unixStr)
				convey.So(unixStr, convey.ShouldNotEqual, 0)
			})
		})
	})
}

func TestYear(t *testing.T) {
	convey.Convey("mock TestYear", t, func() {
		// mock
		convey.Convey("do TestYear", func() {
			// do
			year := year()
			convey.Convey("verify TestYear", func() {
				// verify
				fmt.Printf("year %v\n", year)
				convey.So(year, convey.ShouldNotEqual, 0)
			})
		})
	})
}

func TestMonth(t *testing.T) {
	convey.Convey("mock TestMonth", t, func() {
		// mock
		convey.Convey("do TestMonth", func() {
			// do
			month := month()
			convey.Convey("verify TestMonth", func() {
				// verify
				fmt.Printf("month %v\n", month)
				convey.So(month, convey.ShouldNotEqual, 0)
			})
		})
	})
}

func TestHour(t *testing.T) {
	convey.Convey("mock TestHour", t, func() {
		// mock
		convey.Convey("do TestHour", func() {
			// do
			hour := hour()
			convey.Convey("verify TestHour", func() {
				// verify
				t.Logf("hour %v", hour)
				convey.So("", convey.ShouldEqual, "")
			})
		})
	})
}

func TestDaysOfMonth(t *testing.T) {
	convey.Convey("mock TestDaysOfMonth", t, func() {
		// mock
		now := time.Now()
		year := now.Year()
		fmt.Printf("year %v\n", year)
		currentTime := strconv.FormatInt(now.UTC().Unix(), 10)
		fmt.Printf("currentTime %v\n", currentTime)
		unix := time.Unix(now.UTC().Unix(), 0)
		//if err != nil {
		//	t.Fatalf("mock currentTime number error %s", err)
		//}
		monthStr := fmt.Sprintf("%d", unix.Month())
		fmt.Printf("monthStr %s\n", monthStr)
		month, err := strconv.Atoi(monthStr)
		if err != nil {
			t.Fatalf("mock moth number error %s", err)
		}
		fmt.Printf("month %v\n", month)
		convey.Convey("do TestDaysOfMonth", func() {
			// do
			days := daysOfMonth(year, month)
			convey.Convey("verify TestDaysOfMonth", func() {
				// verify
				fmt.Printf("days %v\n", days)
				convey.So(days, convey.ShouldNotEqual, 0)
			})
		})
	})
}
