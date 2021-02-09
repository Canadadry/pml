package renderer

import (
	"fmt"
	"github.com/canadadry/pml/compiler/lexer"
	"github.com/canadadry/pml/compiler/parser"
	"testing"
)

func TestTextToWords(t *testing.T) {
	tests := []struct {
		in  string
		out []Word
	}{
		{
			in: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce sagittis tincidunt porttitor. Donec",
			out: []Word{
				Word{spaceWidth: 1, width: 5, text: "Lorem"},
				Word{spaceWidth: 1, width: 5, text: "ipsum"},
				Word{spaceWidth: 1, width: 5, text: "dolor"},
				Word{spaceWidth: 1, width: 3, text: "sit"},
				Word{spaceWidth: 1, width: 5, text: "amet,"},
				Word{spaceWidth: 1, width: 11, text: "consectetur"},
				Word{spaceWidth: 1, width: 10, text: "adipiscing"},
				Word{spaceWidth: 1, width: 5, text: "elit."},
				Word{spaceWidth: 1, width: 5, text: "Fusce"},
				Word{spaceWidth: 1, width: 8, text: "sagittis"},
				Word{spaceWidth: 1, width: 9, text: "tincidunt"},
				Word{spaceWidth: 1, width: 10, text: "porttitor."},
				Word{spaceWidth: 1, width: 5, text: "Donec"},
			},
		},
		{
			in: "Lorem\nipsum",
			out: []Word{
				Word{spaceWidth: 1, width: 5, text: "Lorem"},
				Word{text: "\n"},
				Word{spaceWidth: 1, width: 5, text: "ipsum"},
			},
		},
		{
			in: "Lorem\n\nipsum",
			out: []Word{
				Word{spaceWidth: 1, width: 5, text: "Lorem"},
				Word{text: "\n"},
				Word{text: "\n"},
				Word{spaceWidth: 1, width: 5, text: "ipsum"},
			},
		},
	}

	for i, tt := range tests {
		n := NodeText{text: tt.in}
		pdf := &drawer{t: t, charSize: 1}

		result := textToWords(pdf, n)
		if len(tt.out) != len(result) {
			t.Fatalf("[%d] result len got %d exp %d\n got %v\n exp %v on : \n%s", i, len(result), len(tt.out), result, tt.out, tt.in)
		}

		for j := range tt.out {
			if tt.out[j] != result[j] {
				t.Fatalf("[%d] result at %d got '%#v' exp '%#v' on : \n%s", i, j, result[j], tt.out[j], tt.in)
			}
		}
	}
}

