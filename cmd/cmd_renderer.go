package cmd

import (
	"fmt"
	"github.com/canadadry/pml/pkg/adapter/svg"
	"github.com/canadadry/pml/pkg/domain/lexer"
	"github.com/canadadry/pml/pkg/domain/parser"
	"github.com/canadadry/pml/pkg/domain/renderer"
	"io"
)

func Renderer(file string, output io.Writer) error {

	l := lexer.New(string(file))
	p := parser.New(l)
	r := renderer.New(output, svg.New())

	item, err := p.Parse()
	if err != nil {
		return fmt.Errorf("parsing failed : %w", err)
	}
	err = r.Render(item)
	if err != nil {
		return fmt.Errorf("rendering failed : %w", err)
	}
	return nil
}
