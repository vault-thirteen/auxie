package mime

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_GetCategories(t *testing.T) {
	aTest := tester.New(t)

	catsExpected := []string{
		"*",
		"application",
		"audio",
		"font",
		"image",
		"model",
		"multipart",
		"text",
		"video",
	}
	c := GetCategories()
	aTest.MustBeEqual(c, catsExpected)
}

func Test_IsCategoryValid(t *testing.T) {
	aTest := tester.New(t)

	// Test #1.
	aTest.MustBeEqual(IsCategoryValid("application"), true)

	// Test #2.
	aTest.MustBeEqual(IsCategoryValid("*"), true)

	// Test #3.
	aTest.MustBeEqual(IsCategoryValid("zeliboba"), false)
}

func Test_GetMimeTypeParts(t *testing.T) {
	aTest := tester.New(t)
	var (
		category string
		subtype  string
		err      error
	)

	// Test #1.
	category, subtype, err = GetMimeTypeParts("*")
	aTest.MustBeAnError(err)

	// Test #2.
	category, subtype, err = GetMimeTypeParts("*/*")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(category, "*")
	aTest.MustBeEqual(subtype, "*")

	// Test #3.
	category, subtype, err = GetMimeTypeParts("text/*")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(category, "text")
	aTest.MustBeEqual(subtype, "*")

	// Test #4.
	category, subtype, err = GetMimeTypeParts("*/earth")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(category, "")
	aTest.MustBeEqual(subtype, "")

	// Test #5.
	category, subtype, err = GetMimeTypeParts("planet/*")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(category, "")
	aTest.MustBeEqual(subtype, "")

	// Test #6.
	category, subtype, err = GetMimeTypeParts("text/plain")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(category, "text")
	aTest.MustBeEqual(subtype, "plain")
}
