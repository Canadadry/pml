package svgpath

import (
	"errors"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		program  string
		expected []Command
	}{
		{
			"m 234.43804,111.69821 c 50.21866,26.50627 126.75595,-3.87395 151.46369,-35.941621",
			[]Command{
				{'m', []Point{{234.43804, 111.69821}}},
				{'c', []Point{{50.21866, 26.50627}, {126.75595, -3.87395}, {151.46369, -35.941621}}},
			},
		},
		{
			"m 10,12h 234.43804",
			[]Command{
				{'m', []Point{{10, 12}}},
				{'h', []Point{{234.43804, 0}}},
			},
		},
	}

	for i, tt := range tests {
		l := newLexer(tt.program)
		parser := newParser(l)

		result, err := parser.parse()
		if err != nil {
			t.Fatalf("[%d] Failed with error : %v", i, err)
		}
		if len(result) != len(tt.expected) {
			t.Errorf("expected commands %#v", tt.expected)
			t.Errorf("  result commands %#v", result)
			t.Fatalf("[%d] commands wrong size got %d exp %d", i, len(result), len(tt.expected))
		}
		for j := range tt.expected {
			if tt.expected[j].ToString() != result[j].ToString() {
				t.Fatalf("[%d] command %d got %s exp %s", i, j, result[j].ToString(), tt.expected[j].ToString())
			}
		}
	}
}

func TestParserErrors(t *testing.T) {
	tests := []struct {
		program  string
		expected error
	}{
		{
			"234.43804",
			ErrExpectedCommandToken,
		},
		{
			"m234.43804,",
			ErrExpectedFloatToken,
		},
		{
			"m234.43804,12 10m",
			ErrExpectedComaToken,
		}}

	for i, tt := range tests {
		l := newLexer(tt.program)
		parser := newParser(l)

		_, err := parser.parse()
		if !errors.Is(err, tt.expected) {
			t.Fatalf("[%d] error was not the one expected got %s, exp %s", i, err, tt.expected)
		}
	}
}
