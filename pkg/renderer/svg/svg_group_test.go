package svg

import (
	"errors"
	"testing"
)

func TestMatrixFromGAttributes(t *testing.T) {
	tests := []struct {
		attribute string
		expected  matrix
	}{
		{
			attribute: "matrix(1,0,0,0,1,0)",
			expected:  newMatrix(1, 0, 0, 0, 1, 0, 0, 0, 1),
		},
		{
			attribute: "matrix(3.04349,0,0,3.04349,54.9563,54.9563)",
			expected:  newMatrix(3.04349, 0, 0, 3.04349, 54.9563, 54.9563, 0, 0, 1),
		},
	}

	for i, tt := range tests {
		result, err := matrixFromGAttributes(tt.attribute)

		if err != nil {
			t.Fatalf("[%d] should not failed but got %v", i, err)
		}
		err = testMatrixAreEquales(result, tt.expected)
		if err != nil {
			t.Fatalf("[%d] wrong result : %v", i, err)
		}
	}

}

func TestMatrixFromGAttributesFailing(t *testing.T) {
	tests := []struct {
		attribute string
		expected  error
	}{
		{
			attribute: "lol",
			expected:  errCannotParseMainTransformAttr,
		},
		{
			attribute: "azertyu",
			expected:  errCannotParseMainTransformAttr,
		},
		{
			attribute: "matrix",
			expected:  errCannotParseMainTransformAttr,
		},
		{
			attribute: "matrix()",
			expected:  errCannotParseMainTransformAttr,
		},
		{
			attribute: "matrix(1,2,3,4,5,)",
			expected:  errCannotParseMainTransformAttr,
		},
	}

	for i, tt := range tests {
		_, err := matrixFromGAttributes(tt.attribute)

		if !errors.Is(err, tt.expected) {
			t.Fatalf("[%d] should failed with %v but got %v", i, tt.expected, err)
		}
	}

}
