package tsb

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_IsYes(t *testing.T) {
	aTest := tester.New(t)

	// Test the values.
	aTest.MustBeEqual(Yes.IsYes(), true)
	aTest.MustBeEqual(No.IsYes(), false)
	aTest.MustBeEqual(Maybe.IsYes(), false)

	// Test the junk.
	aTest.MustBeEqual(TSB(0).IsYes(), false)
	for i := 4; i <= 255; i++ {
		aTest.MustBeEqual(TSB(i).IsYes(), false)
	}
}

func Test_IsNo(t *testing.T) {
	aTest := tester.New(t)

	// Test the values.
	aTest.MustBeEqual(Yes.IsNo(), false)
	aTest.MustBeEqual(No.IsNo(), true)
	aTest.MustBeEqual(Maybe.IsNo(), false)

	// Test the junk.
	aTest.MustBeEqual(TSB(0).IsNo(), false)
	for i := 4; i <= 255; i++ {
		aTest.MustBeEqual(TSB(i).IsNo(), false)
	}
}

func Test_IsMaybe(t *testing.T) {
	aTest := tester.New(t)

	// Test the values.
	aTest.MustBeEqual(Yes.IsMaybe(), false)
	aTest.MustBeEqual(No.IsMaybe(), false)
	aTest.MustBeEqual(Maybe.IsMaybe(), true)

	// Test the junk.
	aTest.MustBeEqual(TSB(0).IsMaybe(), false)
	for i := 4; i <= 255; i++ {
		aTest.MustBeEqual(TSB(i).IsMaybe(), false)
	}
}

func Test_IsSet(t *testing.T) {
	aTest := tester.New(t)

	// Test the values.
	aTest.MustBeEqual(Yes.IsSet(), true)
	aTest.MustBeEqual(No.IsSet(), true)
	aTest.MustBeEqual(Maybe.IsSet(), true)

	// Test the junk.
	aTest.MustBeEqual(TSB(0).IsSet(), false)
	for i := 4; i <= 255; i++ {
		aTest.MustBeEqual(TSB(i).IsSet(), false)
	}
}
