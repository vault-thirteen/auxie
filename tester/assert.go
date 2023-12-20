package tester

import (
	"fmt"
	"reflect"

	"github.com/kr/pretty"
)

// Assertion methods.
//
// Assertion methods help in comparing values and checking errors. This package
// provides a convenient way to control mismatches. When the comparison fails,
// it shows not only the two values which failed, but also a 'Diff'
// (difference) between them.

// Error messages.
const (
	ErrErrorWasExpected    = "An Error was expected, but None was received"
	ErrfNoErrorWasExpected = "No Error was expected, but One was received: %v"
	ErrfNotEqual           = `Values should be equal, but they are not. 
A=%v 
B=%v 
Diff=%v`
	ErrfNotDifferent = `Values should be different, but they are not. 
A=%v 
B=%v 
Diff=%v`
)

func errorIsSet(err error) bool {
	return err != nil
}

func errorIsEmpty(err error) bool {
	return err == nil
}

func interfacesAreEqual(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}

func interfacesAreDifferent(a any, b any) bool {
	return !reflect.DeepEqual(a, b)
}

// MustBeAnError ensures that the error is not nil.
// If the error is nil, it stops the test.
func (test *Test) MustBeAnError(err error) {
	if errorIsEmpty(err) {
		test.t.Error(ErrErrorWasExpected)
		test.t.FailNow()
	}
}

// MustBeNoError ensures that the error is nil.
// If the error is not nil, it stops the test.
func (test *Test) MustBeNoError(err error) {
	if errorIsSet(err) {
		test.t.Errorf(ErrfNoErrorWasExpected, err)
		test.t.FailNow()
	}
}

// MustBeEqual ensures that two variables have equal values.
// If not, it stops the test.
func (test *Test) MustBeEqual(a any, b any) {
	if interfacesAreDifferent(a, b) {
		msg := fmt.Sprintf(
			ErrfNotEqual,
			pretty.Sprint(a),
			pretty.Sprint(b),
			pretty.Diff(a, b),
		)
		test.t.Errorf(msg)
		test.t.FailNow()
	}
}

// MustBeDifferent ensures that two variables have different values.
// If not, it stops the test.
func (test *Test) MustBeDifferent(a any, b any) {
	if interfacesAreEqual(a, b) {
		msg := fmt.Sprintf(
			ErrfNotDifferent,
			pretty.Sprint(a),
			pretty.Sprint(b),
			pretty.Diff(a, b),
		)
		test.t.Errorf(msg)
		test.t.FailNow()
	}
}
