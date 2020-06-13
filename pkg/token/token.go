package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL     TokenType = "ILLEGAL"
	EOF                   = "EOF"
	IDENTIFIER            = "IDENTIFIER"
	FLOAT                 = "FLOAT"
	STRING                = "STRING"
	COLOR                 = "COLOR"
	DOTS                  = "DOTS"
	LEFT_BRACE            = "LEFT_BRACE"
	RIGHT_BRACE           = "RIGHT_BRACE"
)
