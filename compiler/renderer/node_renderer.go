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
	pdf := r.pdf.Init()
	err = r.draw(rt, pdf, renderBox{})
	if err != nil {
		return err
	}
	return pdf.Output(r.output)
}

type renderBox struct {
	x float64
	y float64
	w float64
	h float64
}

func (rb renderBox) Cut(f Frame) renderBox {
	rb.x = rb.x + f.x
	rb.y = rb.y + f.y
	rb.w = f.width
	rb.h = f.height
	return rb
}

func (r *renderer) draw(node Node, pdf PdfDrawer, rb renderBox) error {
	var err error
	renderChild := true

	switch n := node.(type) {
	case *NodeDocument:
		rb, err = drawDocument(*n, pdf, rb)
	case *NodePage:
		rb, err = drawPage(*n, pdf, rb)
	case *NodeRectangle:
		rb, err = drawRectangle(*n, pdf, rb)
	case *NodeText:
		rb, err = drawText(*n, pdf, rb)
	case *NodeFont:
		rb, err = drawFont(*n, pdf, rb)
	case *NodeImage:
		rb, err = drawImage(*n, pdf, rb)
	case *NodeVector:
		rb, err = drawVector(*n, pdf, rb)
	case *NodeParagraph:
		renderChild = false
		rb, err = drawParagraph(*n, pdf, rb)
	case *NodeContainer:
		rb, err = drawContainer(*n, pdf, rb)
	default:
		return fmt.Errorf("cannot render node type")
	}
	if err != nil {
		return err
	}

	if renderChild {
		for _, child := range node.Children() {
			err := r.draw(child, pdf, rb)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func drawDocument(n NodeDocument, pdf PdfDrawer, rb renderBox) (renderBox, error) {
	return rb, nil
}
func drawContainer(n NodeContainer, pdf PdfDrawer, rb renderBox) (renderBox, error) {
	return rb.Cut(n.Frame), nil
}
func drawPage(n NodePage, pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.AddPage()
	return renderBox{w: PdfWidthMm, h: PdfHeight}, nil
}

func drawRectangle(n NodeRectangle, pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.SetFillColor(n.color)
	rb = rb.Cut(n.Frame)
	pdf.Rect(rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}

func drawText(n NodeText, pdf PdfDrawer, rb renderBox) (renderBox, error) {

	if len(n.fontName) == 0 {
		n.fontName = pdf.GetDefaultFontName()
	}
	pdf.SetFont(n.fontName, n.fontSize)
	pdf.SetTextColor(n.color)
	rb = rb.Cut(n.Frame)
	pdf.Text(n.text, rb.x, rb.y, rb.w, rb.h, PdfTextAlign(n.align))
	return rb, nil
}

func drawFont(n NodeFont, pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.LoadFont(n.name, n.file)
	return rb, nil
}

func drawImage(n NodeImage, pdf PdfDrawer, rb renderBox) (renderBox, error) {
	if len(n.file) == 0 {
		return rb, ErrEmptyImageFileProperty
	}
	var rs io.ReadSeeker
	if n.mode == ImgModeFile {
		file, err := os.Open(n.file)
		if err != nil {
			return rb, fmt.Errorf("%w : %v", ErrCantOpenFile, err)
		}
		defer file.Close()
		rs = file
	} else {
		decoded, err := base64.StdEncoding.DecodeString(n.file)
		if err != nil {
			return rb, fmt.Errorf("%w : %v", ErrB64Read, err)
		}
		rs = bytes.NewReader(decoded)
	}
	rb = rb.Cut(n.Frame)
	pdf.Image(rs, rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}

func drawVector(n NodeVector, pdf PdfDrawer, rb renderBox) (renderBox, error) {
	if len(n.file) == 0 {
		return rb, ErrEmptyImageFileProperty
	}
	file, err := os.Open(n.file)
	if err != nil {
		return rb, fmt.Errorf("%w : %v", ErrCantOpenFile, err)
	}
	defer file.Close()
	rb = rb.Cut(n.Frame)
	pdf.Vector(file, rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}

func drawParagraph(n NodeParagraph, pdf PdfDrawer, rb renderBox) (renderBox, error) {
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
