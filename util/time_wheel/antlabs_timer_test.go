package time_wheel

import (
	aTimer "github.com/antlabs/timer"
	"testing"
	"time"
)

func Test_antlabs_timer_once(t *testing.T) {
	tm := aTimer.NewTimer()

	tm.AfterFunc(1*time.Second, func() {
		t.Logf("AfterFunc 1 [t:%v]\n", time.Now())
	})

	tm.AfterFunc(5*time.Second, func() {
		t.Logf("AfterFunc 5 [t:%v]\n", time.Now())
		tm.Stop()
	})

	tm.Run()
}

func Test_antlabs_timer_Schedule(t *testing.T) {
	tm := aTimer.NewTimer()

	tm.ScheduleFunc(1*time.Second, func() {
		t.Logf("ScheduleFunc 1 [t:%v]\n", time.Now())
	})

	tm.ScheduleFunc(3*time.Second, func() {
		t.Logf("ScheduleFunc 3 [t:%v]\n", time.Now())
	})

	tm.AfterFunc(15*time.Second, func() {
		t.Logf("Stop all Schedule at 15 [t:%v]\n", time.Now())
		tm.Stop()
	})

	tm.Run()
}
