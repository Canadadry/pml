package renderer

import (
	"fmt"
	"github.com/canadadry/pml/language/lexer"
	"github.com/canadadry/pml/language/parser"
	"image/color"
	"io"
	"testing"
)

type fakePdf struct {
	d drawer
}

func (fp *fakePdf) Init() PdfDrawer {
	return &fp.d
}

type drawer struct {
	nextError error
	callStack []string
	t         *testing.T
}

func (d *drawer) MoveTo(x float64, y float64) {
	d.t.Fatalf("should not have been called")
}
func (d *drawer) LineTo(x float64, y float64) {
	d.t.Fatalf("should not have been called")
}
func (d *drawer) BezierTo(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) {
	d.t.Fatalf("should not have been called")
}
func (d *drawer) CloseAndDraw(s SvgStyle) {
	d.t.Fatalf("should not have been called")
}
func (d *drawer) AddPage() {
	d.callStack = append(d.callStack, "AddPage")
}
func (d *drawer) SetFillColor(c color.RGBA) {
	d.callStack = append(d.callStack, fmt.Sprintf("SetFillColor(%v)", c))
}
func (d *drawer) Rect(x float64, y float64, w float64, h float64) {
	d.callStack = append(d.callStack, fmt.Sprintf("Rect(%v,%v,%v,%v)", x, y, w, h))
}
func (d *drawer) SetFont(n string, s float64) {
	d.callStack = append(d.callStack, fmt.Sprintf("SetFont('%v',%v)", n, s))
}
func (d *drawer) GetDefaultFontName() string {
	d.callStack = append(d.callStack, "GetDefaultFontName()")
	return "fakefont"
}
func (d *drawer) SetTextColor(c color.RGBA) {
	d.callStack = append(d.callStack, fmt.Sprintf("SetTextColor(%v)", c))
}
func (d *drawer) Text(s string, x float64, y float64, w float64, h float64, a PdfTextAlign) {
	d.callStack = append(d.callStack, fmt.Sprintf("Text('%s',%v,%v,%v,%v,%s)", s, x, y, w, h, a))
}
func (d *drawer) GetTextMaxLength(text string, maxWidth float64) (int, float64) {
	d.callStack = append(d.callStack, fmt.Sprintf("GetTextMaxLength('%s',%g)", text, maxWidth))
	fixedCharSize := int(14)
	maxLen := int(maxWidth) / fixedCharSize
	lenTxt := len(text)
	if maxLen == 0 {
		return lenTxt, float64(lenTxt * fixedCharSize)
	}
	if lenTxt < maxLen {
		return lenTxt, float64(lenTxt * fixedCharSize)
	}
	return maxLen, float64(maxLen * fixedCharSize)
}
func (d *drawer) Image(image io.ReadSeeker, x float64, y float64, w float64, h float64) {
	d.callStack = append(d.callStack, fmt.Sprintf("Image(%v,%v,%v,%v)", x, y, w, h))
}
func (d *drawer) Vector(image io.Reader, x float64, y float64, w float64, h float64) {
	d.callStack = append(d.callStack, fmt.Sprintf("Vector(%v,%v,%v,%v)", x, y, w, h))
}
func (d *drawer) LoadFont(name string, path string) error {
	d.callStack = append(d.callStack, fmt.Sprintf("LoadFont('%s','%s')", name, path))
	return d.nextError
}
func (d *drawer) Output(out io.Writer) error {
	d.callStack = append(d.callStack, "Output")
	return d.nextError
}

