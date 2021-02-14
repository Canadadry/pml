package eval

import (
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		in       string
		expected float64
	}{
		{"1+1", 2.0},
		{"1+(2**2)", 5.0},
		{"1+(2**2)//3", 2.0},
	}

	for _, tt := range tests {
		result, err := Eval(tt.in)
		if err != nil {
			t.Fatalf("[%s] failed %v", tt.in, err)
		}
		if result != tt.expected {
			t.Fatalf("[%s] expected=%#v, got=%#v", tt.in, result, tt.expected)
		}

	}
}
