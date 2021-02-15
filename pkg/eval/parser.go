package eval

import (
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
)

var precedences = map[TokenType]int{
	PLUS:         SUM,
	MINUS:        SUM,
	PERCENT:      SUM,
	SLASH:        PRODUCT,
	DOUBLE_SLASH: PRODUCT,
	STAR:         PRODUCT,
	DOUBLE_STAR:  PRODUCT,
}

var (
	ErrBadToken         = fmt.Errorf("Bad token type encounter")
	ErrNoPrefixFunction = fmt.Errorf("Cannot find function to parse prefix")
	ErrNoInFixFunction  = fmt.Errorf("Cannot find function to parse infix")
)

type (
	prefixParseFn func() (Node, error)
	infixParseFn  func(Node) (Node, error)
)

type Parser struct {
	lexer   *Lexer
	current Token
	next    Token

	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

func NewParser(lexer *Lexer) *Parser {
	parser := &Parser{
		lexer:   lexer,
		current: Token{},
		next:    Token{},
	}

	parser.prefixParseFns = map[TokenType]prefixParseFn{
		NUMBER:   parser.parseFloatLiteral,
		MINUS:    parser.parsePrefixExpression,
		LEFT_PAR: parser.parseGroupedExpression,
	}

	parser.infixParseFns = map[TokenType]infixParseFn{
		PERCENT:      parser.parseInfixExpression,
		PLUS:         parser.parseInfixExpression,
		MINUS:        parser.parseInfixExpression,
		SLASH:        parser.parseInfixExpression,
		DOUBLE_SLASH: parser.parseInfixExpression,
		STAR:         parser.parseInfixExpression,
		DOUBLE_STAR:  parser.parseInfixExpression,
	}

	parser.moveToNextToken()
	parser.moveToNextToken()

	return parser
}

func (parser *Parser) moveToNextToken() {
	parser.current = parser.next
	parser.next = parser.lexer.GetNextToken()
}

func (parser *Parser) moveToNextTokenIfTypeIs(t TokenType) error {
	if !parser.isNextTokenA(t) {
		return fmt.Errorf("%w : expected %s, got %s", ErrBadToken, t, parser.next.Type)
	}
	parser.moveToNextToken()
	return nil
}

func (parser *Parser) isCurrentTokenA(t TokenType) bool {
	return parser.current.Type == t
}

func (parser *Parser) isNextTokenA(t TokenType) bool {
	return parser.next.Type == t
}

func (parser *Parser) getCurrentPrecedence() int {
	if p, ok := precedences[parser.current.Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) getNextPrecedence() int {
	if p, ok := precedences[parser.next.Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) ParseExpression(precedence int) (Node, error) {
	prefix := parser.prefixParseFns[parser.current.Type]
	if prefix == nil {
		return nil, fmt.Errorf("%w : %s", ErrNoPrefixFunction, parser.current.Type)
	}

	leftExp, err := prefix()
	if err != nil {
		return nil, err
	}

	for precedence < parser.getNextPrecedence() {
		infix := parser.infixParseFns[parser.next.Type]
		if infix == nil {
			return nil, fmt.Errorf("%w : %s", ErrNoInFixFunction, parser.next.Type)
		}
		parser.moveToNextToken()
		var err error
		leftExp, err = infix(leftExp)
		if err != nil {
			return nil, err
		}
	}
	return leftExp, nil
}

func (parser *Parser) parsePrefixExpression() (Node, error) {
	op, ok := PrefixFuncMap[parser.current.Type]
	if !ok {
		return nil, fmt.Errorf("No prefix function for %s", parser.current.Literal)
	}
	parser.moveToNextToken()
	value, err := parser.ParseExpression(PREFIX)
	return &Prefix{
		value:     value,
		operation: op,
	}, err
}

func (parser *Parser) parseInfixExpression(left Node) (Node, error) {
	op, ok := InfixFuncMap[parser.current.Type]
	if !ok {
		return nil, fmt.Errorf("No infix function for %s", parser.current.Literal)
	}
	precedence := parser.getCurrentPrecedence()
	parser.moveToNextToken()
	right, err := parser.ParseExpression(precedence)
	return &Infix{
		operation: op,
		left:      left,
		right:     right,
	}, err
}

func (parser *Parser) parseGroupedExpression() (Node, error) {
	parser.moveToNextToken()
	exp, err := parser.ParseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	if err := parser.moveToNextTokenIfTypeIs(RIGHT_PAR); err != nil {
		return nil, err
	}
	return exp, nil
}

func (parser *Parser) parseFloatLiteral() (Node, error) {
	value, err := strconv.ParseFloat(parser.current.Literal, 64)
	return Value(value), err
}
