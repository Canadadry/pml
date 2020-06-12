package cmd

import (
	"fmt"
	"github.com/canadadry/pml/pkg/lexer"
	"github.com/canadadry/pml/pkg/token"
)

func Lexer(input string) {
	l := lexer.New(input)

	tok := l.GetNextToken()
	for tok.Type != token.EOF {
		fmt.Println(tok)
		tok = l.GetNextToken()
	}
}
