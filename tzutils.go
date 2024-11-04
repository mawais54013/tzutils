package tzutils

import (
	"time"
)

// CurrentTimeIn returns the time in the specified time zone.
func CurrentTimeIn(timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}

func ConvertTime(t time.Time, fromTZ, toTZ string) (time.Time, error) {
	fromLocation, err := time.LoadLocation(fromTZ)
	if err != nil {
		return time.Time{}, err
	}

	toLocation, err := time.LoadLocation(toTZ)
	if err != nil {
		return time.Time{}, err
	}

	// Convert to UTC first 
	utcTime := t.In(fromLocation).UTC()
	return utcTime.In(toLocation), nil
}

// CurrentTimeZone returns the current zone. 
func CurrentTimeZone() (string, int) {
	name, offset := time.Now().Local().Zone()
	return name, offset
}

// Offset returns the current offset for the specified time zone.
func Offset(timezone string) (int, error) {
    loc, err := time.LoadLocation(timezone)
    if err != nil {
        return 0, err
    }
    _, offset := time.Now().In(loc).Zone()
    return offset, nil
}
