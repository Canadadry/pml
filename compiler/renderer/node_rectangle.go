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
	n.children = append(n.children, child)
	return nil
}
func (*NodeRectangle) new(item *ast.Item) (Node, error) {
	var err error
	n := &NodeRectangle{}
	n.color, err = item.GetPropertyAsColorWithDefault("color", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return nil, err
	}
	err = n.Frame.initFrom(item)
	return n, err
}
func (n *NodeRectangle) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.SetFillColor(n.color)
	rb = rb.Cut(n.Frame)
	pdf.Rect(rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}
