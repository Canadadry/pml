package svg

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"pml/pkg/renderer/svg/matrix"
)

type command struct {
	kind byte
	x1   float64
	y1   float64
	x2   float64
	y2   float64
	x3   float64
	y3   float64
}

type svgNode struct {
	worldToLocal matrix.Matrix
	commands     []command
	children     []*svgNode
}

func (sn *svgNode) draw(pdf *gofpdf.Fpdf) error {

	for _, cmd := range sn.commands {
		fmt.Printf("drawing %s, %g, %g\n", string(cmd.kind), cmd.x1, cmd.y2)
		switch cmd.kind {
		case 'M':
			pdf.MoveTo(cmd.x1, cmd.y1)
		case 'L':
			pdf.LineTo(cmd.x1, cmd.y1)
		case 'C':
			pdf.LineTo(cmd.x1, cmd.y1)
		case 'Z':
			pdf.ClosePath()
			pdf.DrawPath("FD")
		}
	}

	for _, child := range sn.children {
		err := child.draw(pdf)
		if err != nil {
			return err
		}
	}
	return nil
}
