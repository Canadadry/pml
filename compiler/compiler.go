package compiler

import (
	"fmt"
	"github.com/canadadry/pml/compiler/lexer"
	"github.com/canadadry/pml/compiler/parser"
	"github.com/canadadry/pml/compiler/renderer"
	"github.com/canadadry/pml/pkg/template"
	"io"
)

type Env struct {
	Input  map[string]io.Reader
	Main   string
	Output io.Writer
	Param  io.Reader
	Pdf    renderer.Pdf
	Funcs  map[string]interface{}
}

func Run(e Env) error {
	out, err := template.ApplyJson(e.Input, e.Main, e.Param, e.Funcs)
	if err != nil {
		return fmt.Errorf("failed to transform template : %w\n", err)
	}

	l := lexer.New(out)
	p := parser.New(l)
	item, err := p.Parse()
	if err != nil {
		return fmt.Errorf("parsing failed : %w on : \n%s", err, out)
	}

	r := renderer.New(e.Output, e.Pdf)
	err = r.Render(item)
	if err != nil {
		return fmt.Errorf("rendering failed : %w", err)
	}
	return nil
}
