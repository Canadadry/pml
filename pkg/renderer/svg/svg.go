package svg

import (
	"image/color"
	"io"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"pml/pkg/renderer/svg/svgpath"
)

type Style struct {
	Fill        bool
	FillColor   color.RGBA
	BorderSize  int
	BorderColor color.RGBA
}

type Drawer interface {
	SetStyle(s Style)
	MoveTo(x float64, y float64)
	LineTo(x float64, y float64)
	BezierTo(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64)
	CloseAndDraw()
}

type svgNode struct {
	worldToLocal matrix.Matrix
	commands     []svgpath.Command
	children     []*svgNode
}

func (sn *svgNode) draw(d Drawer) error {

	for _, cmd := range sn.commands {
		switch cmd.Kind {
		case 'M':
			d.MoveTo(cmd.Points[0].X, cmd.Points[0].Y)
		case 'L':
			d.LineTo(cmd.Points[0].X, cmd.Points[0].Y)
		case 'C':
			d.BezierTo(cmd.Points[0].X, cmd.Points[0].Y, cmd.Points[1].X, cmd.Points[1].Y, cmd.Points[2].X, cmd.Points[2].Y)
		case 'Z':
			d.CloseAndDraw()
		}
	}

	for _, child := range sn.children {
		err := child.draw(d)
		if err != nil {
			return err
		}
	}
	return nil
}

func Draw(d Drawer, svg io.Reader, x float64, y float64, w float64, h float64) error {

	element, err := svgparser.Parse(svg)
	if err != nil {
		return err
	}

	root, err := svgMain(element, viewBox{x, y, w, h})
	if err != nil {
		return err
	}

	return root.draw(d)
}
