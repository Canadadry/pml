package renderer

import (
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"io"
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

func (r *renderer) draw(node Node, pdf PdfDrawer, rb renderBox) error {

	rb, err := node.draw(pdf, rb)
	if err != nil {
		return err
	}

	if node.needToDrawChild() {
		for _, child := range node.Children() {
			err := r.draw(child, pdf, rb)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
