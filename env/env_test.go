// env_test.go.

package env

import (
	"os"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_GetEnv(t *testing.T) {

	var aTest *tester.Test
	var err error
	var result string

	aTest = tester.New(t)

	// Test #1. Normal Data.
	err = os.Setenv("TEST_ENV_A", "XYZ")
	aTest.MustBeNoError(err)
	result, err = GetEnv("TEST_ENV_A")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, "XYZ")

	// Test #2. Empty Variable.
	// Ensure that it is really empty before this Test.
	aTest.MustBeEqual(len(os.Getenv("TEST_ENV_B")), 0)
	result, err = GetEnv("TEST_ENV_B")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, "")
}
