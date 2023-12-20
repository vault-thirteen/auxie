package time

import (
	"fmt"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_GetLocationOffset(t *testing.T) {
	var tst = tester.New(t)
	var err error
	var location *time.Location
	var locationOffset time.Duration
	var locationOffsetSecExpected float64

	// Test #1. Positive Time Zone Offset.
	locationOffsetSecExpected = 5 * SecondsPerHour // 5 Earth-Planet Hours.
	location = time.FixedZone("Somewhere on Earth [EAST]", 5*SecondsPerHour)
	locationOffset, err = GetLocationOffset(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffset.Seconds(), locationOffsetSecExpected)

	// Test #2. Negative Time Zone Offset.
	locationOffsetSecExpected = -4 * SecondsPerHour // 4 Earth-Planet Hours.
	location = time.FixedZone("Somewhere on Earth [WEST]", -4*SecondsPerHour)
	locationOffset, err = GetLocationOffset(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffset.Seconds(), locationOffsetSecExpected)

	// Test #3. Zero Time Zone Offset.
	locationOffsetSecExpected = 0
	location = time.FixedZone("Somewhere on Earth [Greenwich]", 0)
	locationOffset, err = GetLocationOffset(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffset.Seconds(), locationOffsetSecExpected)
}

func Test_ParseDayTimeStringInLocation(t *testing.T) {
	var tst = tester.New(t)
	var dayStartTime time.Time
	var dayStartTimeExpected time.Time
	var err error
	var location *time.Location

	// Test #1. Positive Time Zone Offset.
	location = time.FixedZone("+0400", 4*SecondsPerHour)
	dayStartTimeExpected, err = time.Parse(time.RFC3339, "2019-09-01T00:00:00+04:00")
	tst.MustBeNoError(err)
	dayStartTime, err = ParseDayTimeStringInLocation("2019-09-01", location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(
		dayStartTime.Format(time.RFC3339),
		dayStartTimeExpected.Format(time.RFC3339),
	)
	fmt.Println(dayStartTime.Format(time.RFC3339)) // Report.

	// Test #2. Negative Time Zone Offset.
	location = time.FixedZone("-0800", -8*SecondsPerHour)
	dayStartTimeExpected, err = time.Parse(time.RFC3339, "2019-09-01T00:00:00-08:00")
	tst.MustBeNoError(err)
	dayStartTime, err = ParseDayTimeStringInLocation("2019-09-01", location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(
		dayStartTime.Format(time.RFC3339),
		dayStartTimeExpected.Format(time.RFC3339),
	)
	fmt.Println(dayStartTime.Format(time.RFC3339)) // Report.

	// Test #3. Zero Time Zone Offset.
	location = time.FixedZone("", 0)
	dayStartTimeExpected, err = time.Parse(time.RFC3339, "2019-09-01T00:00:00-00:00")
	tst.MustBeNoError(err)
	dayStartTime, err = ParseDayTimeStringInLocation("2019-09-01", location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(
		dayStartTime.Format(time.RFC3339),
		dayStartTimeExpected.Format(time.RFC3339),
	)
	fmt.Println(dayStartTime.Format(time.RFC3339)) // Report.
}
