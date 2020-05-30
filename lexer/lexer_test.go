package lexer

import (
	"pml/token"
	"testing"
)

func TestNextToken_SingleCharToken(t *testing.T) {

	testedString := `Font{
	id:title
	familly:"helevetica"
	size:12
	weight:3.3
	color:#fafafaff
	color:#fafafa
	fake:#fafaf
	+
	12a
}`

	expectedTokens := []token.Token{
		{Type: token.IDENTIFIER, Literal: "Font", Line: 1, Column: 1},
		{Type: token.LEFT_BRACE, Literal: "{", Line: 1, Column: 5},
		{Type: token.IDENTIFIER, Literal: "id", Line: 2, Column: 2},
		{Type: token.DOTS, Literal: ":", Line: 2, Column: 4},
		{Type: token.IDENTIFIER, Literal: "title", Line: 2, Column: 5},
		{Type: token.IDENTIFIER, Literal: "familly", Line: 3, Column: 2},
		{Type: token.DOTS, Literal: ":", Line: 3, Column: 9},
		{Type: token.STRING, Literal: "helevetica", Line: 3, Column: 10},
		{Type: token.IDENTIFIER, Literal: "size", Line: 4, Column: 2},
		{Type: token.DOTS, Literal: ":", Line: 4, Column: 6},
		{Type: token.INTEGER, Literal: "12", Line: 4, Column: 7},
		{Type: token.IDENTIFIER, Literal: "weight", Line: 5, Column: 2},
		{Type: token.DOTS, Literal: ":", Line: 5, Column: 8},
		{Type: token.FLOAT, Literal: "3.3", Line: 5, Column: 9},
		{Type: token.IDENTIFIER, Literal: "color", Line: 6, Column: 2},
		{Type: token.DOTS, Literal: ":", Line: 6, Column: 7},
		{Type: token.COLOR, Literal: "fafafaff", Line: 6, Column: 8},
		{Type: token.IDENTIFIER, Literal: "color", Line: 7, Column: 2},
		{Type: token.DOTS, Literal: ":", Line: 7, Column: 7},
		{Type: token.COLOR, Literal: "fafafa", Line: 7, Column: 8},
		{Type: token.IDENTIFIER, Literal: "fake", Line: 8, Column: 2},
		{Type: token.DOTS, Literal: ":", Line: 8, Column: 6},
		{Type: token.ILLEGAL, Literal: "fafaf", Line: 8, Column: 7},
		{Type: token.ILLEGAL, Literal: "+", Line: 9, Column: 2},
		{Type: token.ILLEGAL, Literal: "12a", Line: 10, Column: 2},
		{Type: token.RIGHT_BRACE, Literal: "}", Line: 11, Column: 1},
		{Type: token.EOF, Literal: "", Line: 11, Column: 2},
	}

	lexer := New(testedString)

	for _, expectedToken := range expectedTokens {
		actualToken := lexer.GetNextToken()

		testToken(t, actualToken, expectedToken)

	}
}

func testToken(t *testing.T, actual token.Token, expected token.Token) {
	if actual.Type != expected.Type {
		t.Fatalf("token %v type wrong. expected=%q, got=%q", actual, expected.Type, actual.Type)
	}
	if actual.Literal != expected.Literal {
		t.Fatalf("token %v Literal wrong. expected=%q, got=%q", actual, expected.Literal, actual.Literal)
	}
	if actual.Line != expected.Line {
		t.Fatalf("token %v Line wrong. expected=%d, got=%d", actual, expected.Line, actual.Line)
	}
	if actual.Column != expected.Column {
		t.Fatalf("token %v Column wrong. expected=%d, got=%d", actual, expected.Column, actual.Column)
	}
}
