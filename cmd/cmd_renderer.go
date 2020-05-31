package cmd

import (
	"fmt"
	"pml/pkg/lexer"
	"pml/pkg/parser"
	"pml/pkg/renderer"
)

func Renderer(file string, output string) error {

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
