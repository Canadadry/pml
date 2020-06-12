package lexer

import (
	"github.com/canadadry/pml/pkg/token"
)

type Lexer struct {
	source  string
	current int
	read    int
	ch      byte
	line    int
	column  int
}

func New(source string) *Lexer {
	l := &Lexer{
		source:  source,
		current: 0,
		read:    0,
		ch:      0,
		line:    1,
		column:  0,
	}
	l.readChar()
	return l
}

func (lexer *Lexer) GetNextToken() token.Token {
	for isWhiteSpace(lexer.ch) {
		lexer.readChar()
	}
	tok := token.Token{
		Type:    "",
		Literal: string(lexer.ch),
		Line:    lexer.line,
		Column:  lexer.column,
	}

	switch lexer.ch {
	case '{':
		tok.Type = token.LEFT_BRACE
		lexer.readChar()
	case '}':
		tok.Type = token.RIGHT_BRACE
		lexer.readChar()
	case ':':
		tok.Type = token.DOTS
		lexer.readChar()
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
		lexer.readChar()
	case '"':
		tok.Type = token.STRING
		lexer.readChar()
		tok.Literal = lexer.readString()
	case '#':
		lexer.readChar()
		tok.Literal, tok.Type = lexer.readColor()
	default:
		switch {
		case isLetter(lexer.ch):
			tok.Type = token.IDENTIFIER
			tok.Literal = lexer.readIdentifier()
		case isNumeric(lexer.ch):
			literal, tokenType := lexer.readNumeric()
			tok.Literal = literal
			tok.Type = tokenType
		default:
			tok.Type = token.ILLEGAL
			lexer.readChar()
		}
	}

	return tok
}

func (lexer *Lexer) readChar() {

	lexer.column += 1

	if lexer.ch == '\n' {
		lexer.column = 1
		lexer.line += 1
	}

	lexer.ch = 0

	if lexer.read < len(lexer.source) {
		lexer.ch = lexer.source[lexer.read]
	}

	lexer.current = lexer.read
	lexer.read += 1

}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (lexer *Lexer) readIdentifier() string {
	start := lexer.current
	for isLetter(lexer.ch) || isNumeric(lexer.ch) {
		lexer.readChar()
	}
	return lexer.source[start:lexer.current]
}

func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (lexer *Lexer) readNumeric() (string, token.TokenType) {
	start := lexer.current
	for isNumeric(lexer.ch) {
		lexer.readChar()
	}

	switch true {
	case isLetter(lexer.ch):
		{
			for isLetter(lexer.ch) || isNumeric(lexer.ch) {
				lexer.readChar()
			}
			return lexer.source[start:lexer.current], token.ILLEGAL
		}
	case isDot(lexer.ch):
		{
			lexer.readChar()
			for isNumeric(lexer.ch) {
				lexer.readChar()
			}
			return lexer.source[start:lexer.current], token.FLOAT
		}
	}

	return lexer.source[start:lexer.current], token.INTEGER
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func (lexer *Lexer) readString() string {
	start := lexer.current
	for lexer.ch != '"' && lexer.ch != 0 {
		lexer.readChar()
	}
	defer lexer.readChar()
	return lexer.source[start:lexer.current]
}

func isDot(ch byte) bool {
	return ch == '.'
}

func isHexaChar(ch byte) bool {
	return ch >= '0' && ch <= '9' || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}

func (lexer *Lexer) readColor() (string, token.TokenType) {
	start := lexer.current
	for isHexaChar(lexer.ch) {
		lexer.readChar()
	}
	length := lexer.current - start
	if length != 6 && length != 8 {
		return lexer.source[start:lexer.current], token.ILLEGAL
	}
	return lexer.source[start:lexer.current], token.COLOR
}
