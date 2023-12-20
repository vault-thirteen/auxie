package tester

import (
	"errors"
	"testing"
)

func Test_errorIsSet(t *testing.T) {
	var result bool

	// Test #1. Negative.
	result = errorIsSet(nil)
	if result != false {
		t.FailNow()
	}

	// Test #2. Positive.
	result = errorIsSet(errors.New("some error"))
	if result != true {
		t.FailNow()
	}
}

func Test_errorIsEmpty(t *testing.T) {
	var result bool

	// Test #1. Negative.
	result = errorIsEmpty(errors.New("some error"))
	if result != false {
		t.FailNow()
	}

	// Test #2. Positive.
	result = errorIsEmpty(nil)
	if result != true {
		t.FailNow()
	}
}

func Test_interfacesAreEqual(t *testing.T) {
	var result bool

	type TestTypeX struct {
		Age  int
		Name string
	}

	// Test #1. Negative.
	result = interfacesAreEqual(
		TestTypeX{Age: 10, Name: "John"},
		TestTypeX{Age: 11, Name: "Jack"},
	)
	if result != false {
		t.FailNow()
	}

	// Test #2. Positive.
	result = interfacesAreEqual(
		TestTypeX{Age: 12, Name: "Meredith"},
		TestTypeX{Age: 12, Name: "Meredith"},
	)
	if result != true {
		t.FailNow()
	}
}

func Test_interfacesAreDifferent(t *testing.T) {
	var result bool

	type TestTypeX struct {
		Age  int
		Name string
	}

	// Test #1. Negative.
	result = interfacesAreDifferent(
		TestTypeX{Age: 12, Name: "Meredith"},
		TestTypeX{Age: 12, Name: "Meredith"},
	)
	if result != false {
		t.FailNow()
	}

	// Test #2. Positive.
	result = interfacesAreDifferent(
		TestTypeX{Age: 10, Name: "John"},
		TestTypeX{Age: 11, Name: "Jack"},
	)
	if result != true {
		t.FailNow()
	}
}
