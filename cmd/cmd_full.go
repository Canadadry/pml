package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/canadadry/pml/pkg/adapter/pdf"
	"github.com/canadadry/pml/pkg/adapter/svg"
	"github.com/canadadry/pml/pkg/domain/lexer"
	"github.com/canadadry/pml/pkg/domain/parser"
	"github.com/canadadry/pml/pkg/domain/renderer"
	"github.com/canadadry/pml/pkg/domain/template"
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
