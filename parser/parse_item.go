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
		TokenType:  p.current,
		Properties: map[string]ast.Expression{},
	}

	err := p.goToNextTokenIfIsA(token.LEFT_BRACE)
	if err != nil {
		return nil, err
	}

	for p.isNextTokenA(token.IDENTIFIER) {
		p.goToNextToken()
		propertyName := p.current.Literal
		value, err := p.parseProperty()
		if err != nil {
			return nil, fmt.Errorf("In %s, parsing property %s : %w", item.TokenType.Literal, propertyName, err)
		}
		_, ok := item.Properties[propertyName]
		if ok {
			return nil, fmt.Errorf("In %s, property %s : %w", item.TokenType.Literal, propertyName, errPropertyDefinedTwice)
		}
		item.Properties[propertyName] = value
	}

	err = p.goToNextTokenIfIsA(token.RIGHT_BRACE)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (p *parser) parseProperty() (ast.Expression, error) {
	err := p.goToNextTokenIfIsA(token.DOTS)
	if err != nil {
		return nil, err
	}

	p.goToNextToken()

	return &ast.Value{
		PmlToken: p.current,
	}, nil
}
