package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const layout = "2006-01-02 15:04:05"

var (
	tzTaipei = time.FixedZone("Asia/Taipei", int(8*time.Hour/time.Second))
	tzTokyo  = time.FixedZone("Asia/Tokyo", int(9*time.Hour/time.Second))
)

func TestParse(t *testing.T) {
	// in specified timezone indicator
	{
		tm, err := Parse(time.RFC3339, "2020-08-10T11:16:10+09:00")
		assert.NoError(t, err)
		expected := FromTime(time.Date(2020, time.August, 10, 11, 16, 10, 0, tzTokyo))
		assert.Equal(t, expected, tm)
	}
	// in the absence of a timezone indicator
	{
		tm, err := Parse(layout, "2020-08-10 11:16:10")
		assert.NoError(t, err)
		expected := FromTime(time.Date(2020, time.August, 10, 11, 16, 10, 0, time.UTC))
		assert.Equal(t, expected, tm)
	}
}

func TestParseInLocation(t *testing.T) {
	// in specified timezone indicator
	{
		tm, err := ParseInLocation(time.RFC3339, "2020-08-10T11:16:10+09:00", tzTaipei)
		assert.NoError(t, err)
		expected := FromTime(time.Date(2020, time.August, 10, 11, 16, 10, 0, tzTokyo))
		assert.Equal(t, expected, tm)
	}
	// in the absence of a timezone indicator
	{
		tm, err := ParseInLocation(layout, "2020-08-10 11:16:10", tzTaipei)
		assert.NoError(t, err)
		expected := FromTime(time.Date(2020, time.August, 10, 11, 16, 10, 0, tzTaipei))
		assert.Equal(t, expected, tm)
	}
	// with nil loc
	{
		tm, err := ParseInLocation(layout, "2020-08-10 11:16:10", nil)
		assert.Error(t, err)
		assert.Empty(t, tm)
	}
}

func TestToTime(t *testing.T) {
	{
		now := time.Now()
		mesTime := FromTime(now)
		tt := ToTime(mesTime)
		assert.True(t, now.Format(layout) == tt.Format(layout))
	}
}

func TestFormat(t *testing.T) {
	tm := time.Date(2020, time.August, 10, 11, 16, 10, 0, tzTokyo)
	mesTime := FromTime(tm)

	// in specified timezone indicator
	{
		assert.Equal(t, "2020-08-10T10:16:10+08:00", Format(mesTime, tzTaipei, time.RFC3339))

		// without loc
		assert.Equal(t, tm.Local().Format(time.RFC3339), Format(mesTime, nil, time.RFC3339))
	}
	// in the absence of a timezone indicator
	{
		assert.Equal(t, "2020-08-10 10:16:10", Format(mesTime, tzTaipei, layout))

		// without loc
		assert.Equal(t, tm.Local().Format(layout), Format(mesTime, nil, layout))
	}
}
