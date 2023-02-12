package rng

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_New(t *testing.T) {
	aTest := tester.New(t)
	var rng *Range
	var err error

	// Test #1.
	rng, err = New(6, 5)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), "start is reversed")
	aTest.MustBeEqual(rng, (*Range)(nil))

	// Test #2.
	rng, err = New(5, 6)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(rng, &Range{
		start:  5,
		end:    6,
		middle: 5.5,
		radius: 0.5,
	})
}

func Test_Contains(t *testing.T) {
	aTest := tester.New(t)
	var rng *Range
	var err error

	rng, err = New(10, 20)
	aTest.MustBeNoError(err)

	// Tests.
	aTest.MustBeEqual(rng.Contains(1), false)
	aTest.MustBeEqual(rng.Contains(10), true)
	aTest.MustBeEqual(rng.Contains(15), true)
	aTest.MustBeEqual(rng.Contains(20), true)
	aTest.MustBeEqual(rng.Contains(29), false)
}

func Test_GetStart(t *testing.T) {
	aTest := tester.New(t)
	var rng *Range
	var err error

	rng, err = New(10, 20)
	aTest.MustBeNoError(err)

	// Test.
	aTest.MustBeEqual(rng.GetStart(), 10.0)
}

func Test_GetEnd(t *testing.T) {
	aTest := tester.New(t)
	var rng *Range
	var err error

	rng, err = New(10, 20)
	aTest.MustBeNoError(err)

	// Test.
	aTest.MustBeEqual(rng.GetEnd(), 20.0)
}

func Test_GetMiddle(t *testing.T) {
	aTest := tester.New(t)
	var rng *Range
	var err error

	rng, err = New(20, 30)
	aTest.MustBeNoError(err)

	// Test.
	aTest.MustBeEqual(rng.GetMiddle(), 25.0)
}

func Test_GetRadius(t *testing.T) {
	aTest := tester.New(t)
	var rng *Range
	var err error

	rng, err = New(20, 30)
	aTest.MustBeNoError(err)

	// Test.
	aTest.MustBeEqual(rng.GetRadius(), 5.0)
}

func Test_HasIntersectionWith(t *testing.T) {
	aTest := tester.New(t)
	var rngA, rngB *Range
	var err error

	rngA, err = New(10, 20)
	aTest.MustBeNoError(err)

	type Data struct {
		rngBStart      float64
		rngBEnd        float64
		expectedResult bool
	}

	NewData := func(start, end float64, expectedResult bool) (d Data) {
		return Data{
			rngBStart:      start,
			rngBEnd:        end,
			expectedResult: expectedResult,
		}
	}

	tests := []Data{
		NewData(1, 2, false),
		NewData(1, 10, true),
		NewData(1, 15, true),
		NewData(1, 20, true),
		NewData(1, 29, true),

		NewData(10, 10, true),
		NewData(10, 15, true),
		NewData(10, 20, true),
		NewData(10, 29, true),

		NewData(15, 16, true),
		NewData(15, 20, true),
		NewData(15, 29, true),

		NewData(20, 20, true),
		NewData(20, 29, true),

		NewData(29, 30, false),
	}

	for _, tst := range tests {
		rngB, err = New(tst.rngBStart, tst.rngBEnd)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(rngA.HasIntersectionWith(rngB), tst.expectedResult)
	}
}

func Test_IsSequence(t *testing.T) {
	aTest := tester.New(t)
	var rngA, rngB, rngC *Range
	var err error

	rngA, err = New(10, 15)
	aTest.MustBeNoError(err)
	rngB, err = New(15, 40)
	aTest.MustBeNoError(err)
	rngC, err = New(20, 40)
	aTest.MustBeNoError(err)

	// Tests.
	aTest.MustBeEqual(IsSequence(rngA, rngB), true)
	aTest.MustBeEqual(IsSequence(rngA, rngC), false)
}
