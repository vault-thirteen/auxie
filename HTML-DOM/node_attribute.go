package htmldom

import (
	"golang.org/x/net/html"
)

// NodeHasAttribute checks whether the specified node has an attribute with the
// specified name.
func NodeHasAttribute(node *html.Node, attributeName string) (attributeExists bool) {
	if node == nil {
		return false
	}

	for _, attribute := range node.Attr {
		if attribute.Key == attributeName {
			return true
		}
	}

	return false
}

// GetNodeAttributeValue gets the value of an attribute specified by its name.
func GetNodeAttributeValue(node *html.Node, attributeName string) (attributeValue string, attributeExists bool) {
	if node == nil {
		return "", false
	}

	for _, attribute := range node.Attr {
		if attribute.Key == attributeName {
			return attribute.Val, true
		}
	}

	return "", false
}
