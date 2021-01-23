package renderer

import (
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
)

type NodeParagraph struct {
	Frame
	children   []Node
	lineHeight float64
}

func (n *NodeParagraph) Children() []Node      { return n.children }
func (n *NodeParagraph) needToDrawChild() bool { return false }
func (n *NodeParagraph) addChild(child Node) error {
	n.children = append(n.children, child)
	return nil
}
func (*NodeParagraph) new(item *ast.Item) (Node, error) {
	n := &NodeParagraph{}
	var err error

	n.lineHeight, err = item.GetPropertyAsFloatWithDefault("lineHeight", 6)
	if err != nil {
		return nil, err
	}

	err = n.Frame.initFrom(item)
	return n, err
}
func (n *NodeParagraph) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	x := 0.0
	y := 0.0
	rb = rb.Cut(n.Frame)

	for _, child := range n.children {
		offset := 0
		textChild, ok := child.(*NodeText)
		if !ok {
			return rb, fmt.Errorf("Unexpected node in paragraph")
		}
		if len(textChild.fontName) == 0 {
			textChild.fontName = pdf.GetDefaultFontName()
		}

		pdf.SetFont(textChild.fontName, textChild.fontSize)
		pdf.SetTextColor(textChild.color)

		for offset < len(textChild.text) {
			maxSize, textWidth := pdf.GetTextMaxLength(textChild.text[offset:], rb.w-x)
			if textWidth == 0 {
				if x > 0 {
					x = 0
					y = y + n.lineHeight
					continue
				} else {
					break
				}
			}
			text := textChild.text[offset : offset+maxSize]
			pdf.Text(text, x+rb.x, y+rb.y, rb.w, n.lineHeight, "BaselineLeft")
			offset = offset + maxSize
			x = x + textWidth
			if x > rb.w {
				x = 0
				y = y + n.lineHeight
			}
		}
	}
	return rb, nil
}
