package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

type NodePage struct {
	children []Node
}

func (n *NodePage) Children() []Node      { return n.children }
func (n *NodePage) needToDrawChild() bool { return true }
func (n *NodePage) addChild(child Node) error {
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
func (n *NodePage) initFrom(*ast.Item) error { return nil }
func (n *NodePage) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.AddPage()
	return renderBox{w: PdfWidthMm, h: PdfHeight}, nil
}
