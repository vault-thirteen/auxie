package time

import (
	"math"
	"math/big"
	"time"
)

// FormatDayTimeString is the format of day time, i.e. a day identifier.
const FormatDayTimeString = "2006-01-02"

// Settings for internal usage.
const (
	// TimeFormatFor_GetLocationOffset is the time format for the
	// 'GetLocationOffset' function.
	TimeFormatFor_GetLocationOffset = "2006-01-02 15:04:05"
)

// Variables for internal usage.
var (
	sampleTimeString = "2019-09-01 00:00:00"
	emptyTime        time.Time
	dayDuration      = time.Hour * 24
)

// HoursToDuration converts hours into duration.
// Why ? In Go language it is impossible to simply get a 1/3 of an hour.
// Moreover, a float32 number storing a 1/3 will always have an error due to
// its binary nature.
func HoursToDuration(hours *big.Rat) time.Duration {
	// What are we doing ?
	// hours.Num() * time.Hour / hours.Denom().
	x := big.NewRat(time.Hour.Nanoseconds(), 1) // X = 1 Hour.
	x.Mul(x, hours)                             // X = X * hours.
	xFloat64, _ := x.Float64()                  // X -> float64.
	return time.Duration(math.Round(xFloat64))
}
