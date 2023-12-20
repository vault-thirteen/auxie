package time

import (
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_IsEmpty(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var x time.Time

	// Test #1.
	x, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(IsEmpty(x), false)

	// Test #2.
	x = time.Time{}
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(IsEmpty(x), true)
}
