package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

type NodeDocument struct {
	children []Node
}

func (n *NodeDocument) Children() []Node { return n.children }
func (n *NodeDocument) addChild(child Node) error {
	switch child.(type) {
	case *NodeDocument:
		return errChildrenNotAllowed
	case *NodePage:
	case *NodeRectangle:
		return errChildrenNotAllowed
	case *NodeText:
		return errChildrenNotAllowed
	case *NodeFont:
	case *NodeImage:
		return errChildrenNotAllowed
	case *NodeVector:
		return errChildrenNotAllowed
	case *NodeParagraph:
		return errChildrenNotAllowed
	case *NodeContainer:
		return errChildrenNotAllowed
	}
	n.children = append(n.children, child)
	return nil
}
func (n *NodeDocument) initFrom(*ast.Item) error { return nil }
func (n *NodeDocument) needToDrawChild() bool    { return true }
func (n *NodeDocument) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	return rb, nil
}
