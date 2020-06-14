package renderer

import (
	"fmt"
	"github.com/canadadry/pml/pkg/abstract/abstractpdf"
	"github.com/canadadry/pml/pkg/domain/ast"
	"io"
	"os"
)

type renderer struct {
	output io.Writer
	pdf    abstractpdf.Pdf
}

func New(output io.Writer, pdf abstractpdf.Pdf) renderer {
	return renderer{
		output: output,
		pdf:    pdf,
	}
}

func (r *renderer) Render(tree *ast.Item) error {

	rt, err := GenerateFrom(tree)
	if err != nil {
		return err
	}
	return r.draw(rt, nil, 0, 0)
}

func (r *renderer) draw(node Node, pdf abstractpdf.Drawer, xOrigin float64, yOrigin float64) error {

	initialized := false
	renderChild := true

	switch n := node.(type) {
	case *NodeDocument:
		initialized = true
		pdf = r.pdf.Init()
	case *NodePage:
		pdf.AddPage()
	case *NodeRectangle:
		pdf.SetFillColor(n.color)
		pdf.Rect(n.x+xOrigin, n.y+yOrigin, n.width, n.height)
	case *NodeText:
		if len(n.fontName) == 0 {
			n.fontName = pdf.GetDefaultFontName()
		}
		pdf.SetFont(n.fontName, n.fontSize)
		pdf.SetTextColor(n.color)
		pdf.Text(n.text, n.x+xOrigin, n.y+yOrigin, n.width, n.height, abstractpdf.TextAlign(n.align))
	case *NodeFont:
		pdf.LoadFont(n.name, n.file)
	case *NodeImage:
		if len(n.file) == 0 {
			return fmt.Errorf("in image item, you must specify a property file")
		}
		file, err := os.Open(n.file)
		if err != nil {
			return err
		}
		defer file.Close()
		pdf.Image(file, n.x+xOrigin, n.y+yOrigin, n.width, n.height)
	case *NodeVector:
		if len(n.file) == 0 {
			return fmt.Errorf("in vector item, you must specify a property file")
		}
		file, err := os.Open(n.file)
		if err != nil {
			return err
		}
		defer file.Close()
		pdf.Vector(file, n.x+xOrigin, n.y+yOrigin, n.width, n.height)
	case *NodeParagraph:
		renderChild = false
		x := 0.0
		y := 0.0

		for _, child := range node.Chilrend() {
			offset := 0
			textChild, ok := child.(*NodeText)
			if !ok {
				return fmt.Errorf("Unexpected node in paragraph")
			}
			if len(textChild.fontName) == 0 {
				textChild.fontName = pdf.GetDefaultFontName()
			}

			pdf.SetFont(textChild.fontName, textChild.fontSize)
			pdf.SetTextColor(textChild.color)

			for offset < len(textChild.text) {
				maxSize, textWidth := pdf.GetTextMaxLength(textChild.text[offset:], n.width-x)
				text := textChild.text[offset : offset+maxSize]
				pdf.Text(text, n.x+x+xOrigin, n.y+y+yOrigin, n.width, n.lineHeight, "BaselineLeft")
				offset = offset + maxSize
				x = x + textWidth
				if x > n.width {
					x = 0
					y = y + n.lineHeight
				}
			}
		}
	case *NodeContainer:
		xOrigin = xOrigin + n.x
		yOrigin = yOrigin + n.y
	default:
		return fmt.Errorf("cannot render node type")
	}

	if renderChild {
		for _, child := range node.Chilrend() {
			err := r.draw(child, pdf, xOrigin, yOrigin)
			if err != nil {
				return err
			}
		}
	}

	if initialized {
		return pdf.Output(r.output)
	}

	return nil
}
