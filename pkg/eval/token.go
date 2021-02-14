package eval

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Pos     int
}

const (
	ILLEGAL      TokenType = "ILLEGAL"
	EOF                    = "EOF"
	NUMBER                 = "NUMBER"
	LEFT_PAR               = "LEFT_PARENTHESIS"
	RIGHT_PAR              = "RIGHT_PARENTHESIS"
	PLUS                   = "PLUS"
	MINUS                  = "MINUS"
	STAR                   = "STAR"
	SLASH                  = "SLASH"
	DOUBLE_STAR            = "DOUBLE_STAR"
	DOUBLE_SLASH           = "DOUBLE_SLASH"
)
