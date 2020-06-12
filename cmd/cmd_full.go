package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/canadadry/pml/pkg/lexer"
	"github.com/canadadry/pml/pkg/parser"
	"github.com/canadadry/pml/pkg/renderer"
	"github.com/canadadry/pml/pkg/template"
	"io"
)

func Full(input string, output io.Writer, param []byte) error {

	var dat interface{}
	if err := json.Unmarshal(param, &dat); err != nil {
		return fmt.Errorf("Cannot unmarshall json file : %w\n", err)
	}
	out, err := template.Apply(input, dat)
	if err != nil {
		return fmt.Errorf("failed to transform template : %w\n", err)
	}

	l := lexer.New(out)
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
