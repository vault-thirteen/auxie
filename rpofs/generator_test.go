// generator_test.go.

package rpofs //nolint:testpackage

import (
	"strings"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_NewGenerator(t *testing.T) { //nolint:funlen
	type TestData struct {
		// Arguments.
		PasswordFixedSize      int
		AllowedPasswordSymbols []rune

		// State.
		ExpectedObjectState interface{}

		// Flags.
		IsErrorExpected bool
	}

	tests := []TestData{}

	// Test #1. Normal Data.
	tests = append(tests, TestData{
		PasswordFixedSize:      123,
		AllowedPasswordSymbols: []rune{'A', 'B', 'C'},
		ExpectedObjectState: &Generator{
			passwordLength: 123,
			allowedSymbols: []rune{'A', 'B', 'C'},
		},
		IsErrorExpected: false,
	})

	// Test #2. Password Length is too small.
	tests = append(tests, TestData{
		PasswordFixedSize:      -1,
		AllowedPasswordSymbols: []rune{'A', 'B', 'C'},
		ExpectedObjectState:    (*Generator)(nil),
		IsErrorExpected:        true,
	})

	// Test #3. Password Length is too big.
	tests = append(tests, TestData{
		PasswordFixedSize:      9999999,
		AllowedPasswordSymbols: []rune{'A', 'B', 'C'},
		ExpectedObjectState:    (*Generator)(nil),
		IsErrorExpected:        true,
	})

	// Test #4. Allowed Symbols Set is empty.
	tests = append(tests, TestData{
		PasswordFixedSize:      1,
		AllowedPasswordSymbols: []rune{},
		ExpectedObjectState:    (*Generator)(nil),
		IsErrorExpected:        true,
	})

	// Test #5. Allowed Symbols Set is not unique.
	tests = append(tests, TestData{
		PasswordFixedSize:      1,
		AllowedPasswordSymbols: []rune{'A', 'B', 'C', 'B'},
		ExpectedObjectState:    (*Generator)(nil),
		IsErrorExpected:        true,
	})

	var (
		aTest  = tester.New(t)
		result interface{}
		err    error
	)

	for _, test := range tests {
		result, err = NewGenerator(
			test.PasswordFixedSize,
			test.AllowedPasswordSymbols,
		)

		if test.IsErrorExpected {
			aTest.MustBeAnError(err)
			aTest.MustBeEqual(result, (*Generator)(nil))
		} else {
			aTest.MustBeNoError(err)
			aTest.MustBeEqual(result, test.ExpectedObjectState)
		}
	}
}

func Test_CreatePassword(t *testing.T) { //nolint:funlen
	const (
		PasswordLength  = 16
		IterationsCount = 10
	)

	type TestData struct {
		ExpectedPasswordLength int

		// Flags.
		IsErrorExpected bool
	}

	tests := []TestData{}

	// Test. Normal Data.
	tests = append(tests, TestData{
		ExpectedPasswordLength: PasswordLength,
		IsErrorExpected:        false,
	})

	var (
		aTest              = tester.New(t)
		result             *string
		err                error
		generator          *Generator
		generatedPasswords = make(map[string]bool)
		itemExists         bool
		symbolsSet         = []rune{'A', 'B', 'C'}
	)

	generator, err = NewGenerator(PasswordLength, symbolsSet)
	aTest.MustBeNoError(err)

	for _, test := range tests {
		for i := 0; i < IterationsCount; i++ {
			result, err = generator.CreatePassword()

			if test.IsErrorExpected {
				aTest.MustBeAnError(err)
				aTest.MustBeEqual(result, nil)
			} else {
				aTest.MustBeNoError(err)
				aTest.MustBeDifferent(result, nil)
				aTest.MustBeEqual(len(*result), test.ExpectedPasswordLength)
			}

			// Check the "randomness".
			_, itemExists = generatedPasswords[*result]
			if itemExists {
				t.Error("uniqueness failure")
				t.FailNow()
			}

			// Check the Symbols.
			if !stringContainsAllSymbolsFromSet(*result, symbolsSet) {
				t.Error("symbols are not correct")
				t.FailNow()
			}

			generatedPasswords[*result] = true
		}
	}
}

func stringContainsAllSymbolsFromSet(
	s string,
	symbolsSet []rune,
) bool {
	for _, symbol := range symbolsSet {
		if !stringContainsSymbol(s, symbol) {
			return false
		}
	}

	return true
}

func stringContainsSymbol(
	s string,
	symbol rune,
) bool {
	return strings.Contains(s, string(symbol))
}
