// Package tzutils provides utility functions for time zone conversions and operations. 
package tzutils

import (
	"time"
)

// CurrentTimeIn returns the current time in the specified time zone.
// The timezone parameter must be a valid IANA time zone, such as "America/New_York".
// Returns an error if the time zone is invalid.
func CurrentTimeIn(timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}

// ConvertTime converts a given time from one time zone to another.
// Parameters:
// - t: The time to be converted.
// - fromTZ: The source time zone (e.g., "America/New_York").
// - toTZ: The destination time zone (e.g., "Europe/London").
// Returns the converted time or an error if either time zone is invalid.
func ConvertTime(t time.Time, fromTZ, toTZ string) (time.Time, error) {
	fromLocation, err := time.LoadLocation(fromTZ)
	if err != nil {
		return time.Time{}, err
	}

	toLocation, err := time.LoadLocation(toTZ)
	if err != nil {
		return time.Time{}, err
	}

	utcTime := t.In(fromLocation).UTC()
	return utcTime.In(toLocation), nil
}

// CurrentTimeZone returns the current time zone's name and offset from UTC.
// Offset is in seconds.
func CurrentTimeZone() (string, int) {
	name, offset := time.Now().Local().Zone()
	return name, offset
}

// Offset returns the current offset in seconds from UTC for the specified time zone.
// The timezone parameter must be a valid IANA time zone.
func Offset(timezone string) (int, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}
	_, offset := time.Now().In(loc).Zone()
	return offset, nil
}