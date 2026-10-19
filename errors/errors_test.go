package errors

import (
	"errors"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
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

func Test_ErrorFromText(t *testing.T) {
	var aTest = tester.New(t)
	var err error

	// Test #1. Empty text.
	err = ErrorFromText(" ")
	aTest.MustBeEqual(err, nil)

	// Test #2. Non-empty text.
	err = ErrorFromText("  some text  ")
	aTest.MustBeDifferent(err, nil)
	aTest.MustBeEqual(err.Error(), "some text")
}

func Test_ErrorFromFormattedText(t *testing.T) {
	var aTest = tester.New(t)
	var err error

	// Test #1. Empty text.
	err = ErrorFromFormattedText(" ")
	aTest.MustBeEqual(err, nil)

	// Test #2. Non-empty text.
	err = ErrorFromFormattedText("  some text: %v  ", 123)
	aTest.MustBeDifferent(err, nil)
	aTest.MustBeEqual(err.Error(), "some text: 123")
}

func Test_ErrorWithPrefix(t *testing.T) {
	var aTest = tester.New(t)
	var err1 error
	var err2 error

	// Test #1. Empty error.
	err1 = nil
	err2 = ErrorWithPrefix(err1, " prefix ")
	aTest.MustBeEqual(err2, nil)

	// Test #2. Empty prefix.
	err1 = errors.New("abc")
	err2 = ErrorWithPrefix(err1, "   ")
	aTest.MustBeDifferent(err2, nil)
	aTest.MustBeEqual(err2.Error(), "abc")

	// Test #3. Normal situation.
	err1 = errors.New("xyz")
	err2 = ErrorWithPrefix(err1, " some error  ")
	aTest.MustBeDifferent(err2, nil)
	aTest.MustBeEqual(err2.Error(), "some error: xyz")
}
