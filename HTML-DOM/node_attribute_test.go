package htmldom

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
	"golang.org/x/net/html"
)

func Test_NodeHasAttribute(t *testing.T) {

	var aTest = tester.New(t)
	var node *html.Node
	var result bool

	// Test #1. Null Node.
	node = nil
	result = NodeHasAttribute(node, "class")
	aTest.MustBeEqual(result, false)

	// Test #2. Negative.
	node = &html.Node{
		Attr: []html.Attribute{
			{
				Key: "John",
				Val: "Doe",
			},
		},
	}
	result = NodeHasAttribute(node, "class")
	aTest.MustBeEqual(result, false)

	// Test #3. Positive.
	node = &html.Node{
		Attr: []html.Attribute{
			{
				Key: "John",
				Val: "Doe",
			},
			{
				Key: "class",
				Val: "the-class",
			},
		},
	}
	result = NodeHasAttribute(node, "class")
	aTest.MustBeEqual(result, true)
}

func Test_GetNodeAttributeValue(t *testing.T) {

	var aTest = tester.New(t)
	var node *html.Node
	var attributeValue string
	var attributeExists bool

	// Test #1. Null Node.
	node = nil
	attributeValue, attributeExists = GetNodeAttributeValue(node, "class")
	aTest.MustBeEqual(attributeValue, "")
	aTest.MustBeEqual(attributeExists, false)

	// Test #2. Negative.
	node = &html.Node{
		Attr: []html.Attribute{
			{
				Key: "John",
				Val: "Doe",
			},
		},
	}
	attributeValue, attributeExists = GetNodeAttributeValue(node, "class")
	aTest.MustBeEqual(attributeValue, "")
	aTest.MustBeEqual(attributeExists, false)

	// Test #3. Positive.
	node = &html.Node{
		Attr: []html.Attribute{
			{
				Key: "John",
				Val: "Doe",
			},
			{
				Key: "class",
				Val: "the-class",
			},
		},
	}
	attributeValue, attributeExists = GetNodeAttributeValue(node, "class")
	aTest.MustBeEqual(attributeValue, "the-class")
	aTest.MustBeEqual(attributeExists, true)
}
