package htmldom

import "golang.org/x/net/html"

// GetChildNodeByTag searches for the nearest child node having the specified
// tag name.
func GetChildNodeByTag(startingPoint *html.Node, tagName string) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}

	for {
		if childNode.Data == tagName {
			return childNode
		} else {
			childNode = childNode.NextSibling

			if childNode == nil {
				return nil
			}
		}
	}
}

// GetChildNodeByTagAndClass searches for the nearest child node having the
// specified tag name and class.
func GetChildNodeByTagAndClass(startingPoint *html.Node, tagName string, className string) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}

	for {
		var classValue string
		var classExists bool
		classValue, classExists = GetNodeAttributeValue(childNode, AttributeClass)

		if (childNode.Data == tagName) &&
			(classExists) &&
			(classValue == className) {
			return childNode
		} else {
			childNode = childNode.NextSibling

			if childNode == nil {
				return nil
			}
		}
	}
}

// GetChildNodeByTagAndId searches for the nearest child node having the
// specified tag name and id.
func GetChildNodeByTagAndId(startingPoint *html.Node, tagName string, idName string) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}

	for {
		var idValue string
		var idExists bool
		idValue, idExists = GetNodeAttributeValue(childNode, AttributeId)

		if (childNode.Data == tagName) &&
			(idExists) &&
			(idValue == idName) {
			return childNode
		} else {
			childNode = childNode.NextSibling

			if childNode == nil {
				return nil
			}
		}
	}
}

// GetChildNodeByType searches for the nearest child having the specified node type.
func GetChildNodeByType(startingPoint *html.Node, nodeType html.NodeType) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}

	for {
		if childNode.Type == nodeType {
			return childNode
		} else {
			childNode = childNode.NextSibling

			if childNode == nil {
				return nil
			}
		}
	}
}

// GetChildComment searches for the nearest child comment.
func GetChildComment(startingPoint *html.Node) (childNode *html.Node) {
	return GetChildNodeByType(startingPoint, html.CommentNode)
}

// GetChildElement searches for the nearest child element.
func GetChildElement(startingPoint *html.Node) (childNode *html.Node) {
	return GetChildNodeByType(startingPoint, html.ElementNode)
}

// GetChildValue searches for the nearest child [text] value.
func GetChildValue(startingPoint *html.Node) (childNode *html.Node) {
	return GetChildNodeByType(startingPoint, html.TextNode)
}

// GetChildValueNE searches for the nearest child [text] value which is not
// empty.
func GetChildValueNE(startingPoint *html.Node) (childNode *html.Node) {
	node := GetChildNodeByType(startingPoint, html.TextNode)
	if HasNonEmptyValue(node) {
		return node
	}

	for {
		node = GetSiblingNodeByType(node, html.TextNode)
		if node == nil {
			return nil
		}
		if HasNonEmptyValue(node) {
			return node
		}
	}
}

// GetChildNodes returns all child nodes.
func GetChildNodes(startingPoint *html.Node) (childNodes []*html.Node) {
	if startingPoint == nil {
		return nil
	}

	childNodes = make([]*html.Node, 0)
	node := startingPoint.FirstChild

	for node != nil {
		childNodes = append(childNodes, node)
		node = node.NextSibling
	}

	return childNodes
}
