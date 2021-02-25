package renderer

import (
	"image/color"
	"io"
)

const (
	PdfWidthMm = 210
	PdfHeight  = 297
)

type PdfTextAlign string

const (
	AlingTopLeft        PdfTextAlign = "TopLeft"
	AlingTopCenter                   = "TopCenter"
	AlingTopRight                    = "TopRight"
	AlingMiddleLeft                  = "MiddleLeft"
	AlingMiddleCenter                = "MiddleCenter"
	AlingMiddleRight                 = "MiddleRight"
	AlingBottomLeft                  = "BottomLeft"
	AlingBottomCenter                = "BottomCenter"
	AlingBottomRight                 = "BottomRight"
	AlingBaselineLeft                = "BaselineLeft"
	AlingBaselineCenter              = "BaselineCenter"
	AlingBaselineRight               = "BaselineRight"
)

type Pdf interface {
	Init() PdfDrawer
}

type PdfDrawer interface {
	AddPage()
	SetFillColor(c color.RGBA)
	SetStrokeColor(c color.RGBA)
	SetStrokeWidth(w float64)
	Rect(x float64, y float64, width float64, height float64, radius float64)
	LoadFont(fontName string, fontFilePath string) error
	SetFont(fontName string, fontSizeMm float64)
	GetDefaultFontName() string
	SetTextColor(c color.RGBA)
	Text(text string, x float64, y float64, width float64, height float64, align PdfTextAlign)
	GetStringWidth(text string) float64
	Image(image io.ReadSeeker, x float64, y float64, width float64, height float64, keepAspectRation bool)
	Vector(image io.Reader, x float64, y float64, width float64, height float64)
	Output(out io.Writer) error
}
