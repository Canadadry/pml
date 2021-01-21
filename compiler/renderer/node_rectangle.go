package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
	"image/color"
)

type NodeRectangle struct {
	Frame
	children []Node
	color    color.RGBA
}

func (n *NodeRectangle) Children() []Node      { return n.children }
func (n *NodeRectangle) needToDrawChild() bool { return true }
func (n *NodeRectangle) addChild(child Node) error {
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
		return errChildrenNotAllowed
	}
	n.children = append(n.children, child)
	return nil
}
func (n *NodeRectangle) initFrom(item *ast.Item) error {
	var err error

	n.color, err = item.GetPropertyAsColorWithDefault("color", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return err
	}
	return n.Frame.initFrom(item)
}
func (n *NodeRectangle) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.SetFillColor(n.color)
	rb = rb.Cut(n.Frame)
	pdf.Rect(rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}
