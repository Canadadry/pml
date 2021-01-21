package renderer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"io"
	"os"
)

var (
	ErrEmptyImageFileProperty = fmt.Errorf("in image item, you must specify a property file")
	ErrCantOpenFile           = fmt.Errorf("ErrCantOpenFile")
	ErrB64Read                = fmt.Errorf("ErrB64Read")
)

type renderer struct {
	output io.Writer
	pdf    Pdf
}

func New(output io.Writer, pdf Pdf) renderer {
	return renderer{
		output: output,
		pdf:    pdf,
	}
}

func (r *renderer) Render(tree *ast.Item) error {

	rt, err := generate(tree)
	if err != nil {
		return err
	}
	return r.draw(rt, nil, renderBox{})
}

type renderBox struct {
	x float64
	y float64
	w float64
	h float64
}

func (rb renderBox) Cut(x, y, w, h float64) renderBox {
	rb.x = rb.x + x
	rb.y = rb.y + y
	rb.w = w
	rb.h = h
	return rb
}

func (r *renderer) draw(node Node, pdf PdfDrawer, rb renderBox) error {

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
		rb = rb.Cut(n.x, n.y, n.width, n.height)
		pdf.Rect(rb.x, rb.y, rb.w, rb.h)
	case *NodeText:
		if len(n.fontName) == 0 {
			n.fontName = pdf.GetDefaultFontName()
		}
		pdf.SetFont(n.fontName, n.fontSize)
		pdf.SetTextColor(n.color)
		rb = rb.Cut(n.x, n.y, n.width, n.height)
		pdf.Text(n.text, rb.x, rb.y, rb.w, rb.h, PdfTextAlign(n.align))
	case *NodeFont:
		pdf.LoadFont(n.name, n.file)
	case *NodeImage:
		if len(n.file) == 0 {
			return ErrEmptyImageFileProperty
		}
		var rs io.ReadSeeker
		if n.mode == ImgModeFile {
			file, err := os.Open(n.file)
			if err != nil {
				return fmt.Errorf("%w : %v", ErrCantOpenFile, err)
			}
			defer file.Close()
			rs = file
		} else {
			decoded, err := base64.StdEncoding.DecodeString(n.file)
			if err != nil {
				return fmt.Errorf("%w : %v", ErrB64Read, err)
			}
			rs = bytes.NewReader(decoded)
		}
		rb = rb.Cut(n.x, n.y, n.width, n.height)
		pdf.Image(rs, rb.x, rb.y, rb.w, rb.h)
	case *NodeVector:
		if len(n.file) == 0 {
			return ErrEmptyImageFileProperty
		}
		file, err := os.Open(n.file)
		if err != nil {
			return fmt.Errorf("%w : %v", ErrCantOpenFile, err)
		}
		defer file.Close()
		rb = rb.Cut(n.x, n.y, n.width, n.height)
		pdf.Vector(file, rb.x, rb.y, rb.w, rb.h)
	case *NodeParagraph:
		renderChild = false
		x := 0.0
		y := 0.0
		rb = rb.Cut(n.x, n.y, n.width, n.height)

		for _, child := range node.Children() {
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
				pdf.Text(text, x+rb.x, y+rb.y, rb.w, n.lineHeight, "BaselineLeft")
				offset = offset + maxSize
				x = x + textWidth
				if x > rb.w {
					x = 0
					y = y + n.lineHeight
				}
			}
		}
	case *NodeContainer:
		rb = rb.Cut(n.x, n.y, rb.w, rb.h)
	default:
		return fmt.Errorf("cannot render node type")
	}

	if renderChild {
		for _, child := range node.Children() {
			err := r.draw(child, pdf, rb)
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
