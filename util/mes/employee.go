package mes

import (
	"time"
)

// GenerateEmployeeWorkDate returns date of yesterday if before
// 8 a.m, otherwise today's date.
//
// The returned Time is in local timezone. Hours, minutes, seconds,
// and nanoseconds is truncated.
func GenerateEmployeeWorkDate() time.Time {
	tm := time.Now()
	return adjustEmployeeWorkDate(tm)
}

func adjustEmployeeWorkDate(tm time.Time) time.Time {
	if tm.Hour() < 8 {
		tm = tm.AddDate(0, 0, -1)
	}
	y, m, d := tm.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, tm.Location())
}
