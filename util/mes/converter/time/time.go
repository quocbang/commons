package time

import (
	"fmt"
	"time"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/v2/commons"
)

// FromTime returns Time corresponding to the time.Time.
func FromTime(t time.Time) *commons.Time {
	return &commons.Time{Nano: t.UnixNano()}
}

// Parse parses a formatted string and returns Time value it represents.
//
// In the absence of a timezone indicator, Parse parses a formatted
// string as UTC.
func Parse(layout, value string) (*commons.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}
	return FromTime(t), nil
}

// ParseInLocation parses a formatted string and returns Time value it
// represents.
//
// In the absence of a timezone indicator, ParseInLocation uses the
// given location.
//
// ParseInLocation panics if loc is nil.
func ParseInLocation(layout, value string, loc *time.Location) (*commons.Time, error) {
	if loc == nil {
		return nil, fmt.Errorf("missing Location in call to ParseInLocation")
	}
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return nil, err
	}
	return FromTime(t), nil
}

// ToTime returns the local time.Time.
func ToTime(t *commons.Time) time.Time {
	return time.Unix(0, t.GetNano())
}

// Format returns a textual representation of the time value formatted
// according to layout. Format uses local timezone if loc is nil.
func Format(t *commons.Time, loc *time.Location, layout string) string {
	tm := ToTime(t)
	if loc != nil {
		tm = tm.In(loc)
	}
	return tm.Format(layout)
}
