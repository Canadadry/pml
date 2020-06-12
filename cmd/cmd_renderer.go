package cmd

import (
	"fmt"
	"github.com/canadadry/pml/pkg/lexer"
	"github.com/canadadry/pml/pkg/parser"
	"github.com/canadadry/pml/pkg/renderer"
	"io"
)

func Renderer(file string, output io.Writer) error {

	l := lexer.New(string(file))
	p := parser.New(l)
	r := renderer.New(output)

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
