package httphelper

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ParseRecord(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var amt *AcceptedMimeType

	// Test #1.
	amt, err = ParseRecord("")
	aTest.MustBeAnError(err)

	// Test #2.
	amt, err = ParseRecord("*/*")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(amt.MimeType, "*/*")
	aTest.MustBeEqual(amt.Weight, float32(1.0))

	// Test #3.
	amt, err = ParseRecord("*/x")
	aTest.MustBeAnError(err)

	// Test #4.
	amt, err = ParseRecord("application/xml;q=0.9")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(amt.MimeType, "application/xml")
	aTest.MustBeEqual(amt.Weight, float32(0.9))

	// Test #5.
	amt, err = ParseRecord("application/xml;q=0.x")
	aTest.MustBeAnError(err)

	// Test #6.
	amt, err = ParseRecord("application/xml;q=0.5;junk")
	aTest.MustBeAnError(err)
}

func Test_ParseRecordWeight(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var w float32

	// Test #1.
	w, err = ParseRecordWeight("")
	aTest.MustBeAnError(err)

	// Test #2.
	w, err = ParseRecordWeight("=")
	aTest.MustBeAnError(err)

	// Test #3.
	w, err = ParseRecordWeight("=x")
	aTest.MustBeAnError(err)

	// Test #4.
	w, err = ParseRecordWeight("x=")
	aTest.MustBeAnError(err)

	// Test #5.
	w, err = ParseRecordWeight("x=x")
	aTest.MustBeAnError(err)

	// Test #6.
	w, err = ParseRecordWeight("q=")
	aTest.MustBeAnError(err)

	// Test #7.
	w, err = ParseRecordWeight("q=x")
	aTest.MustBeAnError(err)

	// Test #8.
	w, err = ParseRecordWeight("q=0.5")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(w, float32(0.5))
}

func Test_checkMimeType(t *testing.T) {
	aTest := tester.New(t)
	var err error

	// Test #1.
	err = checkMimeType("")
	aTest.MustBeAnError(err)

	// Test #2.
	err = checkMimeType("*/*")
	aTest.MustBeNoError(err)
}
