package renderer

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"pml/pkg/ast"
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

	sig, err := gofpdf.SVGBasicFileParse(properties.file)
	if err != nil {
		return err
	}

	scale := 100 / sig.Wd
	scaleY := 30 / sig.Ht
	if scale > scaleY {
		scale = scaleY
	}
	pdf.SetLineCapStyle("round")
	pdf.SetLineWidth(0.25)
	pdf.SetDrawColor(0, 0, 128)
	pdf.SetXY(*properties.x, *properties.x)
	pdf.SVGBasicWrite(&sig, scale)

	return nil
}
