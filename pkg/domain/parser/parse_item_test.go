package parser

import (
	"errors"
	"github.com/canadadry/pml/pkg/domain/ast"
	"github.com/canadadry/pml/pkg/domain/lexer"
	"github.com/canadadry/pml/pkg/domain/token"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		program  string
		expected *ast.Item
	}{
		{
			"Doc{}",
			&ast.Item{TokenType: tokenIdentifier("Doc")},
		},
		{
			"Doc{prop:prop}",
			&ast.Item{
				TokenType: tokenIdentifier("Doc"),
				Properties: map[string]ast.Expression{
					"prop": &ast.Value{
						PmlToken: tokenIdentifier("prop"),
					},
				},
			},
		},
		{
			"Doc{prop:1}",
			&ast.Item{
				TokenType: tokenIdentifier("Doc"),
				Properties: map[string]ast.Expression{
					"prop": &ast.Value{
						PmlToken: tokenFloat("1"),
					},
				},
			},
		},
		{
			"Doc{prop:1.1}",
			&ast.Item{
				TokenType: tokenIdentifier("Doc"),
				Properties: map[string]ast.Expression{
					"prop": &ast.Value{
						PmlToken: tokenFloat("1.1"),
					},
				},
			},
		},
		{
			"Doc{prop:\"test\"}",
			&ast.Item{
				TokenType: tokenIdentifier("Doc"),
				Properties: map[string]ast.Expression{
					"prop": &ast.Value{
						PmlToken: tokenString("test"),
					},
				},
			},
		},
		{
			"Doc{prop:#123123}",
			&ast.Item{
				TokenType: tokenIdentifier("Doc"),
				Properties: map[string]ast.Expression{
					"prop": &ast.Value{
						PmlToken: tokenColor("123123"),
					},
				},
			},
		},
		{
			"Doc{prop1:prop2 prop3:prop4}",
			&ast.Item{
				TokenType: tokenIdentifier("Doc"),
				Properties: map[string]ast.Expression{
					"prop1": &ast.Value{
						PmlToken: tokenIdentifier("prop2"),
					},
					"prop3": &ast.Value{
						PmlToken: tokenIdentifier("prop4"),
					},
				},
			},
		},
		{
			"Doc{Page{}}",
			&ast.Item{
				TokenType: tokenIdentifier("Doc"),
				Children: []ast.Item{
					{TokenType: tokenIdentifier("Page")},
				},
			},
		},
	}

	for i, tt := range tests {
		l := lexer.New(tt.program)
		parser := New(l)

		result, err := parser.Parse()
		if err != nil {
			t.Fatalf("[%d] Failed with error : %v", i, err)
		}
		testItem(t, i, result, tt.expected)
	}
}

func TestParserError(t *testing.T) {
	tests := []struct {
		program  string
		expected error
	}{
		{"{}", errNextTokenIsNotTheExpectedOne},
		{"a}", errNextTokenIsNotTheExpectedOne},
		{"a{a", errInvalidIdentifier},
		{"a{a:b a:b}", errPropertyDefinedTwice},
		{"a{a:b b:b :", errNextTokenIsNotTheExpectedOne},
		{"a{a::}", errNotAValueType},
		{"a{a{", errNextTokenIsNotTheExpectedOne},
	}

	for i, tt := range tests {
		l := lexer.New(tt.program)
		parser := New(l)

		_, err := parser.Parse()
		if !errors.Is(err, tt.expected) {
			t.Fatalf("[%d] error was not the one expected got %s, exp %s", i, err, tt.expected)
		}
	}
}

func testItem(t *testing.T, index int, actual *ast.Item, expected *ast.Item) {
	if actual.TokenType.Type != expected.TokenType.Type {
		t.Fatalf("[%d] Wrong item type got %s expected %s", index, actual.TokenType.Type, expected.TokenType.Type)
	}
	if actual.TokenType.Literal != expected.TokenType.Literal {
		t.Fatalf("[%d] Wrong item literal got %s expected %s", index, actual.TokenType.Literal, expected.TokenType.Literal)
	}
	if len(actual.Properties) != len(expected.Properties) {
		t.Fatalf(
			"[%d] Wrong number of Properties for item %s got %d expected %d",
			index,
			actual.TokenType.Literal,
			len(actual.Properties),
			len(expected.Properties),
		)
	}

	for k := range expected.Properties {
		testProperty(t, index, k, actual, expected)
	}

	if len(actual.Children) != len(expected.Children) {
		t.Fatalf(
			"[%d] Wrong number of Children for item %s got %d expected %d",
			index,
			actual.TokenType.Literal,
			len(actual.Children),
			len(expected.Children),
		)
	}

	for i := 0; i < len(expected.Children); i++ {
		testItem(t, index, &actual.Children[i], &expected.Children[i])
	}
}

func testProperty(t *testing.T, index int, property string, actual *ast.Item, expected *ast.Item) {
	expectedValue := expected.Properties[property]
	actualValue, ok := actual.Properties[property]
	if !ok {
		t.Fatalf(
			"[%d] on item %s missing property %s",
			index,
			actual.TokenType.Literal,
			property,
		)
	}
	if actualValue.Token().Type != expectedValue.Token().Type {
		t.Fatalf(
			"[%d] Wrong property typeon %s got %s expected %s",
			index,
			property,
			actualValue.Token().Type,
			expectedValue.Token().Type,
		)
	}
	if actualValue.Token().Literal != expectedValue.Token().Literal {
		t.Fatalf(
			"[%d] Wrong property literal %s got %s expected %s",
			index,
			property,
			actualValue.Token().Literal,
			expectedValue.Token().Literal,
		)
	}
}

func tokenIdentifier(literal string) token.Token {
	return token.Token{
		Type:    token.IDENTIFIER,
		Literal: literal,
	}
}

func tokenFloat(literal string) token.Token {
	return token.Token{
		Type:    token.FLOAT,
		Literal: literal,
	}
}

func tokenString(literal string) token.Token {
	return token.Token{
		Type:    token.STRING,
		Literal: literal,
	}
}

func tokenColor(literal string) token.Token {
	return token.Token{
		Type:    token.COLOR,
		Literal: literal,
	}
}
