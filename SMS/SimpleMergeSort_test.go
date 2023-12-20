package sms

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_New(t *testing.T) {
	aTest := tester.New(t)

	// Test #1. Normal data.
	data := [][]ValueType{
		nil,
		{5, 1, 2},
		{},
		{3, 7, 2},
	}

	expectedResult := &SimpleMergeSorter{
		data: [][]ValueType{
			{1, 2, 5},
			{2, 3, 7},
		},
		cursors: nil,
	}

	result, err := New(data...)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, expectedResult)

	// Test #2. No data.
	data = [][]ValueType{
		nil,
		{},
	}

	mergeSorter, err := New(data...)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(mergeSorter, (*SimpleMergeSorter)(nil))
}

func Test_getNonEmptyArrays(t *testing.T) {
	aTest := tester.New(t)

	data := [][]ValueType{
		nil,
		{5, 1, 2},
		{},
		{3, 7, 2},
	}

	expectedResult := [][]ValueType{
		{5, 1, 2},
		{3, 7, 2},
	}

	result := getNonEmptyArrays(data...)
	aTest.MustBeEqual(result, expectedResult)
}

func Test_Sort(t *testing.T) {
	aTest := tester.New(t)

	var data [][]ValueType

	// Test #1. Normal data.
	data = [][]ValueType{
		nil,
		{5, 1, 2},
		{},
		{3, 7, 2},
	}

	mergeSorter, err := New(data...)
	aTest.MustBeNoError(err)
	aTest.MustBeDifferent(mergeSorter, (*SimpleMergeSorter)(nil))

	result := mergeSorter.Sort()
	aTest.MustBeEqual(result, []ValueType{1, 2, 2, 3, 5, 7})
}

func Test_newIntPointer(t *testing.T) {
	aTest := tester.New(t)

	result := newIntPointer(9)
	aTest.MustBeEqual(*result, 9)
}

func Test_findCursorWithLeastValue(t *testing.T) {
	aTest := tester.New(t)

	var (
		cursorIndex int
		err         error
		data        *SimpleMergeSorter
	)

	// Test #1. No least value exists.
	data = &SimpleMergeSorter{
		data: [][]ValueType{
			{9, 9},
			{9, 9},
		},
		cursors: []*int{nil, nil},
	}

	cursorIndex, err = data.findCursorWithLeastValue()
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(cursorIndex, -1)

	// Test #2. One value.
	data = &SimpleMergeSorter{
		data: [][]ValueType{
			{9, 9},
			{9, 9},
		},
		cursors: []*int{nil, newIntPointer(0)},
	}

	cursorIndex, err = data.findCursorWithLeastValue()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(cursorIndex, 1)

	// Test #3. Two values.
	data = &SimpleMergeSorter{
		data: [][]ValueType{
			{99, 88, 77},
			{99, 88, 77},
			{99, 88, 77},
		},
		cursors: []*int{
			newIntPointer(0), // -> 99.
			nil,
			newIntPointer(1), // -> 88.
		},
	}

	cursorIndex, err = data.findCursorWithLeastValue()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(cursorIndex, 2)

	cursor := data.cursors[cursorIndex]
	aTest.MustBeEqual(data.data[cursorIndex][*cursor], ValueType(88))
}
