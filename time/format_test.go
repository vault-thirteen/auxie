package time

import (
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewTimeStringRFC3339(t *testing.T) {
	var aTest = tester.New(t)
	var result string
	var resultExpected string

	resultExpected = "2019-06-11T16:44:05Z"

	// Test.
	result = NewTimeStringRFC3339(
		2019,
		6,
		11,
		16,
		44,
		5,
	)
	aTest.MustBeEqual(result, resultExpected)

	tm, err := time.Parse(time.RFC3339, result)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tm.Year(), 2019)
	aTest.MustBeEqual(tm.Month(), time.Month(6))
	aTest.MustBeEqual(tm.Day(), 11)
	aTest.MustBeEqual(tm.Hour(), 16)
	aTest.MustBeEqual(tm.Minute(), 44)
	aTest.MustBeEqual(tm.Second(), 5)
}
