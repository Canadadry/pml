package abstract

import (
	"image/color"
	"io"
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

type Svg interface {
	Draw(d Drawer, svg io.Reader, x float64, y float64, width float64, height float64) error
}
