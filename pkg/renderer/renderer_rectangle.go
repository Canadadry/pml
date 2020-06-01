package renderer

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"image/color"
	"pml/pkg/ast"
	"strconv"
)

type rectangleProperties struct {
	x      *float64
	y      *float64
	width  *float64
	height *float64
	color  *color.RGBA
}

func (r *renderer) extractRectangleProperties(rectangle *ast.Item) (*rectangleProperties, error) {

	rp := &rectangleProperties{}

	for property, expression := range rectangle.Properties {
		pptType, err := r.definitions.getPropertyType(itemRectangle, property)
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
			rp.x = &value
		case "y":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			rp.y = &value
		case "width":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			rp.width = &value
		case "height":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			rp.height = &value
		case "color":
			value, err := parseHexColor(expression.Token().Literal)
			if err != nil {
				return nil, err
			}
			rp.color = &value
		default:
			return nil, fmt.Errorf("Cannot extract in textItem %s: %w", property, invalidTypeForProperty)

		}
	}
	return rp, nil
}

func (r *renderer) renderRectangle(pdf *gofpdf.Fpdf, rectangle *ast.Item) error {

	properties, err := r.extractRectangleProperties(rectangle)
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
		defaultValue := float64(210)
		properties.width = &defaultValue
	}

	if properties.height == nil {
		defaultValue := float64(297)
		properties.height = &defaultValue
	}

	if properties.color == nil {
		defaultValue := color.RGBA{A: 0xff}
		properties.color = &defaultValue
	}

	pdf.SetFillColor(int(properties.color.R), int(properties.color.G), int(properties.color.B))
	pdf.Rect(*properties.x, *properties.y, *properties.width, *properties.height, "F")

	for _, child := range rectangle.Children {

		if err := r.definitions.validateChildType(itemText, child.TokenType.Literal); err != nil {
			return err
		}
		switch child.TokenType.Literal {
		case itemText:
			if err := r.renderText(pdf, &child); err != nil {
				return err
			}
		case itemRectangle:
			if err := r.renderRectangle(pdf, &child); err != nil {
				return err
			}
		case itemImage:
			if err := r.renderImage(pdf, &child); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Cannot render %s: %w", child.TokenType.Literal, renderingItemNotImplemented)

		}
	}
	return nil
}
