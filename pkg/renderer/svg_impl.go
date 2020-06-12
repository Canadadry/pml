package renderer

import (
	"github.com/jung-kurt/gofpdf"

	"pml/pkg/renderer/svg"
)

type svgToPdf struct {
	pdf *gofpdf.Fpdf
}

func NewSvgToPdf(pdf *gofpdf.Fpdf) *svgToPdf {
	return &svgToPdf{
		pdf: pdf,
	}
}

func (s2p *svgToPdf) MoveTo(x float64, y float64) {
	s2p.pdf.MoveTo(x, y)
}

func (s2p *svgToPdf) LineTo(x float64, y float64) {
	s2p.pdf.LineTo(x, y)
}

func (s2p *svgToPdf) BezierTo(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) {

	s2p.pdf.CurveBezierCubicTo(x1, y1, x2, y2, x3, y3)
}

func (s2p *svgToPdf) CloseAndDraw(s svg.Style) {
	s2p.pdf.ClosePath()

	s2p.pdf.SetDrawColor(int(s.BorderColor.R), int(s.BorderColor.G), int(s.BorderColor.B))
	s2p.pdf.SetFillColor(int(s.FillColor.R), int(s.FillColor.G), int(s.FillColor.B))
	s2p.pdf.SetLineWidth(s.BorderSize)

	mode := ""

	if s.Fill {
		mode += "F"
	}
	if s.BorderSize > 0 {
		mode += "D"
	}

	s2p.pdf.DrawPath(mode)
}
