package renderer

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"image/color"
	"pml/pkg/ast"
	"strconv"
)

type textProperties struct {
	text     string
	fontName string
	fontSize *float64
	x        *float64
	y        *float64
	width    *float64
	height   *float64
	color    *color.RGBA
	align    string
}

func (r *renderer) extractTextProperties(text *ast.Item) (*textProperties, error) {

	alignPossibleValue := map[string]string{
		"TopLeft":        "TL",
		"TopCenter":      "TC",
		"TopRight":       "TR",
		"MiddleLeft":     "LM",
		"MiddleCenter":   "CM",
		"MiddleRight":    "RM",
		"BottomLeft":     "BL",
		"BottomCenter":   "BC",
		"BottomRight":    "BR",
		"BaselineLeft":   "AL",
		"BaselineCenter": "AC",
		"BaselineRight":  "AR",
	}

	tp := &textProperties{}

	for property, expression := range text.Properties {
		pptType, err := r.definitions.getPropertyType(itemText, property)
		if err != nil {
			return nil, err
		}
		if expression.Token().Type != pptType {
			return nil, fmt.Errorf("in textItem, %w property %s exp %s, got %s", invalidTypeForProperty, property, pptType, expression.Token().Type)
		}
		switch property {
		case "text":
			tp.text = expression.Token().Literal
		case "fontName":
			tp.fontName = expression.Token().Literal
		case "fontSize":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			tp.fontSize = &value
		case "x":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			tp.x = &value
		case "y":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			tp.y = &value
		case "width":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			tp.width = &value
		case "height":
			value, err := strconv.ParseFloat(expression.Token().Literal, 64)
			if err != nil {
				return nil, err
			}
			tp.height = &value
		case "color":
			value, err := parseHexColor(expression.Token().Literal)
			if err != nil {
				return nil, err
			}
			tp.color = &value
		case "align":
			align := expression.Token().Literal
			alignValue, ok := alignPossibleValue[align]
			if !ok {
				return nil, fmt.Errorf("%s is not a valid value for align property of text", align)
			}
			tp.align = alignValue
		default:
			return nil, fmt.Errorf("Cannot extract in textItem %s: %w", property, invalidTypeForProperty)

		}
	}
	return tp, nil
}

func (r *renderer) renderText(pdf *gofpdf.Fpdf, text *ast.Item) error {

	properties, err := r.extractTextProperties(text)
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

	if len(properties.align) != 2 {
		properties.align = "TL"
	}

	if len(properties.fontName) == 0 {
		properties.fontName = "Arial"
	}
	if properties.fontSize == nil {
		defaultValue := float64(16)
		properties.fontSize = &defaultValue
	}

	pdf.SetFont(properties.fontName, "", *properties.fontSize)
	pdf.SetTextColor(int(properties.color.R), int(properties.color.G), int(properties.color.B))
	pdf.SetXY(*properties.x, *properties.y)
	pdf.CellFormat(*properties.width, *properties.height, properties.text, "", 0, properties.align, false, 0, "")

	for _, child := range text.Children {

		if err := r.definitions.validateChildType(itemText, child.TokenType.Literal); err != nil {
			return err
		}
		switch child.TokenType.Literal {
		default:
			return fmt.Errorf("Cannot render %s: %w", child.TokenType.Literal, renderingItemNotImplemented)

		}
	}
	return nil
}
