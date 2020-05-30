package ast

import (
	"pml/token"
)

type Item struct {
	TokenType  token.Token
	Properties map[string]Expression
	Child      []Item
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
