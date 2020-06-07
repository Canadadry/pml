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
	BorderSize  float64
	BorderColor color.RGBA
}

type Drawer interface {
	MoveTo(x float64, y float64)
	LineTo(x float64, y float64)
	BezierTo(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64)
	CloseAndDraw(s Style)
}

type svgNode struct {
	worldToLocal matrix.Matrix
	commands     []svgpath.Command
	style        Style
	children     []*svgNode
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
