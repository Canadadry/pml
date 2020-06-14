package abstractpdf

import (
	"github.com/canadadry/pml/pkg/abstract/abstractsvg"
	"image/color"
	"io"
)

type TextAlign string

const (
	AlingTopLeft        TextAlign = "TopLeft"
	AlingTopCenter                = "TopCenter"
	AlingTopRight                 = "TopRight"
	AlingMiddleLeft               = "MiddleLeft"
	AlingMiddleCenter             = "MiddleCenter"
	AlingMiddleRight              = "MiddleRight"
	AlingBottomLeft               = "BottomLeft"
	AlingBottomCenter             = "BottomCenter"
	AlingBottomRight              = "BottomRight"
	AlingBaselineLeft             = "BaselineLeft"
	AlingBaselineCenter           = "BaselineCenter"
	AlingBaselineRight            = "BaselineRight"
)

type Pdf interface {
	Init() Drawer
}

type Drawer interface {
	abstractsvg.Drawer
	AddPage()
	SetFillColor(c color.RGBA)
	Rect(x float64, y float64, width float64, height float64)
	LoadFont(fontName string, fontFilePath string) error
	SetFont(fontName string, fontSizeMm float64)
	SetTextColor(c color.RGBA)
	Text(text string, x float64, y float64, width float64, height float64, align TextAlign)
	GetTextMaxLength(text string, maxWidth float64) (int, float64)
	Image(image io.ReadSeeker, x float64, y float64, width float64, height float64)
	Vector(image io.Reader, x float64, y float64, width float64, height float64)
	Output(out io.Writer) error
}
