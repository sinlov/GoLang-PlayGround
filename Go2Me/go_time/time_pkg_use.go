package go_time

import "time"

func NowTimeStr() string {
	t := time.Now()
	return t.String()
}

// use string
//
//	Mon Jan _2 15:04:05 MST 2006 form time.UnixDate
//
// or
//
//	2006-01-02T15:04:05.999999999Z07:00 MST time.RFC3339Nano
//	to format time
func NowTimeFormat(timeFormat string) string {
	return time.Now().Format(timeFormat)
}
