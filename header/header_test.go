package header

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_MakeListOfHeaders(t *testing.T) {
	var aTest = tester.New(t)
	var result string
	var headers []string

	// Test #1. One Item.
	headers = []string{"a"}
	result = MakeListOfHeaders(headers)
	aTest.MustBeEqual(result, "a")

	// Test #2. Two Items.
	headers = []string{"aa", "bb"}
	result = MakeListOfHeaders(headers)
	aTest.MustBeEqual(result, "aa, bb")

	// Test #3. Three Items.
	headers = []string{"a", "bb", "ccc"}
	result = MakeListOfHeaders(headers)
	aTest.MustBeEqual(result, "a, bb, ccc")

	// Test #4. No Items.
	headers = []string{}
	result = MakeListOfHeaders(headers)
	aTest.MustBeEqual(result, "")
}
