package parser

import (
	"fmt"
	"github.com/canadadry/pml/language/ast"
	"github.com/canadadry/pml/language/token"
)

func (p *parser) parseItem() (*ast.Item, error) {
	if !p.isCurrentTokenA(token.IDENTIFIER) {
		return nil, fmt.Errorf("%w : not an IDENTIFIER", errNextTokenIsNotTheExpectedOne)
	}

	item := &ast.Item{
		TokenType:  p.current,
		Properties: map[string]ast.Expression{},
		Children:   []ast.Item{},
	}

	err := p.goToNextTokenIfIsA(token.LEFT_BRACE)
	if err != nil {
		return nil, err
	}

	for p.isNextTokenA(token.IDENTIFIER) {
		p.goToNextToken()

		if p.isNextTokenA(token.DOTS) {
			if err := p.parsePropertyInItem(item); err != nil {
				return nil, err
			}
		} else if p.isNextTokenA(token.LEFT_BRACE) {
			child, err := p.parseItem()
			if err != nil {
				return nil, err
			}
			item.Children = append(item.Children, *child)
		} else {
			return nil, fmt.Errorf("In %s, token.IDENTIFIER %s is not a property : %w", item.TokenType.Literal, p.current.Literal, errInvalidIdentifier)
		}
	}

	err = p.goToNextTokenIfIsA(token.RIGHT_BRACE)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (p *parser) parsePropertyInItem(item *ast.Item) error {
	propertyName := p.current.Literal
	p.goToNextToken()
	p.goToNextToken()

	_, ok := item.Properties[propertyName]
	if ok {
		return fmt.Errorf("In %s, property %s : %w", item.TokenType.Literal, propertyName, errPropertyDefinedTwice)
	}
	value, err := p.parseValue()
	if err != nil {
		return fmt.Errorf("In %s, property %s : %w", item.TokenType.Literal, propertyName, err)
	}
	item.Properties[propertyName] = value

	return nil
}

func (p *parser) parseValue() (*ast.Value, error) {
	v := &ast.Value{
		PmlToken: p.current,
	}

	switch p.current.Type {
	case token.IDENTIFIER:
		return v, nil
	case token.FLOAT:
		return v, nil
	case token.STRING:
		return v, nil
	case token.COLOR:
		return v, nil
	}

	return nil, fmt.Errorf("%w : got %s", errNotAValueType, string(p.current.Type))
}
