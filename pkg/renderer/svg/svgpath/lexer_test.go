package svgpath

import (
	"testing"
)

func TestNextToken_SingleCharToken(t *testing.T) {

	testedString := "m 234.43804,111.69821 c 50.21866,26.50627 126.75595,-3.87395 151.46369,-35.941621#"

	expectedTokens := []Token{
		{Type: COMMAND, Literal: "m"},
		{Type: NUMBER, Literal: "234.43804"},
		{Type: COMA, Literal: ","},
		{Type: NUMBER, Literal: "111.69821"},
		{Type: COMMAND, Literal: "c"},
		{Type: NUMBER, Literal: "50.21866"},
		{Type: COMA, Literal: ","},
		{Type: NUMBER, Literal: "26.50627"},
		{Type: NUMBER, Literal: "126.75595"},
		{Type: COMA, Literal: ","},
		{Type: NUMBER, Literal: "-3.87395"},
		{Type: NUMBER, Literal: "151.46369"},
		{Type: COMA, Literal: ","},
		{Type: NUMBER, Literal: "-35.941621"},
		{Type: ILLEGAL, Literal: "#"},
		{Type: EOF, Literal: ""},
	}

	lexer := newLexer(testedString)

	for _, expectedToken := range expectedTokens {
		actualToken := lexer.getNextToken()

		testToken(t, actualToken, expectedToken)

	}
}

func testToken(t *testing.T, actual Token, expected Token) {
	if actual.Type != expected.Type {
		t.Fatalf("token %v type wrong. expected=%q, got=%q", actual, expected.Type, actual.Type)
	}
	if actual.Literal != expected.Literal {
		t.Fatalf("token %v Literal wrong. expected=%q, got=%q", actual, expected.Literal, actual.Literal)
	}
}
