package renderer

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	"pml/pkg/ast"
	"pml/pkg/renderer/svg"
	"strconv"
)

type vectorProperties struct {
	x      *float64
	y      *float64
	width  *float64
	height *float64
	file   string
}

func (r *renderer) extractVectorProperties(vector *ast.Item) (*vectorProperties, error) {

	ip := &vectorProperties{}

	for property, expression := range vector.Properties {
		pptType, err := r.definitions.getPropertyType(itemVector, property)
		if err != nil {
			return nil, err
		}
		if expression.Token().Type != pptType {
			return nil, fmt.Errorf("in textItem, %w property %s exp %s, got %s", invalidTypeForProperty, property, pptType, expression.Token().Type)
		}
		switch property {
		case "x":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			ip.x = &value
		case "y":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			ip.y = &value
		case "width":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			ip.width = &value
		case "height":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			ip.height = &value
		case "file":
			ip.file = expression.Token().Literal
		default:
			return nil, fmt.Errorf("Cannot extract in textItem %s: %w", property, invalidTypeForProperty)

		}
	}
	return ip, nil
}

func (r *renderer) renderVector(pdf *gofpdf.Fpdf, vector *ast.Item) error {

	properties, err := r.extractVectorProperties(vector)
	if err != nil {
		return err
	}
	if properties.x == nil {
		defaultValue := float64(0)
		properties.x = &defaultValue
	}

	if properties.y == nil {
		defaultValue := float64(0)
		properties.y = &defaultValue
	}
	if properties.width == nil {
		defaultValue := float64(0)
		properties.width = &defaultValue
	}

	if properties.height == nil {
		defaultValue := float64(0)
		properties.height = &defaultValue
	}

	if len(properties.file) == 0 {
		return fmt.Errorf("in vector item, you must specify a property file")
	}
	svgFile, err := os.Open(properties.file)
	if err != nil {
		return err
	}
	defer svgFile.Close()

	return svg.Draw(NewSvgToPdf(pdf), svgFile, *properties.x, *properties.y, *properties.width, *properties.height)
}

type svgToPdf struct {
	pdf *gofpdf.Fpdf
}

func NewSvgToPdf(pdf *gofpdf.Fpdf) *svgToPdf {
	return &svgToPdf{
		pdf: pdf,
	}
}

func (s2p *svgToPdf) SetStyle(s svg.Style) {
	s2p.pdf.SetDrawColor(0xff, 0x00, 0x00)
	s2p.pdf.SetFillColor(0x99, 0x99, 0x99)
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

func (s2p *svgToPdf) CloseAndDraw() {
	s2p.pdf.ClosePath()
	s2p.pdf.SetFillColor(200, 200, 200)
	s2p.pdf.SetLineWidth(3)
	s2p.pdf.DrawPath("DF")
}
