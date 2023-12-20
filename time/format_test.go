package time

import (
	"testing"

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
}
