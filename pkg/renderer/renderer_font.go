package renderer

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	"path/filepath"
	"pml/pkg/ast"
)

type fontProperties struct {
	file string
	name string
}

func (r *renderer) extractFontProperties(font *ast.Item) (*fontProperties, error) {

	fp := &fontProperties{}

	for property, expression := range font.Properties {
		pptType, err := r.definitions.getPropertyType(itemFont, property)
		if err != nil {
			return nil, err
		}
		if expression.Token().Type != pptType {
			return nil, fmt.Errorf("in textItem, %w property %s exp %s, got %s", invalidTypeForProperty, property, pptType, expression.Token().Type)
		}
		switch property {
		case "file":
			fp.file = expression.Token().Literal
		case "name":
			fp.name = expression.Token().Literal
		default:
			return nil, fmt.Errorf("Cannot extract in textItem %s: %w", property, invalidTypeForProperty)

		}
	}
	return fp, nil
}

func (r *renderer) renderFont(pdf *gofpdf.Fpdf, font *ast.Item) error {

	properties, err := r.extractFontProperties(font)
	if err != nil {
		return err
	}
	if len(properties.file) == 0 {
		return fmt.Errorf("in font item, you must specify a property file")
	}
	if len(properties.name) == 0 {
		return fmt.Errorf("in font item, you must specify a property name")
	}

	fmt.Printf("%#v\n", properties)

	dir := filepath.Dir(properties.file)

	err = gofpdf.MakeFont(properties.file, dir+"/cp1258.map", dir, os.Stdout, true)
	if err != nil {
		return err
	}

	pdf.AddUTF8Font(properties.name, "", properties.file)
	return nil
}