func TestRender(t *testing.T) {
	tests := []struct {
		in    string
		calls []string
	}{
		{
			in:    "Document{}",
			calls: []string{"Output"},
		},
		{
			in: "Document{Page{}}",
			calls: []string{
				"AddPage",
				"Output",
			},
		},
		{
			in: "Document{Font{} Page{}}",
			calls: []string{
				"LoadFont('','')",
				"AddPage",
				"Output",
			},
		},
		{
			in: "Document{ Page{ Rectangle{}}}",
			calls: []string{
				"AddPage",
				"SetFillColor({0 0 0 0})",
				"Rect(0,0,0,0)",
				"Output",
			},
		},
		{
			in: "Document{ Page{ Text{}}}",
			calls: []string{
				"AddPage",
				"GetDefaultFontName()",
				"SetFont('fakefont',6)",
				"SetTextColor({0 0 0 0})",
				"Text('',0,0,0,0,TopLeft)",
				"Output",
			},
		},
		{
			in: `Document{ Page{ Image{ file:"node_renderer.go"}}}`,
			calls: []string{
				"AddPage",
				"Image(0,0,0,0)",
				"Output",
			},
		},

		{
			in: `Document{ Page{ Image{ mode: b64 file:"dGVzdA=="}}}`,
			calls: []string{
				"AddPage",
				"Image(0,0,0,0)",
				"Output",
			},
		},
		{
			in: `Document{ Page{ Vector{ file:"node_renderer.go"}}}`,
			calls: []string{
				"AddPage",
				"Vector(0,0,0,0)",
				"Output",
			},
		},
		{
			in: `Document{ Page{ Paragraph{ }}}`,
			calls: []string{
				"AddPage",
				"Output",
			},
		},
		{
			in: `Document{ Page{ Paragraph{Text{} Text{}}}}`,
			calls: []string{
				"AddPage",
				"GetDefaultFontName()",
				"SetFont('fakefont',6)",
				"SetTextColor({0 0 0 0})",
				"GetDefaultFontName()",
				"SetFont('fakefont',6)",
				"SetTextColor({0 0 0 0})",
				"Output",
			},
		},
		{
			in: `Document{ Font{} Page{ Rectangle{Rectangle{}} Paragraph{Text{} Text{}}}}`,
			calls: []string{
				"LoadFont('','')",
				"AddPage",
				"SetFillColor({0 0 0 0})",
				"Rect(0,0,0,0)",
				"SetFillColor({0 0 0 0})",
				"Rect(0,0,0,0)",
				"GetDefaultFontName()",
				"SetFont('fakefont',6)",
				"SetTextColor({0 0 0 0})",
				"GetDefaultFontName()",
				"SetFont('fakefont',6)",
				"SetTextColor({0 0 0 0})",
				"Output",
			},
		},
		{
			in: `Document{ 
				Page{ 
					Container{
						x:100
						y:100

						Image{ 
							x:10
							y:10
							width:80
							height:80
							file:"node_renderer.go"
						}
					}
				}
			}`,
			calls: []string{
				"AddPage",
				"Image(110,110,80,80)",
				"Output",
			},
		},
		{
			in: `Document{ 
				Page{ 
					Rectangle{
						x:100
						y:100
						width:100
						height:100
						color: #ff0000

						Image{ 
							x:10
							y:10
							width:80
							height:80
							file:"node_renderer.go"
						}
					}
				}
			}`,
			calls: []string{
				"AddPage",
				"SetFillColor({255 0 0 255})",
				"Rect(100,100,100,100)",
				"Image(110,110,80,80)",
				"Output",
			},
		},
		{
			in: `Document{ 
				Page{ 
					Paragraph{
						width:100
						Text{text:"mon chien"}
						Text{text:" va bien"}
					}
				}
			}`,
			calls: []string{
				"AddPage",
				"GetDefaultFontName()",
				"SetFont('fakefont',6)",
				"SetTextColor({0 0 0 0})",
				"GetTextMaxLength('mon chien',100)",
				"Text('mon chi',0,0,100,6,BaselineLeft)",
				"GetTextMaxLength('en',2)",
				"Text('en',98,0,100,6,BaselineLeft)",
				"GetDefaultFontName()",
				"SetFont('fakefont',6)",
				"SetTextColor({0 0 0 0})",
				"GetTextMaxLength(' va bien',100)",
				"Text(' va bie',0,6,100,6,BaselineLeft)",
				"GetTextMaxLength('n',2)",
				"Text('n',98,6,100,6,BaselineLeft)",
				"Output",
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
				t: t,
			},
		}

		r := New(nil, &fpdf)
		err = r.Render(item)
		if err != nil {
			t.Fatalf("[%d] rendering failed : %v on : \n%s", i, err, tt.in)
		}

		if len(tt.calls) != len(fpdf.d.callStack) {
			t.Fatalf("[%d] callstack len got %d exp %d\n got %v\n exp %v on : \n%s", i, len(fpdf.d.callStack), len(tt.calls), fpdf.d.callStack, tt.calls, tt.in)
		}

		for j := range tt.calls {
			if tt.calls[j] != fpdf.d.callStack[j] {
				t.Fatalf("[%d] callstack at %d got '%s' exp '%s' on : \n%s", i, j, fpdf.d.callStack[j], tt.calls[j], tt.in)
			}
		}
	}
}
