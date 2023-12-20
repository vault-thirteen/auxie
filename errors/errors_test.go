package errors

import (
	"errors"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_Combine(t *testing.T) {
	var aTest = tester.New(t)
	var err1 error
	var err2 error
	var errCombined error
	var errCombinedExpected error

	// Test #1. ++
	err1 = errors.New("1")
	err2 = errors.New("2")
	errCombinedExpected = errors.New("1" + ErrorsSeparator + "2")
	errCombined = Combine(err1, err2)
	aTest.MustBeEqual(errCombined, errCombinedExpected)

	// Test #2. +-
	err1 = errors.New("1")
	err2 = nil
	errCombinedExpected = err1
	errCombined = Combine(err1, err2)
	aTest.MustBeEqual(errCombined, errCombinedExpected)

	// Test #3. -+
	err1 = nil
	err2 = errors.New("2")
	errCombinedExpected = err2
	errCombined = Combine(err1, err2)
	aTest.MustBeEqual(errCombined, errCombinedExpected)

	// Test #4. --
	err1 = nil
	err2 = nil
	errCombinedExpected = nil
	errCombined = Combine(err1, err2)
	aTest.MustBeEqual(errCombined, errCombinedExpected)
}
