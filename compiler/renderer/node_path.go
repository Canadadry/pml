package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
	"image/color"
)

type NodePath struct {
	Frame
	color       color.RGBA
	borderColor color.RGBA
	borderWidth float64
	path        string
}

func (n *NodePath) Children() []Node          { return nil }
func (n *NodePath) needToDrawChild() bool     { return false }
func (n *NodePath) addChild(child Node) error { return errChildrenNotAllowed }

func (*NodePath) new(item *ast.Item) (Node, error) {
	var err error
	n := &NodePath{}
	n.color, err = item.GetPropertyAsColorWithDefault("color", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return nil, err
	}
	n.borderColor, err = item.GetPropertyAsColorWithDefault("borderColor", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return nil, err
	}
	n.borderWidth, err = item.GetPropertyAsFloatWithDefault("borderWidth", 0)
	if err != nil {
		return nil, err
	}
	n.path, err = item.GetPropertyAsStringWithDefault("path", "")
	if err != nil {
		return nil, err
	}
	err = n.Frame.initFrom(item)
	return n, err
}
func (n *NodePath) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.SetFillColor(n.color)
	pdf.SetStrokeColor(n.borderColor)
	pdf.SetStrokeWidth(n.borderWidth)
	rb = rb.Cut(n.Frame)
	pdf.Path(n.path, rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}
