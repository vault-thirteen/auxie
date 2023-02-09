package sma

import (
	"container/list"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_New(t *testing.T) {
	aTest := tester.New(t)

	var (
		calculator *Calculator
		err        error
	)

	// Test #1. Wrong windows size.
	calculator, err = New(0)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(calculator, (*Calculator)(nil))

	// Test #2. OK.
	calculator, err = New(123)
	aTest.MustBeNoError(err)
	aTest.MustBeDifferent(calculator, (*Calculator)(nil))
	aTest.MustBeEqual(calculator.size, 123)
	aTest.MustBeDifferent(calculator.items, (*list.List)(nil))
	aTest.MustBeEqual(calculator.previousValue, (*float64)(nil))
}

func Test_IsCold(t *testing.T) {
	aTest := tester.New(t)

	var (
		calculator *Calculator
		err        error
	)

	calculator, err = New(3)
	aTest.MustBeNoError(err)

	// Test #0. No items.
	aTest.MustBeEqual(true, calculator.IsCold())

	// Test #1. One item.
	calculator.items.PushFront(ValueType(1))
	aTest.MustBeEqual(true, calculator.IsCold())

	// Test #2. Two items.
	calculator.items.PushFront(ValueType(2))
	aTest.MustBeEqual(true, calculator.IsCold())

	// Test #3. Three items.
	calculator.items.PushFront(ValueType(3))
	aTest.MustBeEqual(false, calculator.IsCold())

	// Test #4. Four items.
	calculator.items.PushFront(ValueType(4))
	aTest.MustBeEqual(false, calculator.IsCold())
}

func Test_AddItemAndGetSMA(t *testing.T) {
	aTest := tester.New(t)

	var (
		calculator *Calculator
		sma        ValueType
		err        error
	)

	calculator, err = New(3)
	aTest.MustBeNoError(err)

	// Test #1. One item.
	sma, err = calculator.AddItemAndGetSMA(ValueType(8))
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(sma, ValueType(0))

	// Test #2. Two items.
	sma, err = calculator.AddItemAndGetSMA(ValueType(9))
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(sma, ValueType(0))

	// Test #3. Three items.
	sma, err = calculator.AddItemAndGetSMA(ValueType(10))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(sma, ValueType(9))

	// Test #4. Four items.
	sma, err = calculator.AddItemAndGetSMA(ValueType(314))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(sma, ValueType(111)) // 9+10+314=333.

	// Test #5. Emulation of bad average calculation for full coverage.
	calculator.items.PushFront(ValueType(99))
	calculator.previousValue = nil
	sma, err = calculator.AddItemAndGetSMA(ValueType(5))
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(sma, ValueType(0))
}

func Test_getItemsAverage(t *testing.T) {
	aTest := tester.New(t)

	var (
		calculator *Calculator
		average    ValueType
		err        error
	)

	calculator, err = New(3)
	aTest.MustBeNoError(err)

	calculator.items.PushFront(ValueType(1.0))
	calculator.items.PushFront(ValueType(2.0))

	// Test #1. Not enough Data.
	average, err = calculator.getItemsAverage()
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(average, ValueType(0))

	// Test #2. OK.
	calculator.items.PushFront(ValueType(12.0))

	average, err = calculator.getItemsAverage()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(average, ValueType(5))
}
