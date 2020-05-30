package parser

import (
	"errors"
	"pml/ast"
	"pml/lexer"
	"pml/token"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		program  string
		expected *ast.Item
	}{
		{"Doc{}", &ast.Item{TokenType: tokenIdentifier("Doc")}},
	}

	for i, tt := range tests {
		l := lexer.New(tt.program)
		parser := New(l)

		result, err := parser.Parse()
		if err != nil {
			t.Fatalf("%v", err)
		}
		testItem(t, i, result, tt.expected)
	}
}

func TestParserError(t *testing.T) {
	tests := []struct {
		program  string
		expected error
	}{
		{"{}", errNextTokenIsNotTheExpectedOne},
		{"a}", errNextTokenIsNotTheExpectedOne},
		{"a{a", errNextTokenIsNotTheExpectedOne},
	}

	for _, tt := range tests {
		l := lexer.New(tt.program)
		parser := New(l)

		_, err := parser.Parse()
		if !errors.Is(err, tt.expected) {
			t.Fatalf("error was not the one expected got %s, exp %s", err, tt.expected)
		}
	}
}

func tokenIdentifier(literal string) token.Token {
	return token.Token{
		Type:    token.IDENTIFIER,
		Literal: literal,
	}
}

func testItem(t *testing.T, index int, actual *ast.Item, expected *ast.Item) {
	if actual.TokenType.Type != expected.TokenType.Type {
		t.Fatalf("[%d] Wrong item type got %s expected %s", index, actual.TokenType.Type, expected.TokenType.Type)
	}
	if actual.TokenType.Literal != expected.TokenType.Literal {
		t.Fatalf("[%d] Wrong item literal got %s expected %s", index, actual.TokenType.Literal, expected.TokenType.Literal)
	}

}
