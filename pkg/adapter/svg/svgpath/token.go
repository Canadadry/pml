package svgpath

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"
	NUMBER            = "NUMBER"
	COMMAND           = "COMMAND"
	COMA              = "COMA"
)
