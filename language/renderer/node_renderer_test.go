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
	d.callStack = append(d.callStack, fmt.Sprintf("SetFont(%v,%v)", n, s))
}
func (d *drawer) GetDefaultFontName() string {
	d.callStack = append(d.callStack, "GetDefaultFontName()")
	return "fakefont"
}
func (d *drawer) SetTextColor(c color.RGBA) {
	d.callStack = append(d.callStack, fmt.Sprintf("SetTextColor(%v)", c))
}
func (d *drawer) Text(string, float64, float64, float64, float64, PdfTextAlign) {
	d.callStack = append(d.callStack, "Text")
}
func (d *drawer) GetTextMaxLength(text string, maxWidth float64) (int, float64) {
	d.callStack = append(d.callStack, "GetTextMaxLength")
	return 0, 0
}
func (d *drawer) Image(image io.ReadSeeker, x float64, y float64, width float64, height float64) {
	d.callStack = append(d.callStack, "Image")
}
func (d *drawer) Vector(image io.Reader, x float64, y float64, width float64, height float64) {
	d.callStack = append(d.callStack, "Vector")
}
func (d *drawer) LoadFont(fontName string, fontFilePath string) error {
	d.callStack = append(d.callStack, "LoadFont")
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
				"LoadFont",
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
				"SetFont(fakefont,6)",
				"SetTextColor({0 0 0 0})",
				"Text",
				"Output",
			},
		},
		{
			in: `Document{ Page{ Image{ file:"node_renderer.go"}}}`,
			calls: []string{
				"AddPage",
				"Image",
				"Output",
			},
		},
		{
			in: `Document{ Page{ Vector{ file:"node_renderer.go"}}}`,
			calls: []string{
				"AddPage",
				"Vector",
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
				"SetFont(fakefont,6)",
				"SetTextColor({0 0 0 0})",
				"GetDefaultFontName()",
				"SetFont(fakefont,6)",
				"SetTextColor({0 0 0 0})",
				"Output",
			},
		},
		{
			in: `Document{ Font{} Page{ Rectangle{Rectangle{}} Paragraph{Text{} Text{}}}}`,
			calls: []string{
				"LoadFont",
				"AddPage",
				"SetFillColor({0 0 0 0})",
				"Rect(0,0,0,0)",
				"SetFillColor({0 0 0 0})",
				"Rect(0,0,0,0)",
				"GetDefaultFontName()",
				"SetFont(fakefont,6)",
				"SetTextColor({0 0 0 0})",
				"GetDefaultFontName()",
				"SetFont(fakefont,6)",
				"SetTextColor({0 0 0 0})",
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
			t.Fatalf("[%d] callstack len got %d exp %d\n got %v\n exp %v", i, len(fpdf.d.callStack), len(tt.calls), fpdf.d.callStack, tt.calls)
		}

		for j := range tt.calls {
			if tt.calls[j] != fpdf.d.callStack[j] {
				t.Fatalf("[%d] callstack at %d got '%s' exp '%s'", i, j, fpdf.d.callStack[j], tt.calls[j])
			}
		}
	}
}
