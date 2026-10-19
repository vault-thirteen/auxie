package nts

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ByteArrayToStrings(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result []string

	// Test #1.
	result, err = ByteArrayToStrings([]byte{})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, []string(nil))

	// Test #2.
	result, err = ByteArrayToStrings([]byte{1, 2, 3})
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, []string(nil))

	// Test #3.
	result, err = ByteArrayToStrings([]byte{'T', 'E', 'S', 'T', NUL})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, []string{"TEST"})

	// Test #4.
	result, err = ByteArrayToStrings([]byte{NUL})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, []string{""})

	// Test #5.
	result, err = ByteArrayToStrings([]byte{NUL, 'A', NUL})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, []string{"", "A"})

	// Test #6.
	result, err = ByteArrayToStrings(
		[]byte{NUL, 'A', NUL, 'B', 'C', NUL, NUL, NUL, 'D', NUL, NUL})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, []string{"", "A", "BC", "", "", "D", ""})
}
