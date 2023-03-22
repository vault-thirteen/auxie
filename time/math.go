package time

import (
	"math/big"
	"time"
)

// Minimum selects the minimum time of the two specified.
func Minimum(a time.Time, b time.Time) time.Time {
	if a.After(b) {
		return b
	}
	return a
}

// Maximum selects the maximum time of the two specified.
func Maximum(a time.Time, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

// AddHours adds hours to the time.
func AddHours(timeStart time.Time, deltaHours *big.Rat) time.Time {
	return timeStart.Add(HoursToDuration(deltaHours))
}

// SubHours subtracts hours from the time.
func SubHours(timeStart time.Time, deltaHours *big.Rat) time.Time {
	return timeStart.Add(-HoursToDuration(deltaHours))
}

// IntervalDurationHours calculates the interval duration between two specified
// time values in hours.
func IntervalDurationHours(timeStart time.Time, timeEnd time.Time) (delta *big.Rat) {
	delta = new(big.Rat)
	return delta.SetFloat64(timeEnd.Sub(timeStart).Hours())
}
