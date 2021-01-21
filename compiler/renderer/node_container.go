package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

type NodeContainer struct {
	Frame
	children []Node
}

func (n *NodeContainer) Children() []Node      { return n.children }
func (n *NodeContainer) needToDrawChild() bool { return true }
func (n *NodeContainer) addChild(child Node) error {
	switch child.(type) {
	case *NodeDocument:
		return errChildrenNotAllowed
	case *NodePage:
		return errChildrenNotAllowed
	case *NodeRectangle:
	case *NodeText:
	case *NodeFont:
		return errChildrenNotAllowed
	case *NodeImage:
	case *NodeVector:
	case *NodeParagraph:
	case *NodeContainer:
	}
	n.children = append(n.children, child)
	return nil
}
func (n *NodeContainer) initFrom(item *ast.Item) error {
	return n.Frame.initFrom(item)
}

func (n *NodeContainer) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	return rb.Cut(n.Frame), nil
}
