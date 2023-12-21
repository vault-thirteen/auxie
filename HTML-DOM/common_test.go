package htmldom

import (
	"errors"
	"io"
	"os"
	"strings"

	"go.uber.org/multierr"
	"golang.org/x/net/html"
)

// Auxiliary functions for tests.

const SampleSourceFileName = "test/sample.html"

func _test_parseHtmlSourceIntoDom(htmlPageSource string) (domNode *html.Node, err error) {
	domNode, err = html.Parse(strings.NewReader(htmlPageSource))
	if err != nil {
		return nil, err
	}

	if domNode == nil {
		return nil, errors.New(ErrDomNodeIsNotFound)
	}

	return domNode, nil
}

func _test_getSampleHtmlSource() (source string, err error) {
	var file *os.File
	file, err = os.Open(SampleSourceFileName)
	if err != nil {
		return "", err
	}

	defer func() {
		derr := file.Close()
		err = multierr.Combine(err, derr)
	}()

	var fileContents []byte
	fileContents, err = io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(fileContents), nil
}

func _test_createSimpleDomForTests() (domNode *html.Node, err error) {
	var source string
	source, err = _test_getSampleHtmlSource()
	if err != nil {
		return nil, err
	}

	return _test_parseHtmlSourceIntoDom(source)
}

func _test_getContainerNodeFromSimpleDomForTests(domNode *html.Node) (containerNode *html.Node, err error) {
	// Fool Check.
	if domNode == nil {
		return nil, errors.New(ErrStartingPointIsNotSet)
	}

	// Search.
	var node = domNode.FirstChild // -> <html>
	if node == nil {
		return nil, errors.New(ErrDomNodeIsNotFound)
	}

	node = node.FirstChild // -> <head>
	if node == nil {
		return nil, errors.New(ErrDomNodeIsNotFound)
	}

	node = node.NextSibling.NextSibling // -> <body>
	if node == nil {
		return nil, errors.New(ErrDomNodeIsNotFound)
	}

	node = node.FirstChild.NextSibling // -> <div class="class-container">
	if node == nil {
		return nil, errors.New(ErrDomNodeIsNotFound)
	}

	return node, nil
}

func _test_prepareTestNode() (containerNode *html.Node, err error) {
	var domNode *html.Node
	domNode, err = _test_createSimpleDomForTests()
	if err != nil {
		return nil, err
	}

	return _test_getContainerNodeFromSimpleDomForTests(domNode)
}
