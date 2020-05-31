package renderer

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"pml/ast"
)

type textProperties struct {
	text string
}

func (r *renderer) extractTextProperties(text *ast.Item) (*textProperties, error) {

	tp := &textProperties{}

	for property, expression := range text.Properties {
		pptType, err := r.definitions.getPropertyType(itemText, property)
		if err != nil {
			return nil, err
		}
		if expression.Token().Type != pptType {
			return nil, fmt.Errorf("in textItem, %w exp %s, got %s", invalidTypeForProperty, pptType, expression.Token().Type)
		}
		switch property {
		case "text":
			tp.text = expression.Token().Literal
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

	pdf.CellFormat(190, 7, properties.text, "0", 0, "CM", false, 0, "")

	for _, child := range text.Children {

		if err := r.definitions.validateChildType(itemText, child.TokenType.Literal); err != nil {
			return err
		}
		switch child.TokenType.Literal {
		case itemText:
			if err := r.renderText(pdf, &child); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Cannot render %s: %w", child.TokenType.Literal, renderingItemNotImplemented)

		}
	}
	return nil
}
