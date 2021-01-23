package renderer

import (
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"image/color"
	"strings"
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

type Word struct {
	text     string
	width    float64
	fontSize float64
	fontName string
	color    color.RGBA
}

func textToWords(pdf PdfDrawer, node NodeText) []Word {
	words := []Word{}

	pdf.SetFont(node.fontName, node.fontSize)

	splitted := strings.Split(node.text, " ")
	for _, part := range splitted {
		if len(part) == 0 {
			continue
		}
		_, width := pdf.GetTextMaxLength(part, PdfWidthMm)
		words = append(words, Word{
			text:     part,
			width:    width,
			fontSize: node.fontSize,
			fontName: node.fontName,
			color:    node.color,
		})
	}
	return words
}

type Line struct {
	words   []Word
	spacing float64
}

func wordsToLines(words []Word, width float64) []Line {
	x := 0.0
	lines := []Line{}
	line := Line{}
	for _, w := range words {
		if (x + w.width) > width {
			if x == 0 {
				continue
			}
			line.spacing = width - x
			lines = append(lines, line)
			line = Line{}
			x = 0
		}
		x = x + w.width
		line.words = append(line.words, w)
	}
	if len(line.words) > 0 {
		line.spacing = width - x
		lines = append(lines, line)
		line = Line{}
	}
	return lines
}

func (n *NodeParagraph) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	rb = rb.Cut(n.Frame)

	words := []Word{}

	for _, child := range n.children {
		textChild, ok := child.(*NodeText)
		if !ok {
			return rb, fmt.Errorf("Unexpected node in paragraph")
		}

		words = append(words, textToWords(pdf, *textChild)...)
	}

	lines := wordsToLines(words, rb.w)

	x := 0.0
	y := 0.0
	for _, l := range lines {
		spacing := l.spacing / float64(len(l.words)-1)
		for _, w := range l.words {
			pdf.SetFont(w.fontName, w.fontSize)
			pdf.SetTextColor(w.color)
			pdf.Text(w.text, x+rb.x, y+rb.y, rb.w, n.lineHeight, "BaselineLeft")
			x = x + w.width + spacing
		}
		x = 0
		y = y + n.lineHeight
	}
	return rb, nil
}
