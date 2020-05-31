package renderer

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io"
	"pml/pkg/ast"
)

type renderer struct {
	output      io.Writer
	definitions itemDefinitions
}

func New(output io.Writer) renderer {
	return renderer{
		output:      output,
		definitions: items,
	}
}

func (r *renderer) Render(tree *ast.Item) error {

	if tree.TokenType.Literal != itemDocument {
		return fmt.Errorf("%w : got %s exp %s", rootMustBeDocumentItem, tree.TokenType.Literal, itemDocument)
	}
	return r.renderDocument(tree)
}

func (r *renderer) renderDocument(document *ast.Item) error {

	pdf := gofpdf.New("P", "mm", "A4", "")

	for _, child := range document.Children {

		if err := r.definitions.validateChildType(itemDocument, child.TokenType.Literal); err != nil {
			return err
		}
		switch child.TokenType.Literal {
		case itemPage:
			if err := r.renderPage(pdf, &child); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Cannot render %s: %w", child.TokenType.Literal, renderingItemNotImplemented)
		}
	}

	return pdf.Output(r.output)
}

func (r *renderer) renderPage(pdf *gofpdf.Fpdf, page *ast.Item) error {

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	for _, child := range page.Children {

		if err := r.definitions.validateChildType(itemPage, child.TokenType.Literal); err != nil {
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
		default:
			return fmt.Errorf("Cannot render %s: %w", child.TokenType.Literal, renderingItemNotImplemented)

		}
	}
	return nil
}
