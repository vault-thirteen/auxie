package saac

import (
	"math/big"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_New(t *testing.T) {
	aTest := tester.New(t)

	var (
		calculator *Calculator
	)

	// Test.
	calculator = New()
	aTest.MustBeDifferent(calculator, (*Calculator)(nil))
	aTest.MustBeEqual(calculator.n.String(), "0")
	aTest.MustBeEqual(calculator.previousValue, (*big.Float)(nil))
	aTest.MustBeEqual(calculator.one.String(), "1")
}

func Test_AddItemAndGetAverage(t *testing.T) {
	aTest := tester.New(t)

	var (
		calculator *Calculator
		avg        string
		err        error
	)

	calculator = New()

	// Test #1. Bad number.
	avg, err = calculator.AddItemAndGetAverage("x")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(avg, "")

	// Test #2. "1".
	avg, err = calculator.AddItemAndGetAverage("1")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(avg, "1")

	// Test #3. "1" + "2" -> "1.5".
	avg, err = calculator.AddItemAndGetAverage("2")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(avg, "1.5")

	// Test #4. "1" + "2" + "30" -> "11".
	avg, err = calculator.AddItemAndGetAverage("30")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(avg, "11")

	// Test #5. "1" + "2" + "30" + "4" -> 37/4 -> 9 1/4.
	avg, err = calculator.AddItemAndGetAverage("4")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(avg, "9.25")

	// Test #6. "1" + "2" + "30" + "4" + "3" -> 40/5 -> 8.
	avg, err = calculator.AddItemAndGetAverage("3")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(avg, "8")
}
