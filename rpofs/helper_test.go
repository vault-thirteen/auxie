package rpofs

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_isSetOfUniqueSymbols(t *testing.T) {
	type TestData struct {
		Set            []rune
		ResultExpected bool
	}

	tests := []TestData{}

	// Test #1. Empty Set.
	tests = append(tests, TestData{
		Set:            []rune{},
		ResultExpected: false,
	})

	// Test #2. Set with a single Item
	tests = append(tests, TestData{
		Set:            []rune{'A'},
		ResultExpected: true,
	})

	// Test #3. Set with unique Symbols
	tests = append(tests, TestData{
		Set:            []rune{'J', 'A', 'C', 'K'},
		ResultExpected: true,
	})

	// Test #4. Set with non-unique Symbols
	tests = append(tests, TestData{
		Set:            []rune{'J', 'A', 'C', 'K', 'C'},
		ResultExpected: false,
	})

	var (
		aTest  = tester.New(t)
		result bool
	)

	for _, test := range tests {
		result = isSetOfUniqueSymbols(test.Set)
		aTest.MustBeEqual(result, test.ResultExpected)
	}
}