func TestWordsToLines(t *testing.T) {
	tests := []struct {
		in    []Word
		width float64
		out   []Line
	}{
		{
			in: []Word{
				Word{spaceWidth: 1, width: 5, text: "Lorem"},
				Word{spaceWidth: 1, width: 5, text: "ipsum"},
				Word{spaceWidth: 1, width: 5, text: "dolor"},
				Word{spaceWidth: 1, width: 3, text: "sit"},
				Word{spaceWidth: 1, width: 5, text: "amet,"},
				Word{spaceWidth: 1, width: 11, text: "consectetur"},
				Word{spaceWidth: 1, width: 10, text: "adipiscing"},
				Word{spaceWidth: 1, width: 5, text: "elit."},
				Word{spaceWidth: 1, width: 5, text: "Fusce"},
				Word{spaceWidth: 1, width: 8, text: "sagittis"},
				Word{spaceWidth: 1, width: 9, text: "tincidunt"},
				Word{spaceWidth: 1, width: 10, text: "porttitor."},
				Word{spaceWidth: 1, width: 5, text: "Donec"},
				Word{spaceWidth: 1, width: 5, text: "marro"},
			},
			width: 50,
			out: []Line{
				Line{
					words: []Word{
						Word{spaceWidth: 1, width: 5, text: "Lorem"},
						Word{spaceWidth: 1, width: 5, text: "ipsum"},
						Word{spaceWidth: 1, width: 5, text: "dolor"},
						Word{spaceWidth: 1, width: 3, text: "sit"},
						Word{spaceWidth: 1, width: 5, text: "amet,"},
						Word{spaceWidth: 1, width: 11, text: "consectetur"},
						Word{spaceWidth: 1, width: 10, text: "adipiscing"},
					},
				},
				Line{
					words: []Word{
						Word{spaceWidth: 1, width: 5, text: "elit."},
						Word{spaceWidth: 1, width: 5, text: "Fusce"},
						Word{spaceWidth: 1, width: 8, text: "sagittis"},
						Word{spaceWidth: 1, width: 9, text: "tincidunt"},
						Word{spaceWidth: 1, width: 10, text: "porttitor."},
						Word{spaceWidth: 1, width: 5, text: "Donec"},
					},
				},
				Line{
					words: []Word{
						Word{spaceWidth: 1, width: 5, text: "marro"},
					},
				},
			},
		},
		{
			in: []Word{
				Word{spaceWidth: 1, width: 5, text: "Lorem"},
				Word{spaceWidth: 1, width: 5, text: "ipsum"},
				Word{text: "\n"},
				Word{spaceWidth: 1, width: 5, text: "dolor"},
				Word{spaceWidth: 1, width: 3, text: "sit"},
			},
			width: 50,
			out: []Line{
				Line{
					words: []Word{
						Word{spaceWidth: 1, width: 5, text: "Lorem"},
						Word{spaceWidth: 1, width: 5, text: "ipsum"},
					},
				},
				Line{
					words: []Word{
						Word{spaceWidth: 1, width: 5, text: "dolor"},
						Word{spaceWidth: 1, width: 3, text: "sit"},
					},
				},
			},
		},
		{
			in: []Word{
				Word{spaceWidth: 1, width: 5, text: "Lorem"},
				Word{spaceWidth: 1, width: 5, text: "ipsum"},
				Word{text: "\n"},
				Word{text: "\n"},
				Word{spaceWidth: 1, width: 5, text: "dolor"},
				Word{spaceWidth: 1, width: 3, text: "sit"},
			},
			width: 50,
			out: []Line{
				Line{
					words: []Word{
						Word{spaceWidth: 1, width: 5, text: "Lorem"},
						Word{spaceWidth: 1, width: 5, text: "ipsum"},
					},
				},
				Line{
					words: []Word{},
				},
				Line{
					words: []Word{
						Word{spaceWidth: 1, width: 5, text: "dolor"},
						Word{spaceWidth: 1, width: 3, text: "sit"},
					},
				},
			},
		},
	}

	for i, tt := range tests {

		result := wordsToLines(tt.in, tt.width)
		if len(tt.out) != len(result) {
			t.Fatalf("[%d] result len got %d exp %d\n got %v\n exp %v ", i, len(result), len(tt.out), result, tt.out)
		}

		for j := range tt.out {
			if len(tt.out[j].words) != len(result[j].words) {
				t.Fatalf("[%d] result len got %d exp %d\n got %v\n exp %v ", i, len(result[j].words), len(tt.out[j].words), result[j].words, tt.out[j].words)
			}

			for k := range tt.out[j].words {
				if tt.out[j].words[k] != result[j].words[k] {
					t.Fatalf("[%d] result at %d:%d got '%#v' exp '%#v' ", i, j, k, result[j].words[k], tt.out[j].words[k])
				}
			}
		}
	}
}
func TestGetLineSize(t *testing.T) {
	tests := []struct {
		in    Line
		width float64
		out   LineSize
	}{
		{

			in: Line{
				words: []Word{
					Word{spaceWidth: 1, width: 5, text: "elit."},
					Word{spaceWidth: 1, width: 5, text: "Fusce"},
					Word{spaceWidth: 1, width: 8, text: "sagittis"},
					Word{spaceWidth: 1, width: 9, text: "tincidunt"},
					Word{spaceWidth: 1, width: 10, text: "porttitor."},
					Word{spaceWidth: 1, width: 5, text: "Donec"},
				},
			},
			width: 50,
			out: LineSize{
				spaceLeft: 3,
				wordCount: 6,
			},
		},
		{

			in: Line{
				words: []Word{
					Word{spaceWidth: 1, width: 5, text: "Lorem"},
					Word{spaceWidth: 1, width: 5, text: "ipsum"},
					Word{spaceWidth: 1, width: 5, text: "dolor"},
					Word{spaceWidth: 1, width: 3, text: "sit"},
					Word{spaceWidth: 1, width: 5, text: "amet,"},
					Word{spaceWidth: 1, width: 11, text: "consectetur"},
					Word{spaceWidth: 1, width: 10, text: "adipiscing"},
				},
			},
			width: 50,
			out: LineSize{
				spaceLeft: 0,
				wordCount: 7,
			},
		},
	}

	for i, tt := range tests {
		result := getLineSize(tt.in, tt.width)
		if tt.out != result {
			t.Fatalf("[%d] result got %v exp %v", i, result, tt.out)
		}
	}
}

