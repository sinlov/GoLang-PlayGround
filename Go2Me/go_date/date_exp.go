package go_date

import (
	"strconv"
	"time"
	"fmt"
)

func unix() int {
	return int(time.Now().UTC().Unix())
}

func unixStr() string {
	return strconv.FormatInt(time.Now().UTC().Unix(), 10)
}

func year() int {
	return time.Now().Year()
}

func month() int {
	monthStr := fmt.Sprintf("%d", time.Now().UTC().Month())
	month, _ := strconv.Atoi(monthStr)
	return month
}

func hour() int {
	hourStr := fmt.Sprintf("%d", time.Now().UTC().Hour())
	hour, _ := strconv.Atoi(hourStr)
	return hour
}

func daysOfMonth(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}
