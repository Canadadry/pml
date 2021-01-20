package svgdrawer

import (
	"fmt"
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

type DrawFunc func(d Drawer, svg io.Reader, x float64, y float64, width float64, height float64) error

type ForTesting struct {
	Callstack []string
}

func (dcs *ForTesting) MoveTo(x float64, y float64) {
	dcs.Callstack = append(dcs.Callstack,
		fmt.Sprintf("MoveTo x:%g, y:%g", x, y),
	)
}

func (dcs *ForTesting) LineTo(x float64, y float64) {
	dcs.Callstack = append(dcs.Callstack,
		fmt.Sprintf("LineTo x:%g, y:%g", x, y),
	)
}

func (dcs *ForTesting) BezierTo(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) {
	dcs.Callstack = append(dcs.Callstack,
		fmt.Sprintf("BezierTo %g,%g, anchor 1 %g,%g anchor 2 %g,%g", x3, y3, x1, y1, x2, y2),
	)
}

func (dcs *ForTesting) CloseAndDraw(s Style) {
	dcs.Callstack = append(dcs.Callstack,
		fmt.Sprintf("CloseAndDraw"),
	)
}
