package parser

import (
	"fmt"
	"pml/ast"
	"pml/token"
)

func (p *parser) parseItem() (*ast.Item, error) {
	if !p.isCurrentTokenA(token.IDENTIFIER) {
		return nil, fmt.Errorf("%w : not an IDENTIFIER", errNextTokenIsNotTheExpectedOne)
	}

	item := &ast.Item{
		TokenType: p.current,
	}

	err := p.goToNextTokenIfIsA(token.LEFT_BRACE)
	if err != nil {
		return nil, err
	}
	err = p.goToNextTokenIfIsA(token.RIGHT_BRACE)
	if err != nil {
		return nil, err
	}

	return item, nil
}
