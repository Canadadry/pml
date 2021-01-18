package language

import (
	"fmt"
	"github.com/canadadry/pml/language/lexer"
	"github.com/canadadry/pml/language/parser"
	"github.com/canadadry/pml/language/renderer"
	"github.com/canadadry/pml/pkg/pdf"
	"github.com/canadadry/pml/pkg/svg"
	"github.com/canadadry/pml/pkg/template"
	"io"
)

func Run(input io.Reader, output io.Writer, param io.Reader) error {
	out, err := template.ApplyJson(input, param)
	if err != nil {
		return fmt.Errorf("failed to transform template : %w\n", err)
	}

	l := lexer.New(out)
	p := parser.New(l)
	r := renderer.New(output, pdf.New(svg.New()))

	item, err := p.Parse()
	if err != nil {
		return fmt.Errorf("parsing failed : %w on : \n%s", err, out)
	}
	err = r.Render(item)
	if err != nil {
		return fmt.Errorf("rendering failed : %w", err)
	}
	return nil
}
