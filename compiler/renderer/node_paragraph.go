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

const utf8LineBreak = "\xE2\x80\xA9"

type getAlginStartAndSpacing func(ls LineSize) (float64, float64)

var getAlginStartAndSpacingByParagraphMode = map[string]getAlginStartAndSpacing{
	ParagraphLeft: func(ls LineSize) (float64, float64) {
		return 0, 0
	},
	ParagraphRight: func(ls LineSize) (float64, float64) {
		return ls.spaceLeft, 0
	},
	ParagraphCenter: func(ls LineSize) (float64, float64) {
		return ls.spaceLeft / 2, 0
	},
	ParagraphJustify: func(ls LineSize) (float64, float64) {
		if ls.breakLine {
			return 0, 0
		}
		return 0, ls.spaceLeft / float64(ls.wordCount-1)
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

func (n *NodeParagraph) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	rb = rb.Cut(n.Frame)
	words, err := childrenToWords(pdf, n.children)
	if err != nil {
		return rb, err
	}
	lines := wordsToLines(words, rb.w)
	n.drawLines(pdf, rb, n.lineHeight, lines)
	return rb, nil
}

func (n *NodeParagraph) drawLines(pdf PdfDrawer, rb renderBox, lineHeight float64, lines []Line) {
	y := 0.0
	for _, l := range lines {
		ls := getLineSize(l, rb.w)
		x, spacing := n.align(ls)
		for _, w := range l.words {
			pdf.SetFont(w.fontName, w.fontSize)
			pdf.SetTextColor(w.color)
			pdf.Text(w.text, x+rb.x, y+rb.y, rb.w, lineHeight, AlingBaselineLeft)
			x = x + w.width + spacing + w.spaceWidth
		}
		y = y + n.lineHeight
	}
}

type Word struct {
	text       string
	width      float64
	fontSize   float64
	fontName   string
	color      color.RGBA
	spaceWidth float64
}

func childrenToWords(pdf PdfDrawer, nodes []Node) ([]Word, error) {
	words := []Word{}

	for _, child := range nodes {
		textChild, ok := child.(*NodeText)
		if !ok {
			return nil, fmt.Errorf("Unexpected node in paragraph")
		}

		words = append(words, textToWords(pdf, *textChild)...)
	}
	return words, nil
}

func textToWords(pdf PdfDrawer, node NodeText) []Word {
	pdf.SetFont(node.fontName, node.fontSize)
	spaceWidth := pdf.GetStringWidth(" ")

	splitted := strings.Split(node.text, "\n")
	words := []Word{}
	for i := 0; i < len(splitted); i++ {
		n := NodeText{
			text:     strings.ReplaceAll(splitted[i], utf8LineBreak, ""),
			fontSize: node.fontSize,
			fontName: node.fontName,
			color:    node.color,
		}
		words = append(words, lineToWords(spaceWidth, pdf, n)...)
		if i < (len(splitted) - 1) {
			words = append(words, Word{text: "\n"})
		}
	}
	return words
}

func lineToWords(spaceWidth float64, pdf PdfDrawer, node NodeText) []Word {
	words := []Word{}

	splitted := strings.Split(node.text, " ")
	for _, part := range splitted {
		if len(part) == 0 {
			continue
		}
		width := pdf.GetStringWidth(part)
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
	words     []Word
	breakLine bool
}

func wordsToLines(words []Word, width float64) []Line {
	x := 0.0
	lines := []Line{}
	line := Line{}
	for _, w := range words {
		if (x+w.width) > width || w.text == "\n" {
			if x == 0 && w.text != "\n" {
				continue
			}
			line.breakLine = w.text == "\n"
			lines = append(lines, line)
			line = Line{}
			x = 0
		}
		x = x + w.width + w.spaceWidth
		if w.text != "\n" {
			line.words = append(line.words, w)
		}
	}
	if len(line.words) > 0 {
		line.breakLine = true
		lines = append(lines, line)
	}
	return lines
}

type LineSize struct {
	spaceLeft float64
	wordCount int
	breakLine bool
}

func getLineSize(l Line, width float64) LineSize {
	if len(l.words) == 0 {
		return LineSize{}
	}
	realWidth := 0.0
	for _, w := range l.words {
		realWidth = realWidth + w.width + w.spaceWidth
	}
	realWidth = realWidth - l.words[len(l.words)-1].spaceWidth
	return LineSize{
		spaceLeft: width - realWidth,
		wordCount: len(l.words),
		breakLine: l.breakLine,
	}
}
