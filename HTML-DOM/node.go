package htmldom

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_UnsupportedNodeType   = "unsupported node type"
	Err_NodeIsNotSet          = "node is not set"
	Err_DomNodeIsNotFound     = "DOM node is not found"
	Err_StartingPointIsNotSet = "starting point node is not set"
)

func HasNonEmptyValue(node *html.Node) bool {
	if node == nil {
		return false
	}

	return len(strings.TrimSpace(node.Data)) > 0
}

func GetCleanValue(node *html.Node) (cleanValue string) {
	if node == nil {
		return cleanValue
	}

	return strings.TrimSpace(node.Data)
}

func GetOuterHtml(node *html.Node) (outerHtml string, err error) {
	if node == nil {
		return outerHtml, errors.ErrorFromText(Err_NodeIsNotSet)
	}

	var bufInnerHtml bytes.Buffer
	w := io.Writer(&bufInnerHtml)
	err = html.Render(w, node)
	return bufInnerHtml.String(), nil
}

func GetInnerHtml(node *html.Node) (innerHtml string, err error) {
	if node == nil {
		return innerHtml, errors.ErrorFromText(Err_NodeIsNotSet)
	}

	switch node.Type {
	case html.TextNode:
		return node.Data, nil
	case html.ElementNode:
		return concatenateOuterHtmlOfNodes(GetChildNodes(node))
	default:
		return "", errors.ErrorFromText(Err_UnsupportedNodeType)
	}
}

func concatenateOuterHtmlOfNodes(nodes []*html.Node) (outerHtml string, err error) {
	var sb strings.Builder
	var buf string

	for _, node := range nodes {
		buf, err = GetOuterHtml(node)
		if err != nil {
			return "", err
		}

		_, err = sb.WriteString(buf)
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}
