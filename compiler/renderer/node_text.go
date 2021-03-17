package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
	"image/color"
)

type NodeText struct {
	Frame
	text     string
	color    color.RGBA
	align    string
	fontName string
	fontSize float64
}

func (n *NodeText) Children() []Node          { return nil }
func (n *NodeText) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeText) needToDrawChild() bool     { return true }
func (*NodeText) new(item *ast.Item) (Node, error) {
	n := &NodeText{}
	var err error
	n.text, err = item.GetPropertyAsStringWithDefault("text", "")
	if err != nil {
		return nil, err
	}
	n.color, err = item.GetPropertyAsColorWithDefault("color", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return nil, err
	}
	values := []string{
		string(AlingTopLeft),
		string(AlingTopCenter),
		string(AlingTopRight),
		string(AlingMiddleLeft),
		string(AlingMiddleCenter),
		string(AlingMiddleRight),
		string(AlingBottomLeft),
		string(AlingBottomCenter),
		string(AlingBottomRight),
		string(AlingBaselineLeft),
		string(AlingBaselineCenter),
		string(AlingBaselineRight),
	}
	n.align, err = item.GetPropertyAsIdentifierFromListWithDefault("align", values[0], values)
	if err != nil {
		return nil, err
	}
	n.fontName, err = item.GetPropertyAsStringWithDefault("fontName", "Arial")
	if err != nil {
		return nil, err
	}
	n.fontSize, err = item.GetPropertyAsFloatWithDefault("fontSize", 6)
	if err != nil {
		return nil, err
	}
	err = n.Frame.initFrom(item)
	return n, err
}
func (n *NodeText) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {

	if len(n.fontName) == 0 {
		n.fontName = pdf.GetDefaultFontName()
	}
	pdf.SetFont(n.fontName, n.fontSize*rb.s*n.scale)
	pdf.SetTextColor(n.color)
	rb = rb.Cut(n.Frame)
	pdf.Text(n.text, rb.x, rb.y, rb.w, rb.h, PdfTextAlign(n.align))
	return rb, nil
}
