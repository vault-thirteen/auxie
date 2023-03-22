package time

import (
	"testing"
	"time"

	"github.com/vault-thirteen/tester"
)

func Test_ToSecondStart(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	// Test #1.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:51.123Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:51.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToSecondStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:05.000Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:05.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToSecondStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:04.999Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:04.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToSecondStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToMinuteStart(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	// Test #1.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:51.123Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMinuteStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:00.000Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMinuteStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:47:59.999Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:47:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMinuteStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToHourStart(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	// Test #1.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:51.123Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToHourStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:00:00.000Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToHourStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T11:59:59.999Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T11:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToHourStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToDayStart(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	// Test #1.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:51.123Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToDayStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-11T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToDayStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-10T23:59:59.999Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-10T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToDayStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToMonthStart(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	// Test #1.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:51.123Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-06-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2020-02-29T01:02:03.123Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2020-02-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToNextMonthStart(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	// Test #1.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:51.123Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-07-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToNextMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-07-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToNextMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2020-02-29T23:59:59.999Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2020-03-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToNextMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToPreviousMonthStart(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	// Test #1.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:51.123Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-05-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToPreviousMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2019-05-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToPreviousMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2020-02-29T23:59:59.999Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00.000Z")
	aTest.MustBeNoError(err)
	t1Processed = ToPreviousMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}
