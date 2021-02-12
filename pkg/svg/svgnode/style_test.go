package svgnode

import (
	"image/color"
	"testing"
)

func TestParseColorParam(t *testing.T) {
	tests := []struct {
		in       string
		expected color.RGBA
	}{
		{"blanchedalmond", color.RGBA{255, 235, 205, 0}},
		{"darkgoldenrod", color.RGBA{184, 134, 11, 0}},
		{"rgb(255,0,0)", color.RGBA{255, 0, 0, 0}},
		{"#ff0000", color.RGBA{255, 0, 0, 0}},
	}

	for _, tt := range tests {
		result, err := parseColorParam(tt.in)
		if err != nil {
			t.Fatalf("[%s] failed  %v", tt.in, err)
		}
		if result != tt.expected {
			t.Fatalf("[%s] exp %#v got %#v", tt.in, tt.expected, result)
		}
	}
}
