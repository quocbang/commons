package mes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_adjustEmployeeWorkDate(t *testing.T) {
	tmFormat := "2006-01-02 15:04:05"
	// 2020-10-10 15:12:33 local timezone
	{
		tm, err := time.ParseInLocation(tmFormat, "2020-10-10 15:12:33", time.Local)
		assert.NoError(t, err)
		tm = adjustEmployeeWorkDate(tm.Local())

		expected := time.Date(2020, 10, 10, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, tm)
		assert.Equal(t, "2020-10-10 00:00:00", tm.Format(tmFormat))
	}
	// 2020-01-01 07:59:59 local timezone
	{
		tm, err := time.ParseInLocation(tmFormat, "2020-01-01 07:59:59", time.Local)
		assert.NoError(t, err)
		tm = adjustEmployeeWorkDate(tm.Local())

		expected := time.Date(2019, 12, 31, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, tm)
		assert.Equal(t, "2019-12-31 00:00:00", tm.Format(tmFormat))
	}
	// 2020-01-01 08:00:00 local timezone
	{
		tm, err := time.ParseInLocation(tmFormat, "2020-01-01 08:00:00", time.Local)
		assert.NoError(t, err)
		tm = adjustEmployeeWorkDate(tm.Local())

		expected := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, tm)
		assert.Equal(t, "2020-01-01 00:00:00", tm.Format(tmFormat))
	}
	// 2020-09-09 00:54:11 local timezone
	{
		tm, err := time.ParseInLocation(tmFormat, "2020-09-09 00:54:11", time.Local)
		assert.NoError(t, err)
		tm = adjustEmployeeWorkDate(tm.Local())

		expected := time.Date(2020, 9, 8, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expected, tm)
		assert.Equal(t, "2020-09-08 00:00:00", tm.Format(tmFormat))
	}
}
