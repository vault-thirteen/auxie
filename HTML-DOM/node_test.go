package htmldom

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
	"golang.org/x/net/html"
)

func Test_HasNonEmptyValue(t *testing.T) {
	var aTest = tester.New(t)
	var node *html.Node

	// Test #1.
	aTest.MustBeEqual(HasNonEmptyValue(nil), false)

	// Test #2.
	node = &html.Node{
		Data: "  \n  \t  \r\n  ",
	}
	aTest.MustBeEqual(HasNonEmptyValue(node), false)

	// Test #3.
	node = &html.Node{
		Data: "  \n  \t  Voxel  \r\n  ",
	}
	aTest.MustBeEqual(HasNonEmptyValue(node), true)
}

func Test_GetCleanValue(t *testing.T) {
	var aTest = tester.New(t)
	var node *html.Node

	// Test #1.
	aTest.MustBeEqual(GetCleanValue(nil), "")

	// Test #2.
	node = &html.Node{
		Data: "  \n  \t  Срамота   \r\n  ",
	}
	aTest.MustBeEqual(GetCleanValue(node), "Срамота")
}

func Test_GetOuterHtml(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result string
	var err error

	// Test #1.
	result, err = GetOuterHtml(nil)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, "")

	// Test #2. Comment.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	pivot = GetChildComment(pivot)
	//
	result, err = GetOuterHtml(pivot)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, "<!-- comment #1 -->")

	// Test #3. Text Value.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	pivot = GetChildValue(pivot)
	//
	result, err = GetOuterHtml(pivot)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, "Text A")

	// Test #4. Simple Element.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	pivot = GetChildNodeByTag(pivot, TagB)
	//
	result, err = GetOuterHtml(pivot)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, `<b id="ob-b1">Element 1</b>`)

	// Test #5. Composite Element.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-c")
	//
	result, err = GetOuterHtml(pivot)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, `<div class="class-object-c" id="id-object-c">Text B<!-- comment #2 --><b id="oc-b2">Element 2</b>
                <div class="class-object-c-1" id="id-object-c-1">C-1</div>
                <div class="class-object-c-2" id="id-object-c-2">C-2</div>Ѧ<p>John Carmack</p>Ѫ
                <div class="class-object-c-3" id="id-object-c-3">C-3</div>Ѫ<p>Voxel</p>Ѧ
            </div>`)
}

func Test_GetInnerHtml(t *testing.T) {
	var aTest = tester.New(t)
	var pivot *html.Node
	var result string
	var err error

	// Test #1.
	result, err = GetInnerHtml(nil)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, "")

	// Test #2. Comment.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	pivot = GetChildComment(pivot)
	//
	result, err = GetInnerHtml(pivot)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, "")

	// Test #3. Text Value.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-b")
	pivot = GetChildValue(pivot)
	//
	result, err = GetInnerHtml(pivot)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, "Text A")

	// Test #4. Element.
	pivot, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	pivot = GetChildNodeByTagAndId(pivot, TagDiv, "id-object-c")
	//
	result, err = GetInnerHtml(pivot)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, `Text B<!-- comment #2 --><b id="oc-b2">Element 2</b>
                <div class="class-object-c-1" id="id-object-c-1">C-1</div>
                <div class="class-object-c-2" id="id-object-c-2">C-2</div>Ѧ<p>John Carmack</p>Ѫ
                <div class="class-object-c-3" id="id-object-c-3">C-3</div>Ѫ<p>Voxel</p>Ѧ
            `)
}
