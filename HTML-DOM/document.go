package htmldom

import (
	"golang.org/x/net/html"
)

// GetElementById searches all descendant nodes for a first node with
// specified ID.
func GetElementById(root *html.Node, id string) *html.Node {
	// Loop over all descendants in depth-first order.
	for node := range root.Descendants() {
		if node.Type == html.ElementNode {
			for _, attr := range node.Attr {
				if attr.Key == AttributeId && attr.Val == id {
					return node
				}
			}
		}
	}
	return nil
}

// GetElementsByClassName returns all descendant nodes with specified class
// name.
func GetElementsByClassName(root *html.Node, className string) (out []*html.Node) {
	out = make([]*html.Node, 0)

	// Loop over all descendants in depth-first order.
	var cn string
	var ok bool
	for node := range root.Descendants() {
		if node.Type != html.ElementNode {
			continue
		}

		cn, ok = GetNodeAttributeValue(node, AttributeClass)
		if !ok {
			continue
		}

		if cn != className {
			continue
		}

		out = append(out, node)
	}

	return out
}

// GetElementsByName returns all descendant nodes with specified name.
func GetElementsByName(root *html.Node, name string) (out []*html.Node) {
	out = make([]*html.Node, 0)

	// Loop over all descendants in depth-first order.
	var n string
	var ok bool
	for node := range root.Descendants() {
		if node.Type != html.ElementNode {
			continue
		}

		n, ok = GetNodeAttributeValue(node, AttributeName)
		if !ok {
			continue
		}

		if n != name {
			continue
		}

		out = append(out, node)
	}

	return out
}

// GetElementsByTagName returns all descendant nodes with specified tag name.
func GetElementsByTagName(root *html.Node, tagName string) (out []*html.Node) {
	out = make([]*html.Node, 0)

	// Loop over all descendants in depth-first order.
	for node := range root.Descendants() {
		if node.Type != html.ElementNode {
			continue
		}

		if node.Data != tagName {
			continue
		}

		out = append(out, node)
	}

	return out
}
