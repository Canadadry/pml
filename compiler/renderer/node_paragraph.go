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
	switch child.(type) {
	case *NodeDocument:
		return errChildrenNotAllowed
	case *NodePage:
		return errChildrenNotAllowed
	case *NodeRectangle:
		return errChildrenNotAllowed
	case *NodeText:
	case *NodeFont:
		return errChildrenNotAllowed
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
func (n *NodeParagraph) initFrom(item *ast.Item) error {
	var err error

	n.lineHeight, err = item.GetPropertyAsFloatWithDefault("lineHeight", 6)
	if err != nil {
		return err
	}
	xvalues := []string{Left, Center, Right, Fill, Layout, Free}
	n.xAlign, err = item.GetPropertyAsIdentifierFromListWithDefault("xAlign", xvalues[5], xvalues)
	if err != nil {
		return err
	}
	yvalues := []string{Top, Middle, Bottom, Fill, Layout, Free}
	n.yAlign, err = item.GetPropertyAsIdentifierFromListWithDefault("yAlign", yvalues[5], yvalues)
	if err != nil {
		return err
	}

	return n.Frame.initFrom(item)
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
			maxSize, textWidth := pdf.GetTextMaxLength(textChild.text[offset:], n.width-x)
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
