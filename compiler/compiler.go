package compiler

import (
	"fmt"
	"github.com/canadadry/pml/compiler/lexer"
	"github.com/canadadry/pml/compiler/parser"
	"github.com/canadadry/pml/compiler/renderer"
	"github.com/canadadry/pml/pkg/template"
	"io"
)

func Run(input io.Reader, output io.Writer, param io.Reader, pdf renderer.Pdf) error {
	out, err := template.ApplyJson(input, param)
	if err != nil {
		return fmt.Errorf("failed to transform template : %w\n", err)
	}

	l := lexer.New(out)
	p := parser.New(l)
	item, err := p.Parse()
	if err != nil {
		return fmt.Errorf("parsing failed : %w on : \n%s", err, out)
	}

	r := renderer.New(output, pdf)
	err = r.Render(item)
	if err != nil {
		return fmt.Errorf("rendering failed : %w", err)
	}
	return nil
}
