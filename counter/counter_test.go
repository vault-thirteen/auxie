package c

import (
	"math/big"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewCounter(t *testing.T) {
	aTest := tester.New(t)
	c := NewCounter()
	aTest.MustBeEqual(c.bic.String(), "0")
}

func Test_RawValue(t *testing.T) {
	aTest := tester.New(t)
	c := &Counter{
		bic: big.NewInt(123),
	}
	aTest.MustBeEqual(c.RawValue().String(), "123")
}

func Test_ToString(t *testing.T) {
	aTest := tester.New(t)
	c := &Counter{
		bic: big.NewInt(1234567),
	}
	aTest.MustBeEqual(c.ToString(), "1234567")
}

func Test_Load(t *testing.T) {
	aTest := tester.New(t)
	c := NewCounter()
	var err error

	// Test #1. Bad text value.
	err = c.Load("junk")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), Err_Load)

	// Test #2. Good text value.
	err = c.Load("12345")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(c.bic.String(), "12345")

	// Test #3. Good text value of great size.
	err = c.Load("1000000000000000000000000000000000000000000000")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(c.bic.String(), "1000000000000000000000000000000000000000000000")
}

func Test_Inc(t *testing.T) {
	aTest := tester.New(t)
	var err error

	// Test #1. Small value.
	c := &Counter{
		bic: big.NewInt(1234567),
	}
	c.Inc()
	aTest.MustBeEqual(c.ToString(), "1234568")

	// Test #2. Big value.
	c = NewCounter()
	x := "1000000000000000000000000000000000000000000000"
	err = c.Load(x + "0")
	aTest.MustBeNoError(err)
	c.Inc()
	aTest.MustBeEqual(c.ToString(), x+"1")

	// Test #3. Very big value.
	manyZeroes := "00000000000000000000000000000000000000000000000000000000000000"
	c = NewCounter()
	err = c.Load(x + manyZeroes + "0")
	aTest.MustBeNoError(err)
	c.Inc()
	aTest.MustBeEqual(c.ToString(), x+manyZeroes+"1")

	// Test #4. Huge value.
	base := x + manyZeroes + manyZeroes + manyZeroes + manyZeroes + manyZeroes + manyZeroes
	text := base + "0"
	textExpected := base + "1"
	c = NewCounter()
	err = c.Load(text)
	aTest.MustBeNoError(err)
	c.Inc()
	aTest.MustBeEqual(c.ToString(), textExpected)
}
