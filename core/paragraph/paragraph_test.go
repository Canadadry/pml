package paragraph

import (
	"testing"
)

type fakeSizer struct{}

func (fs fakeSizer) GetStringWidth(str string, fontName string, font float64) float64 {
	return float64(len(str))
}

func TestBlocsToWords(t *testing.T) {
	tests := map[string]struct {
		in  []string
		out []Word
	}{
		"simple text": {
			in: []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce sagittis tincidunt porttitor. Donec"},
			out: []Word{
				Word{Text: "Lorem"},
				Word{Text: "ipsum"},
				Word{Text: "dolor"},
				Word{Text: "sit"},
				Word{Text: "amet,"},
				Word{Text: "consectetur"},
				Word{Text: "adipiscing"},
				Word{Text: "elit."},
				Word{Text: "Fusce"},
				Word{Text: "sagittis"},
				Word{Text: "tincidunt"},
				Word{Text: "porttitor."},
				Word{Text: "Donec"},
			},
		},
		"text with one line break": {
			in: []string{"Lorem\nipsum"},
			out: []Word{
				Word{Text: "Lorem"},
				Word{Text: "\n"},
				Word{Text: "ipsum"},
			},
		},
		"text with two line break": {
			in: []string{"Lorem\n\nipsum"},
			out: []Word{
				Word{Text: "Lorem"},
				Word{Text: "\n"},
				Word{Text: "\n"},
				Word{Text: "ipsum"},
			},
		},
		"text on two bloc": {
			in: []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "Fusce sagittis tincidunt porttitor. Donec"},
			out: []Word{
				Word{Text: "Lorem"},
				Word{Text: "ipsum"},
				Word{Text: "dolor"},
				Word{Text: "sit"},
				Word{Text: "amet,"},
				Word{Text: "consectetur"},
				Word{Text: "adipiscing"},
				Word{Text: "elit."},
				Word{Text: "Fusce"},
				Word{Text: "sagittis"},
				Word{Text: "tincidunt"},
				Word{Text: "porttitor."},
				Word{Text: "Donec"},
			},
		},
		"text with two line break on two bloc": {
			in: []string{"Lorem\n", "\nipsum"},
			out: []Word{
				Word{Text: "Lorem"},
				Word{Text: "\n"},
				Word{Text: "\n"},
				Word{Text: "ipsum"},
			},
		},
	}

	for title, tt := range tests {
		b := []Block{}
		for _, s := range tt.in {
			b = append(b, Block{Text: s})
		}

		result := blocsToWords(b)
		if len(tt.out) != len(result) {
			t.Fatalf("[%s] result len got %d exp %d\n got %v\n exp %v", title, len(result), len(tt.out), result, tt.out)
		}

		for j := range tt.out {
			if tt.out[j].Text != result[j].Text {
				t.Fatalf("[%s] result at %d got '%s' exp '%s'", title, j, result[j].Text, tt.out[j].Text)
			}
		}
	}
}
