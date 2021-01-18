package svgpath

type Lexer struct {
	source  string
	current int
	read    int
	ch      byte
}

func newLexer(source string) *Lexer {
	l := &Lexer{
		source:  source,
		current: 0,
		read:    0,
		ch:      0,
	}
	l.readChar()
	return l
}

func (lexer *Lexer) getNextToken() Token {
	for isWhiteSpace(lexer.ch) {
		lexer.readChar()
	}
	tok := Token{
		Type:    "",
		Literal: string(lexer.ch),
	}

	switch lexer.ch {
	case ',':
		tok.Type = COMA
		lexer.readChar()
	case 0:
		tok.Type = EOF
		tok.Literal = ""
		lexer.readChar()
	case '-':
		lexer.readChar()
		literal, tokenType := lexer.readNumeric()
		tok.Literal = "-" + literal
		tok.Type = tokenType
	default:
		switch {
		case isLetter(lexer.ch):
			tok.Type = COMMAND
			tok.Literal = string(lexer.ch)
			lexer.readChar()
		case isNumeric(lexer.ch):
			literal, tokenType := lexer.readNumeric()
			tok.Literal = literal
			tok.Type = tokenType
		default:
			tok.Type = ILLEGAL
			lexer.readChar()
		}
	}

	return tok
}

func (lexer *Lexer) readChar() {
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

func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (lexer *Lexer) readNumeric() (string, TokenType) {
	start := lexer.current
	for isNumeric(lexer.ch) {
		lexer.readChar()
	}

	if isDot(lexer.ch) {
		lexer.readChar()
		for isNumeric(lexer.ch) {
			lexer.readChar()
		}
	}

	return lexer.source[start:lexer.current], NUMBER
}
func isDot(ch byte) bool {
	return ch == '.'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}
