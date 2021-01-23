package renderer

import (
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"image/color"
	"strings"
)

const (
	ParagraphLeft    = "left"
	ParagraphRight   = "right"
	ParagraphCenter  = "center"
	ParagraphJustify = "justify"
)

type getAlginStartAndSpacing func(l Line) (float64, float64)

var getAlginStartAndSpacingByParagraphMode = map[string]getAlginStartAndSpacing{
	ParagraphLeft: func(l Line) (float64, float64) {
		return 0, l.spaceWidth
	},
	ParagraphRight: func(l Line) (float64, float64) {
		return l.spaceLeft - l.spaceWidth*float64(len(l.words)), l.spaceWidth
	},
	ParagraphCenter: func(l Line) (float64, float64) {
		return (l.spaceLeft - l.spaceWidth*float64(len(l.words))) / 2, l.spaceWidth
	},
	ParagraphJustify: func(l Line) (float64, float64) {
		return 0, l.spaceLeft / float64(len(l.words))
	},
}

type NodeParagraph struct {
	Frame
	children   []Node
	lineHeight float64
	align      getAlginStartAndSpacing
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
	values := []string{ParagraphLeft, ParagraphRight, ParagraphCenter, ParagraphJustify}
	align, err := item.GetPropertyAsIdentifierFromListWithDefault("align", values[3], values)
	if err != nil {
		return nil, err
	}
	var ok bool
	n.align, ok = getAlginStartAndSpacingByParagraphMode[align]
	if !ok {
		return nil, fmt.Errorf("Paragraph mode don't implement this mode %s", align)
	}
	err = n.Frame.initFrom(item)
	return n, err
}

type Word struct {
	text       string
	width      float64
	fontSize   float64
	fontName   string
	color      color.RGBA
	spaceWidth float64
}

func textToWords(pdf PdfDrawer, node NodeText) []Word {
	words := []Word{}

	pdf.SetFont(node.fontName, node.fontSize)
	_, spaceWidth := pdf.GetTextMaxLength(" ", PdfWidthMm)

	splitted := strings.Split(node.text, " ")
	for _, part := range splitted {
		if len(part) == 0 {
			continue
		}
		_, width := pdf.GetTextMaxLength(part, PdfWidthMm)
		words = append(words, Word{
			text:       part,
			width:      width,
			fontSize:   node.fontSize,
			fontName:   node.fontName,
			color:      node.color,
			spaceWidth: spaceWidth,
		})
	}
	return words
}

type Line struct {
	words      []Word
	spaceLeft  float64
	spaceWidth float64
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
			line.spaceLeft = width - x
			line.spaceWidth = line.spaceWidth / float64(len(line.words))
			lines = append(lines, line)
			line = Line{}
			x = 0
		}
		x = x + w.width + w.spaceWidth
		line.words = append(line.words, w)
		line.spaceWidth = line.spaceWidth + w.spaceWidth
	}
	if len(line.words) > 0 {
		line.spaceLeft = width - x
		line.spaceWidth = line.spaceWidth / float64(len(line.words))
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

	y := 0.0
	for _, l := range lines {
		x, spacing := n.align(l)
		for _, w := range l.words {
			pdf.SetFont(w.fontName, w.fontSize)
			pdf.SetTextColor(w.color)
			pdf.Text(w.text, x+rb.x, y+rb.y, rb.w, n.lineHeight, AlingBaselineLeft)
			x = x + w.width + spacing
		}
		y = y + n.lineHeight
	}
	return rb, nil
}
