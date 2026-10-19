package htmldom

import "golang.org/x/net/html"

// GetSiblingNodeByTag searches for the nearest sibling node having the
// specified tag name.
func GetSiblingNodeByTag(startingPoint *html.Node, tagName string) (siblingNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	siblingNode = startingPoint.NextSibling
	if siblingNode == nil {
		return nil
	}

	for {
		if siblingNode.Data == tagName {
			return siblingNode
		} else {
			siblingNode = siblingNode.NextSibling

			if siblingNode == nil {
				return nil
			}
		}
	}
}

// GetSiblingNodeByTagAndClass searches for the nearest sibling node having the
// specified tag name and class.
func GetSiblingNodeByTagAndClass(startingPoint *html.Node, tagName string, className string) (siblingNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	siblingNode = startingPoint.NextSibling
	if siblingNode == nil {
		return nil
	}

	for {
		var classValue string
		var classExists bool
		classValue, classExists = GetNodeAttributeValue(siblingNode, AttributeClass)
		if (siblingNode.Data == tagName) &&
			(classExists) &&
			(classValue == className) {
			return siblingNode
		} else {
			siblingNode = siblingNode.NextSibling

			if siblingNode == nil {
				return nil
			}
		}
	}
}

// GetSiblingNodeByTagAndId searches for the nearest sibling node having the
// specified tag name and ID.
func GetSiblingNodeByTagAndId(startingPoint *html.Node, tagName string, idName string) (siblingNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	siblingNode = startingPoint.NextSibling
	if siblingNode == nil {
		return nil
	}

	for {
		var idValue string
		var idExists bool
		idValue, idExists = GetNodeAttributeValue(siblingNode, AttributeId)
		if (siblingNode.Data == tagName) &&
			(idExists) &&
			(idValue == idName) {
			return siblingNode
		} else {
			siblingNode = siblingNode.NextSibling

			if siblingNode == nil {
				return nil
			}
		}
	}
}

// GetSiblingNodeByType searches for the nearest sibling having the specified node type.
func GetSiblingNodeByType(startingPoint *html.Node, nodeType html.NodeType) (siblingNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	siblingNode = startingPoint.NextSibling
	if siblingNode == nil {
		return nil
	}

	for {
		if siblingNode.Type == nodeType {
			return siblingNode
		} else {
			siblingNode = siblingNode.NextSibling

			if siblingNode == nil {
				return nil
			}
		}
	}
}

// GetSiblingComment searches for the nearest sibling comment.
func GetSiblingComment(startingPoint *html.Node) (siblingNode *html.Node) {
	return GetSiblingNodeByType(startingPoint, html.CommentNode)
}

// GetSiblingElement searches for the nearest sibling element.
func GetSiblingElement(startingPoint *html.Node) (siblingNode *html.Node) {
	return GetSiblingNodeByType(startingPoint, html.ElementNode)
}

// GetSiblingValue searches for the nearest sibling [text] value.
func GetSiblingValue(startingPoint *html.Node) (siblingNode *html.Node) {
	return GetSiblingNodeByType(startingPoint, html.TextNode)
}

// GetSiblingValueNE searches for the nearest sibling [text] value which is not
// empty.
func GetSiblingValueNE(startingPoint *html.Node) (siblingNode *html.Node) {
	node := GetSiblingNodeByType(startingPoint, html.TextNode)
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
