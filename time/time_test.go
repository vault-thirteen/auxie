package time

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/vault-thirteen/tester"
)

const SecondsPerHour = 60 * 60

func Test_SelfCheck(t *testing.T) {
	var aTest = tester.New(t)
	var dur time.Duration

	// Check the minimal time unit.
	dur = time.Duration(1)
	aTest.MustBeEqual(dur, time.Nanosecond)

	// Check the scale.
	dur = time.Second
	aTest.MustBeEqual(dur, time.Duration(1_000_000_000))

	// Duration has int64 type.
	// This may change in the future, so we need to check this.
	var durAsInt64 int64

	// Biggest positive unit.
	dur = time.Duration(math.MaxInt64)
	durAsInt64 = int64(dur)
	aTest.MustBeEqual(time.Duration(durAsInt64), dur)

	// Smallest positive unit.
	dur = time.Duration(1)
	durAsInt64 = int64(dur)
	aTest.MustBeEqual(time.Duration(durAsInt64), dur)

	// Biggest negative unit.
	dur = time.Duration(math.MinInt64)
	durAsInt64 = int64(dur)
	aTest.MustBeEqual(time.Duration(durAsInt64), dur)

	// Smallest negative unit.
	dur = time.Duration(-1)
	durAsInt64 = int64(dur)
	aTest.MustBeEqual(time.Duration(durAsInt64), dur)
}

func Test_HoursToDuration(t *testing.T) {
	var aTest = tester.New(t)
	var hours = big.NewRat(1, 3) // 1/3 of an hour.

	dur := HoursToDuration(hours)
	aTest.MustBeEqual(dur, time.Minute*20)
	aTest.MustBeEqual(dur, time.Second*1200)
	aTest.MustBeEqual(dur, time.Nanosecond*1200*1_000_000_000)
}

// Attention ! Some of these functions have been created due to the dumbness of
// the developers of the Go programming language. The results of the tests may
// change in later versions of the Go programming language. The author of these
// tests does not guarantee the correctness !
