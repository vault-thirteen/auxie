package httphelper

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewAcceptedMimeTypesFromHeader(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var amts *AcceptedMimeTypes
	var amt *AcceptedMimeType

	// Test #1.
	amts, err = NewAcceptedMimeTypesFromHeader("")
	aTest.MustBeAnError(err)

	// Test #2.
	amts, err = NewAcceptedMimeTypesFromHeader("text/html")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(len(amts.values), 1)

	// Test #3.
	amts, err = NewAcceptedMimeTypesFromHeader("*/*;q=0.8, text/html, application/xml;q=0.9, image/webp")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(len(amts.values), 4)
	//
	amt, err = amts.Next()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(amt.MimeType, "text/html")
	aTest.MustBeEqual(amt.Weight, float32(1.0))
	//
	amt, err = amts.Next()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(amt.MimeType, "image/webp")
	aTest.MustBeEqual(amt.Weight, float32(1.0))
	//
	amt, err = amts.Next()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(amt.MimeType, "application/xml")
	aTest.MustBeEqual(amt.Weight, float32(0.9))
	//
	amt, err = amts.Next()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(amt.MimeType, "*/*")
	aTest.MustBeEqual(amt.Weight, float32(0.8))

}

func Test_newAcceptedMimeTypes(t *testing.T) {
	aTest := tester.New(t)
	var amts *AcceptedMimeTypes
	var values []*AcceptedMimeType

	// Test #1.
	values = []*AcceptedMimeType{}
	amts = newAcceptedMimeTypes(values)
	aTest.MustBeEqual(amts.values, values)
	aTest.MustBeEqual(amts.lastIdx, -1)
	aTest.MustBeEqual(amts.cursor, -1)
}

func Test_resetIterator(t *testing.T) {
	aTest := tester.New(t)
	var amts *AcceptedMimeTypes

	// Test.
	amts = &AcceptedMimeTypes{
		cursor: 100,
	}
	amts.resetIterator()
	aTest.MustBeEqual(amts.cursor, -1)
}

func Test_parseAcceptHttpHeader(t *testing.T) {
	aTest := tester.New(t)
	var types []*AcceptedMimeType
	var err error

	// Test #1.
	types, err = parseAcceptHttpHeader("")
	aTest.MustBeAnError(err)

	// Test #2.
	types, err = parseAcceptHttpHeader(",")
	aTest.MustBeAnError(err)

	// Test #3.
	types, err = parseAcceptHttpHeader("text/html,")
	aTest.MustBeAnError(err)

	// Test #4.
	types, err = parseAcceptHttpHeader("text/html, image/webp")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(len(types), 2)
}

func Test_Next(t *testing.T) {
	aTest := tester.New(t)
	var amts *AcceptedMimeTypes
	var amt *AcceptedMimeType
	var err error

	// Test.
	amts, err = NewAcceptedMimeTypesFromHeader("text/html")
	aTest.MustBeNoError(err)
	//
	amt, err = amts.Next()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(amt.MimeType, "text/html")
	aTest.MustBeEqual(amt.Weight, float32(1.0))
	//
	amt, err = amts.Next()
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), ErrEndOfList)
}

func Test_Reset(t *testing.T) {
	aTest := tester.New(t)
	var amts *AcceptedMimeTypes

	// Test.
	amts = &AcceptedMimeTypes{
		cursor: 100,
	}
	amts.Reset()
	aTest.MustBeEqual(amts.cursor, -1)
}
