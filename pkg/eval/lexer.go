package eval

type Lexer struct {
	source  string
	current int
	read    int
	ch      byte
}

func NewLexer(source string) *Lexer {
	l := &Lexer{
		source: source,
	}
	l.readChar()
	return l
}

func (lexer *Lexer) GetNextToken() Token {
	for isWhiteSpace(lexer.ch) {
		lexer.readChar()
	}
	tok := Token{
		Type:    "",
		Literal: string(lexer.ch),
		Pos:     lexer.current,
	}
	switch lexer.ch {
	case 0:
		tok.Type = EOF
		tok.Literal = ""
		lexer.readChar()
	case '(':
		tok.Type = LEFT_PAR
		lexer.readChar()
	case ')':
		tok.Type = RIGHT_PAR
		lexer.readChar()
	case '+':
		tok.Type = PLUS
		lexer.readChar()
	case '-':
		tok.Type = MINUS
		lexer.readChar()
	case '%':
		tok.Type = PERCENT
		lexer.readChar()
	case '*':
		if lexer.peekChar() == '*' {
			tok.Type = DOUBLE_STAR
			tok.Literal = "**"
			lexer.readChar()
			lexer.readChar()
		} else {
			tok.Type = STAR
			lexer.readChar()
		}
	case '/':
		if lexer.peekChar() == '/' {
			tok.Type = DOUBLE_SLASH
			tok.Literal = "//"
			lexer.readChar()
			lexer.readChar()
		} else {
			tok.Type = SLASH
			lexer.readChar()
		}
	default:
		if !isNumeric(lexer.ch) {
			tok.Type = ILLEGAL
			lexer.readChar()
		} else {
			tok.Type = NUMBER
			tok.Literal = lexer.readNumeric()
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

func (lexer *Lexer) peekChar() byte {
	var ch byte
	if lexer.read < len(lexer.source) {
		ch = lexer.source[lexer.read]
	}
	return ch
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9' || ch == '.'
}

func (lexer *Lexer) readNumeric() string {
	start := lexer.current
	for isNumeric(lexer.ch) {
		lexer.readChar()
	}
	return lexer.source[start:lexer.current]
}
