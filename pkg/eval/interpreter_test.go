package eval

import (
	"testing"
)

func TestNodeEval(t *testing.T) {
	tests := []struct {
		tree     Node
		expected float64
	}{
		{Value(1.0), 1.0},
		{Prefix{Value(1.0), Neg}, -1.0},
		{Infix{Value(1.0), Value(2.0), Add}, 3.0},
		{Infix{Value(1.0), Value(1.0), Sub}, 0.0},
		{Infix{Value(1.0), Value(1.0), Mul}, 1.0},
		{Infix{Value(1.0), Value(1.0), Div}, 1.0},
		{Infix{Value(3.0), Value(2.0), IntDiv}, 1.0},
		{Infix{Prefix{Value(1.0), Neg}, Value(2.0), Mul}, -2.0},
	}

	for i, tt := range tests {
		result := tt.tree.Eval()
		if result != tt.expected {
			t.Fatalf("[%d] expected=%#v, got=%#v", i, tt.expected, result)
		}

	}
}
