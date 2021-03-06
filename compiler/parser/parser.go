package parser

import (
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"github.com/canadadry/pml/compiler/lexer"
	"github.com/canadadry/pml/compiler/token"
)

type parser struct {
	current token.Token
	next    token.Token
	lexer   *lexer.Lexer
}

func New(l *lexer.Lexer) parser {
	return parser{
		current: l.GetNextToken(),
		next:    l.GetNextToken(),
		lexer:   l,
	}
}

func (p *parser) Parse() (*ast.Item, error) {
	return p.parseItem()
}

func (p *parser) goToNextToken() {
	p.current = p.next
	p.next = p.lexer.GetNextToken()
}

func (p *parser) isCurrentTokenA(t token.TokenType) bool {
	return p.current.Type == t
}
func (p *parser) isNextTokenA(t token.TokenType) bool {
	return p.next.Type == t
}

func (p *parser) goToNextTokenIfIsA(t token.TokenType) error {
	if !p.isNextTokenA(t) {
		return fmt.Errorf("%w got %s", errNextTokenIsNotTheExpectedOne, string(t))
	}
	p.goToNextToken()

	return nil
}