//left (width 50)
// ****-****-****-****-****-****-****-****-****-****-
// Lorem ipsum dolor sit amet, consectetur adipiscing
// elit. Fusce sagittis tincidunt porttitor. Donec
// nec fringilla risus. Aliquam erat volutpat. Nunc
// sodales, orci nec efficitur
//
// right (width 50)
// ****-****-****-****-****-****-****-****-****-****-
// Lorem ipsum dolor sit amet, consectetur adipiscing
//    elit. Fusce sagittis tincidunt porttitor. Donec
//   nec fringilla risus. Aliquam erat volutpat. Nunc
//                        sodales, orci nec efficitur
//

func TestParagraphAlign(t *testing.T) {
	tests := []struct {
		in    string
		calls []string
	}{
		// {
		// 	in: `Document{
		// 		Page{
		// 			Paragraph{
		// 				width:50
		// 				align:left
		// 				lineHeight:1
		// 				Text{text:"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce sagittis tincidunt porttitor. Donec nec fringilla risus. Aliquam erat volutpat. Nunc sodales, orci nec efficitur"}
		// 			}
		// 		}
		// 	}`,
		// 	calls: []string{
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Lorem", 0, 0),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "ipsum", 6, 0),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "dolor", 12, 0),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "sit", 18, 0),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "amet,", 22, 0),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "consectetur", 28, 0),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "adipiscing", 40, 0),

		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "elit.", 0, 1),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Fusce", 6, 1),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "sagittis", 12, 1),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "tincidunt", 21, 1),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "porttitor.", 31, 1),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Donec", 42, 1),

		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "nec", 0, 2),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "fringilla", 4, 2),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "risus.", 14, 2),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Aliquam", 21, 2),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "erat", 29, 2),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "volutpat.", 34, 2),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Nunc", 44, 2),

		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "sodales,", 0, 3),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "orci", 9, 3),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "nec", 14, 3),
		// 		fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "efficitur", 18, 3),
		// 	},
		// },
		{
			in: `Document{ 
				Page{ 
					Paragraph{
						width:50
						align:right
						lineHeight:1
						Text{text:"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce sagittis tincidunt porttitor. Donec nec fringilla risus. Aliquam erat volutpat. Nunc sodales, orci nec efficitur"}
					}
				}
			}`,
			calls: []string{
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Lorem", 0, 0),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "ipsum", 6, 0),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "dolor", 12, 0),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "sit", 18, 0),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "amet,", 22, 0),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "consectetur", 28, 0),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "adipiscing", 40, 0),

				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "elit.", 3, 1),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Fusce", 9, 1),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "sagittis", 15, 1),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "tincidunt", 24, 1),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "porttitor.", 34, 1),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Donec", 45, 1),

				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "nec", 2, 2),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "fringilla", 6, 2),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "risus.", 16, 2),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Aliquam", 23, 2),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "erat", 31, 2),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "volutpat.", 36, 2),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "Nunc", 46, 2),

				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "sodales,", 23, 3),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "orci", 32, 3),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "nec", 37, 3),
				fmt.Sprintf("Text('%s',%d,%d,50,1,BaselineLeft)", "efficitur", 41, 3),
			},
		},
	}

	for i, tt := range tests {

		l := lexer.New(tt.in)
		p := parser.New(l)
		item, err := p.Parse()
		if err != nil {
			t.Fatalf("[%d] parsing failed : %v on : \n%s", i, err, tt.in)
		}

		fpdf := fakePdf{
			d: drawer{
				t:        t,
				charSize: 1,
			},
		}

		r := New(nil, &fpdf)
		err = r.Render(item)
		if err != nil {
			t.Fatalf("[%d] rendering error got  %v ", i, err)
		}

		j := 0
		for _, c := range fpdf.d.callStack {
			if c[:4] != "Text" {
				continue
			}
			if tt.calls[j] != c {
				t.Fatalf("[%d] callstack at %d got '%s' exp '%s'", i, j, c, tt.calls[j])
			}
			j = j + 1
		}
	}
}
