package time

import (
	"math/big"
	"testing"
	"time"

	"github.com/vault-thirteen/tester"
)

func Test_Minimum(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result time.Time
	var resultExpected time.Time
	var time1 time.Time
	var time2 time.Time

	time1, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	time2, err = time.Parse(time.RFC3339, "2020-06-24T15:02:55Z")
	aTest.MustBeNoError(err)

	// Test #1.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time1, time2)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time2, time1)
	aTest.MustBeEqual(result, resultExpected)

	// Test #3.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time1, time1)
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Maximum(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result time.Time
	var resultExpected time.Time
	var time1 time.Time
	var time2 time.Time

	time1, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	time2, err = time.Parse(time.RFC3339, "2020-06-24T15:02:55Z")
	aTest.MustBeNoError(err)

	// Test #1.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time1, time2)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time2, time1)
	aTest.MustBeEqual(result, resultExpected)

	// Test #3.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time2, time2)
	aTest.MustBeEqual(result, resultExpected)
}

func Test_AddHours(t *testing.T) {
	var aTest = tester.New(t)
	var t1 time.Time
	var t2 time.Time

	t1 = time.Now()
	t2 = t1.Add(time.Minute * 20)
	t1PlusDelta := AddHours(t1, big.NewRat(1, 3))
	aTest.MustBeEqual(t1PlusDelta, t2)
}

func Test_SubHours(t *testing.T) {
	var aTest = tester.New(t)
	var t1 time.Time
	var t2 time.Time

	t1 = time.Now()
	t2 = t1.Add(time.Minute * 20)
	t2MinusT1 := SubHours(t2, big.NewRat(1, 3))
	aTest.MustBeEqual(t2MinusT1, t1)
}

func Test_IntervalDurationHours(t *testing.T) {
	var aTest = tester.New(t)
	var deltaRat *big.Rat
	var t1 time.Time
	var t2 time.Time

	t1 = time.Now()
	t2 = t1.Add(time.Minute * 20)
	deltaRat = IntervalDurationHours(t1, t2)
	deltaFloat64, isDeltaExact := deltaRat.Float64()
	aTest.MustBeEqual(deltaFloat64, float64(0.3333333333333333))
	aTest.MustBeEqual(isDeltaExact, true)
	deltaDur := HoursToDuration(deltaRat)
	aTest.MustBeEqual(deltaDur.Minutes(), float64(20))
}
