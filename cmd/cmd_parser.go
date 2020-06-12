package cmd

import (
	"fmt"
	"github.com/canadadry/pml/pkg/lexer"
	"github.com/canadadry/pml/pkg/parser"
)

func Parser(file string) error {

	l := lexer.New(file)
	p := parser.New(l)

	item, err := p.Parse()
	if err != nil {
		return fmt.Errorf("parsing failed : %w", err)
	}
	fmt.Println(item)
	return nil
}
