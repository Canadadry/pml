package eval

import (
	"testing"
)

func TestNextToken_SingleCharToken(t *testing.T) {

	testedString := "()+* /-**//\n12.32%#"

	expectedTokens := []Token{
		{Type: LEFT_PAR, Literal: "(", Pos: 0},
		{Type: RIGHT_PAR, Literal: ")", Pos: 1},
		{Type: PLUS, Literal: "+", Pos: 2},
		{Type: STAR, Literal: "*", Pos: 3},
		{Type: SLASH, Literal: "/", Pos: 5},
		{Type: MINUS, Literal: "-", Pos: 6},
		{Type: DOUBLE_STAR, Literal: "**", Pos: 7},
		{Type: DOUBLE_SLASH, Literal: "//", Pos: 9},
		{Type: NUMBER, Literal: "12.32", Pos: 12},
		{Type: PERCENT, Literal: "%", Pos: 17},
		{Type: ILLEGAL, Literal: "#", Pos: 18},
		{Type: EOF, Literal: "", Pos: 19},
	}

	lexer := NewLexer(testedString)

	for i, expectedToken := range expectedTokens {
		actualToken := lexer.GetNextToken()

		if actualToken != expectedToken {
			t.Fatalf("[%d] expected=%#v, got=%#v", i, expectedToken, actualToken)
		}
	}
}
