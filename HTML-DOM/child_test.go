package htmldom

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
	"golang.org/x/net/html"
)

func Test_GetChildNodeByTag(t *testing.T) {

	// Preparation.
	var aTest = tester.New(t)
	var err error
	var containerNode *html.Node
	containerNode, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	var divA *html.Node
	divA = GetChildNodeByTag(containerNode, TagDiv)
	aTest.MustBeDifferent(divA, (*html.Node)(nil))
	var divA1 *html.Node
	divA1 = GetChildNodeByTag(divA, TagDiv)
	aTest.MustBeDifferent(divA1, (*html.Node)(nil))
	var result *html.Node

	// Test #1. Fool Check.
	result = GetChildNodeByTag(nil, TagDiv)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Child Match.
	result = GetChildNodeByTag(containerNode, TagDiv)
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a",
		},
		{
			Key: AttributeId,
			Val: "id-object-a",
		},
	})

	// Test #3. No Matches.
	result = GetChildNodeByTag(containerNode, TagB)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Child.
	result = GetChildNodeByTag(divA1, TagDiv)
	aTest.MustBeEqual(result, (*html.Node)(nil))
}

func Test_GetChildNodeByTagAndClass(t *testing.T) {

	// Preparation.
	var aTest = tester.New(t)
	var containerNode *html.Node
	var err error
	containerNode, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	var divA *html.Node
	divA = GetChildNodeByTagAndClass(containerNode, TagDiv, "class-object-a")
	aTest.MustBeDifferent(divA, (*html.Node)(nil))
	aTest.MustBeEqual(divA.Data, TagDiv)
	aTest.MustBeEqual(divA.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a",
		},
		{
			Key: AttributeId,
			Val: "id-object-a",
		},
	})
	var divA1 *html.Node
	divA1 = GetChildNodeByTagAndClass(divA, TagDiv, "class-object-a-1")
	aTest.MustBeDifferent(divA1, (*html.Node)(nil))
	var result *html.Node

	// Test #1. Fool Check.
	result = GetChildNodeByTagAndClass(nil, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Child Match.
	result = GetChildNodeByTagAndClass(divA, TagDiv, "class-object-a-3")
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-3",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-3",
		},
	})

	// Test #3. No Matches.
	result = GetChildNodeByTagAndClass(containerNode, TagB, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Child.
	result = GetChildNodeByTagAndClass(divA1, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))
}

func Test_GetChildNodeByTagAndId(t *testing.T) {

	// Preparation.
	var aTest = tester.New(t)
	var containerNode *html.Node
	var err error
	containerNode, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	var divA *html.Node
	divA = GetChildNodeByTagAndId(containerNode, TagDiv, "id-object-a")
	aTest.MustBeDifferent(divA, (*html.Node)(nil))
	aTest.MustBeEqual(divA.Data, TagDiv)
	aTest.MustBeEqual(divA.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a",
		},
		{
			Key: AttributeId,
			Val: "id-object-a",
		},
	})
	var divA1 *html.Node
	divA1 = GetChildNodeByTagAndId(divA, TagDiv, "id-object-a-1")
	aTest.MustBeDifferent(divA1, (*html.Node)(nil))
	var result *html.Node

	// Test #1. Fool Check.
	result = GetChildNodeByTagAndId(nil, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Child Match.
	result = GetChildNodeByTagAndId(divA, TagDiv, "id-object-a-3")
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-3",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-3",
		},
	})

	// Test #3. No Matches.
	result = GetChildNodeByTagAndId(containerNode, TagB, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Child.
	result = GetChildNodeByTagAndId(divA1, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))
}

func Test_GetChildNodeByType(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result *html.Node
	var err error

	// Test #1. Fool Check.
	result = GetChildNodeByType(nil, html.TextNode)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. Nothing is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a")
	//
	result = GetChildNodeByType(pivot, html.CommentNode)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #3. Comment is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	//
	result = GetChildNodeByType(pivot, html.CommentNode)
	aTest.MustBeEqual(result.Data, " comment #1 ")

	// Test #4. Comment is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-c")
	//
	result = GetChildNodeByType(pivot, html.CommentNode)
	aTest.MustBeEqual(result.Data, " comment #2 ")

	// Test #5. Text is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	//
	result = GetChildNodeByType(pivot, html.TextNode)
	aTest.MustBeEqual(result.Data, "Text A")

	// Test #6. Text is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-c")
	//
	result = GetChildNodeByType(pivot, html.TextNode)
	aTest.MustBeEqual(result.Data, "Text B")

	// Test #6. Element is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	//
	result = GetChildNodeByType(pivot, html.ElementNode)
	aTest.MustBeEqual(result.Data, "b")
	aTest.MustBeEqual(result.Attr[0], html.Attribute{Key: "id", Val: "ob-b1"})

	// Test #7. Element is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-c")
	//
	result = GetChildNodeByType(pivot, html.ElementNode)
	aTest.MustBeEqual(result.Data, "b")
	aTest.MustBeEqual(result.Attr[0], html.Attribute{Key: "id", Val: "oc-b2"})
}

func Test_GetChildComment(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result *html.Node
	var err error

	// Test #1. Fool Check.
	result = GetChildComment(nil)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. Nothing is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a")
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a-1")
	//
	result = GetChildComment(pivot)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #3. Child is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	//
	result = GetChildComment(pivot)
	aTest.MustBeEqual(result.Data, " comment #1 ")
}

func Test_GetChildElement(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result *html.Node
	var err error

	// Test #1. Fool Check.
	result = GetChildElement(nil)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. Nothing is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a")
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a-1")
	//
	result = GetChildElement(pivot)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #3. Child is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	//
	result = GetChildElement(pivot)
	aTest.MustBeEqual(result.Data, "b")
	aTest.MustBeEqual(result.Attr[0], html.Attribute{Key: "id", Val: "ob-b1"})

	// Test #4. Child is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagSpan, "id-object-d")
	//
	result = GetChildElement(pivot)
	aTest.MustBeEqual(result.Data, "a")
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{Key: "class", Val: "class-object-d-1"},
		{Key: "id", Val: "id-object-d-1"},
	})
}

func Test_GetChildValue(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result *html.Node
	var err error

	// Test #1. Fool Check.
	result = GetChildValue(nil)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. Nothing is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a")
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a-1")
	//
	result = GetChildValue(pivot)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #3. Child is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	//
	result = GetChildValue(pivot)
	aTest.MustBeEqual(result.Data, "Text A")
}

func Test_GetChildValueNE(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result *html.Node
	var err error

	// Test.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a")
	//
	result = GetChildValueNE(pivot)
	aTest.MustBeEqual(GetCleanValue(result), "NE Value")
}

func Test_GetChildNodes(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result []*html.Node
	var err error

	// Test #1. Fool Check.
	result = GetChildNodes(nil)
	aTest.MustBeEqual(result, ([]*html.Node)(nil))

	// Test #2. Nothing is found.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-a")
	//
	result = GetChildNodes(pivot)
	aTest.MustBeEqual(len(result), 4+3)
	aTest.MustBeEqual(result[1].Data, "div")
	aTest.MustBeEqual(result[3].Data, "div")
	aTest.MustBeEqual(GetCleanValue(result[4]), "NE Value")
	aTest.MustBeEqual(result[5].Data, "div")
}
