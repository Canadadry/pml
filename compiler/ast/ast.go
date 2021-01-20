package ast

import (
	"errors"
	"fmt"
	"github.com/canadadry/pml/compiler/token"
	"image/color"
	"strconv"
)

var (
	ErrInvalidTypeForProperty     = errors.New("invalidTypeForProperty")
	errPropertyDefinitionNotFound = errors.New("errPropertyDefinitionNotFound")
)

type Item struct {
	TokenType  token.Token
	Properties map[string]Expression
	Children   []Item
}

type Expression interface {
	Token() token.Token
}

type Value struct {
	PmlToken token.Token
}

func (v *Value) Token() token.Token {
	return v.PmlToken
}

func (i *Item) GetPropertyAsColorWithDefault(name string, defaultValue color.RGBA) (color.RGBA, error) {

	v, ok := i.Properties[name]
	if !ok {
		return defaultValue, nil
	}

	if v.Token().Type != token.COLOR {
		return defaultValue, ErrInvalidTypeForProperty
	}
	s := v.Token().Literal

	c := color.RGBA{
		A: 0xff,
	}

	var err error = nil

	switch len(s) {
	case 6:
		_, err = fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
	default:
		err = fmt.Errorf("invalid length, must be 6")

	}
	return c, err
}
func (i *Item) GetPropertyAsFloatWithDefault(name string, defaultValue float64) (float64, error) {
	v, ok := i.Properties[name]
	if !ok {
		return defaultValue, nil
	}

	if v.Token().Type != token.FLOAT {
		return defaultValue, ErrInvalidTypeForProperty
	}

	value, err := strconv.ParseFloat(v.Token().Literal, 64)
	if err != nil {
		return defaultValue, err
	}
	return value, nil
}
func (i *Item) GetPropertyAsStringWithDefault(name string, defaultValue string) (string, error) {

	v, ok := i.Properties[name]
	if !ok {
		return defaultValue, nil
	}

	if v.Token().Type != token.STRING {
		return defaultValue, ErrInvalidTypeForProperty
	}
	return v.Token().Literal, nil
}
func (i *Item) GetPropertyAsIdentifierWithDefault(name string, defaultValue string) (string, error) {

	v, ok := i.Properties[name]
	if !ok {
		return defaultValue, nil
	}

	if v.Token().Type != token.IDENTIFIER {
		return defaultValue, ErrInvalidTypeForProperty
	}
	return v.Token().Literal, nil
}
