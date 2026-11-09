package htmldom

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
	"golang.org/x/net/html"
)

func Test_GetElementById(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result *html.Node
	var err error
	var className string
	var ok bool

	// Test #1.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	result = GetElementById(pivot, "id-object-e")
	className, ok = GetNodeAttributeValue(result, AttributeClass)
	aTest.MustBeEqual(ok, true)
	aTest.MustBeEqual(className, "class-object-e")
}

func Test_GetElementsByClassName(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result []*html.Node
	var err error

	// Test #1.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	result = GetElementsByClassName(pivot, "sub-f")
	aTest.MustBeEqual(len(result), 2)

	// Check values.
	aTest.MustBeEqual(GetTextContent(result[0]), "SF-1")
	aTest.MustBeEqual(GetTextContent(result[1]), "SF-2")
}

func Test_GetElementsByName(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result []*html.Node
	var err error

	// Test #1.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	result = GetElementsByName(pivot, "universe")
	aTest.MustBeEqual(len(result), 2)

	// Check values.
	aTest.MustBeEqual(GetTextContent(result[0]), "G.1")
	aTest.MustBeEqual(GetTextContent(result[1]), "G.2")
}

func Test_GetElementsByTagName(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result []*html.Node
	var err error

	// Test #1.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetElementById(pivot, "id-object-h")
	result = GetElementsByTagName(pivot, TagSpan)
	aTest.MustBeEqual(len(result), 2)

	// Check values.
	aTest.MustBeEqual(GetNodeTagName(result[0]), TagSpan)
	aTest.MustBeEqual(GetNodeTagName(result[1]), TagSpan)
	aTest.MustBeEqual(GetTextContent(result[0]), "H-1")
	aTest.MustBeEqual(GetTextContent(result[1]), "H-2")
}
