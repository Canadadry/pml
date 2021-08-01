package paragraph

import (
	"fmt"
	"testing"
)

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
		testWords(t, title, result, tt.out)
	}
}

func testWords(t *testing.T, title string, result, expected []Word) {
	if len(expected) != len(result) {
		t.Fatalf("[%s] len got %d exp %d", title, len(result), len(expected))
	}

	for j := range expected {
		if expected[j].Text != result[j].Text {
			t.Fatalf("[%s:%d] got '%s' exp '%s'", title, j, result[j].Text, expected[j].Text)
		}
	}
}

type fakeSizer struct{}

func (fs fakeSizer) GetStringWidth(str string, fontName string, font float64) float64 {
	// fmt.Println("'"+str+"'", len(str))
	return float64(len(str))
}

func TestFormat(t *testing.T) {
	tests := map[string]struct {
		in    string
		width float64
		out   []Line
	}{
		"one line text": {
			in:    "Lorem ipsum dolor sit amet, consectetur",
			width: 50,
			out: []Line{
				{
					Words: []Word{
						Word{Text: "Lorem", Width: 5},
						Word{Text: "ipsum", Width: 5},
						Word{Text: "dolor", Width: 5},
						Word{Text: "sit", Width: 3},
						Word{Text: "amet,", Width: 5},
						Word{Text: "consectetur", Width: 11},
					},
					StartingOffset: 0,
					MaxWidth:       50,
					Overflow:       false,
				},
			},
		},
		"two lines text with \\n": {
			in:    "Lorem ipsum dolor sit \n amet, consectetur",
			width: 50,
			out: []Line{
				{
					Words: []Word{
						Word{Text: "Lorem", Width: 5},
						Word{Text: "ipsum", Width: 5},
						Word{Text: "dolor", Width: 5},
						Word{Text: "sit", Width: 3},
					},
					StartingOffset: 0,
					MaxWidth:       50,
					Overflow:       false,
				},
				{
					Words: []Word{
						Word{Text: "amet,", Width: 5},
						Word{Text: "consectetur", Width: 11},
					},
					StartingOffset: 0,
					MaxWidth:       50,
					Overflow:       false,
				},
			},
		},
		"two line text with overflow": {
			in:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce sagittis tincidunt porttitor. Donec",
			width: 50,
			out: []Line{
				{
					Words: []Word{
						Word{Text: "Lorem", Width: 5},
						Word{Text: "ipsum", Width: 5},
						Word{Text: "dolor", Width: 5},
						Word{Text: "sit", Width: 3},
						Word{Text: "amet,", Width: 5},
						Word{Text: "consectetur", Width: 11},
						Word{Text: "adipiscing", Width: 10},
					},
					StartingOffset: 0,
					MaxWidth:       50,
					Overflow:       true,
				},
				{
					Words: []Word{
						Word{Text: "elit.", Width: 5},
						Word{Text: "Fusce", Width: 5},
						Word{Text: "sagittis", Width: 8},
						Word{Text: "tincidunt", Width: 9},
						Word{Text: "porttitor.", Width: 9},
						Word{Text: "Donec", Width: 5},
					},
					StartingOffset: 0,
					MaxWidth:       50,
					Overflow:       false,
				},
			},
		},
	}

	for title, tt := range tests {
		b := append([]Block{}, Block{Text: tt.in})

		result := Format(b, tt.width, fakeSizer{})

		if len(result) != len(tt.out) {
			t.Fatalf("[%s] result len got %d exp %d", title, len(result), len(tt.out))
		}
		for j := range tt.out {
			testLine(t, fmt.Sprintf("%s:line[%d]", title, j), result[j], tt.out[j])
		}
	}
}

func testLine(t *testing.T, title string, result, expected Line) {
	if result.StartingOffset != expected.StartingOffset {
		t.Fatalf("[%s] StartingOffset got '%v' exp '%v'", title, result.StartingOffset, expected.StartingOffset)
	}
	if result.MaxWidth != expected.MaxWidth {
		t.Fatalf("[%s] MaxWidth got '%v' exp '%v'", title, result.MaxWidth, expected.MaxWidth)
	}
	if result.Overflow != expected.Overflow {
		t.Fatalf("[%s] Overflow got '%v' exp '%v'", title, result.Overflow, expected.Overflow)
	}
	testWords(t, title+".words", result.Words, expected.Words)
}
