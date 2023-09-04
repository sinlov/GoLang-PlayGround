package go_date

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnix(t *testing.T) {
	t.Logf("~> mock Unix")
	// mock Unix

	t.Logf("~> do Unix")
	// do Unix
	unix := unix()

	// verify Unix
	assert.NotEqual(t, 0, unix)
}

func TestUnixStr(t *testing.T) {
	unixStr := unixStr()
	assert.NotEqual(t, "", unixStr)
	t.Logf("unixStr %v", unixStr)
}

func TestYear(t *testing.T) {
	year := year()
	assert.NotEqual(t, 0, year)
	t.Logf("year %v", year)
}

func TestMonth(t *testing.T) {
	month := month()
	assert.NotEqual(t, 0, month)
	t.Logf("month %v", month)
}

func TestHour(t *testing.T) {
	hour := hour()
	assert.NotEqual(t, 0, hour)
	t.Logf("hour %v", hour)
}

func TestDaysOfMonth(t *testing.T) {
	year := year()
	month := month()
	days := daysOfMonth(year, month)
	assert.NotEqual(t, 0, days)
	t.Logf("days %v", days)
}
