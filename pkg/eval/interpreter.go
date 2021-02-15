package eval

import (
	"math"
)

var (
	InfixFuncMap = map[TokenType]InfixFunc{
		PLUS:         Add,
		MINUS:        Sub,
		STAR:         Mul,
		SLASH:        Div,
		PERCENT:      math.Mod,
		DOUBLE_STAR:  math.Pow,
		DOUBLE_SLASH: IntDiv,
	}
	PrefixFuncMap = map[TokenType]PrefixFunc{
		MINUS: Neg,
	}
)

type Value float64

func (v Value) Eval() float64 { return float64(v) }

type InfixFunc func(l, r float64) float64

type Infix struct {
	left      Node
	right     Node
	operation InfixFunc
}

func (in Infix) Eval() float64 { return in.operation(in.left.Eval(), in.right.Eval()) }

func Add(l, r float64) float64    { return l + r }
func Sub(l, r float64) float64    { return l - r }
func Mul(l, r float64) float64    { return l * r }
func Div(l, r float64) float64    { return l / r }
func IntDiv(l, r float64) float64 { return math.Floor(l / r) }

type PrefixFunc func(v float64) float64

type Prefix struct {
	value     Node
	operation PrefixFunc
}

func (p Prefix) Eval() float64 { return p.operation(p.value.Eval()) }

func Neg(v float64) float64 { return -v }
