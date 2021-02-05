package renderer

import (
	"errors"
	"github.com/canadadry/pml/compiler/ast"
	"github.com/canadadry/pml/compiler/lexer"
	"github.com/canadadry/pml/compiler/parser"
	"reflect"
	"strings"
	"testing"
)

func TestGenerateFrom(t *testing.T) {
	tests := []struct {
		program string
		result  Node
	}{
		{
			program: `Document{
				Page{
					Rectangle{
						x:11
						y:21
						width: 31
						height: 41
					}
					Text{
						x:12
						y:22
						width: 32
						height: 42
						text:"text"
						align: TopLeft
					}
					Image{
						x:13
						y:23
						width: 33
						height: 43
						file:"filename"
					}
					Vector{
						x:14
						y:24
						width: 34
						height: 44
						file:"vecname"
					}
					Paragraph{
						x:15
						y:25
						width: 35
						height: 45
					}
					Container{
						x:16
						y:26
					}
				}
				Font{
					file:"testfile"
					name:"testname"
				}
			}`,
			result: &NodeDocument{
				children: []Node{
					&NodePage{
						children: []Node{
							&NodeRectangle{
								Frame: Frame{
									x:       11,
									y:       21,
									width:   31,
									height:  41,
									xAnchor: Relative,
									yAnchor: Relative,
								},
							},
							&NodeText{
								Frame: Frame{
									x:       12,
									y:       22,
									width:   32,
									height:  42,
									xAnchor: Relative,
									yAnchor: Relative,
								},
								text:  "text",
								align: "TopLeft",
							},
							&NodeImage{
								Frame: Frame{
									x:       13,
									y:       23,
									width:   33,
									height:  43,
									xAnchor: Relative,
									yAnchor: Relative,
								},
								file: "filename",
								mode: "file",
							},
							&NodeVector{
								Frame: Frame{
									x:       14,
									y:       24,
									width:   34,
									height:  44,
									xAnchor: Relative,
									yAnchor: Relative,
								},
								file: "vecname",
							},
							&NodeParagraph{
								Frame: Frame{
									x:       15,
									y:       25,
									width:   35,
									height:  45,
									xAnchor: Relative,
									yAnchor: Relative,
								},
								lineHeight: 6,
							},
							&NodeContainer{
								Frame: Frame{
									x:       16,
									y:       26,
									xAnchor: Relative,
									yAnchor: Relative,
								},
							},
						},
					},
					&NodeFont{
						file: "testfile",
						name: "testname",
					},
				},
			},
		},
	}

	for i, tt := range tests {
		l := lexer.New(tt.program)
		p := parser.New(l)
		item, err := p.Parse()
		if err != nil {
			t.Fatalf("[%d]parsing failed : %v on : \n%s", i, err, tt.program)
		}
		n, err := generate(item)
		if err != nil {
			t.Fatalf("[%d]generator failed : %v on : \n%s", i, err, tt.program)
		}
		testNodeMatch(t, i, n, tt.result)
	}
}

func testNodeMatch(t *testing.T, i int, got Node, exp Node) {
	if reflect.TypeOf(got) != reflect.TypeOf(exp) {
		t.Fatalf("[%d] generator failed got type %s expected %s", i, reflect.TypeOf(got), reflect.TypeOf(exp))
	}

	if len(got.Children()) != len(exp.Children()) {
		t.Fatalf("[%d] generator failed in %s got %d children expected %d", i, reflect.TypeOf(got), len(got.Children()), len(exp.Children()))
	}

	switch realGot := got.(type) {
	case *NodeDocument:
	case *NodePage:
	case *NodeRectangle:
		realExp, _ := exp.(*NodeRectangle)
		testNodeRectangle(t, i, *realGot, *realExp)
	case *NodeText:
		realExp, _ := exp.(*NodeText)
		testNodeText(t, i, *realGot, *realExp)
	case *NodeFont:
		realExp, _ := exp.(*NodeFont)
		testNodeFont(t, i, *realGot, *realExp)
	case *NodeImage:
		realExp, _ := exp.(*NodeImage)
		testNodeImage(t, i, *realGot, *realExp)
	case *NodeVector:
		realExp, _ := exp.(*NodeVector)
		testNodeVector(t, i, *realGot, *realExp)
	case *NodeParagraph:
		realExp, _ := exp.(*NodeParagraph)
		testNodeParagraph(t, i, *realGot, *realExp)
	case *NodeContainer:
		realExp, _ := exp.(*NodeContainer)
		testNodeContainer(t, i, *realGot, *realExp)
	default:
		t.Fatalf("unknown node %T", got)
	}

	for i := range got.Children() {
		testNodeMatch(t, i, got.Children()[i], exp.Children()[i])
	}
}

func testFrame(t *testing.T, i int, in string, got Frame, exp Frame) {
	testFloatProperty(t, i, in, "x", got.x, exp.x)
	testFloatProperty(t, i, in, "y", got.y, exp.y)
	testFloatProperty(t, i, in, "width", got.width, exp.width)
	testFloatProperty(t, i, in, "height", got.height, exp.height)
	testStringProperty(t, i, in, "xAnchor", got.xAnchor, exp.xAnchor)
	testStringProperty(t, i, in, "yAnchor", got.yAnchor, exp.yAnchor)
}

func testNodeRectangle(t *testing.T, i int, got NodeRectangle, exp NodeRectangle) {
	testFrame(t, i, "NodeRectangle", got.Frame, exp.Frame)
	// testColorProperty(t, i, "NodeRectangle", "color", got.color, exp.color)
}

func testNodeText(t *testing.T, i int, got NodeText, exp NodeText) {
	testFrame(t, i, "NodeText", got.Frame, exp.Frame)
	testStringProperty(t, i, "NodeText", "text", got.text, exp.text)
	testStringProperty(t, i, "NodeText", "align", got.align, exp.align)
	// testColorProperty(t, i, "NodeText", "color", got.color, exp.color)
}

func testNodeFont(t *testing.T, i int, got NodeFont, exp NodeFont) {
	testStringProperty(t, i, "NodeFont", "file", got.file, exp.file)
	testStringProperty(t, i, "NodeFont", "name", got.name, exp.name)
}

func testNodeImage(t *testing.T, i int, got NodeImage, exp NodeImage) {
	testFrame(t, i, "NodeImage", got.Frame, exp.Frame)
	testStringProperty(t, i, "NodeImage", "file", got.file, exp.file)
	testStringProperty(t, i, "NodeImage", "mode", got.mode, exp.mode)
}

func testNodeVector(t *testing.T, i int, got NodeVector, exp NodeVector) {
	testFrame(t, i, "NodeVector", got.Frame, exp.Frame)
	testStringProperty(t, i, "NodeVector", "file", got.file, exp.file)
}

func testNodeParagraph(t *testing.T, i int, got NodeParagraph, exp NodeParagraph) {
	testFrame(t, i, "NodeParagraph", got.Frame, exp.Frame)
	testFloatProperty(t, i, "NodeParagraph", "lineHeight", got.lineHeight, exp.lineHeight)
}

func testNodeContainer(t *testing.T, i int, got NodeContainer, exp NodeContainer) {
	testFrame(t, i, "NodeContainer", got.Frame, exp.Frame)
}

func testFloatProperty(t *testing.T, i int, node string, property string, got float64, exp float64) {
	if got != exp {
		t.Fatalf("[%d] generator failed in %s property %s : got %g expected %g", i, node, property, got, exp)
	}
}

func testStringProperty(t *testing.T, i int, node string, property string, got string, exp string) {
	if got != exp {
		t.Fatalf("[%d] generator failed in %s property %s : got '%s' expected '%s'", i, node, property, got, exp)
	}
}

func TestGenerateFrom_Hierarchy(t *testing.T) {
	tests := []struct {
		err      error
		hierachy []string
	}{
		{err: rootMustBeDocumentItem, hierachy: []string{"Page"}},
		{err: rootMustBeDocumentItem, hierachy: []string{"Rectangle"}},
		{err: rootMustBeDocumentItem, hierachy: []string{"Text"}},
		{err: rootMustBeDocumentItem, hierachy: []string{"Image"}},
		{err: rootMustBeDocumentItem, hierachy: []string{"Vector"}},
		{err: rootMustBeDocumentItem, hierachy: []string{"Container"}},
		{err: rootMustBeDocumentItem, hierachy: []string{"Paragraph"}},
		{err: rootMustBeDocumentItem, hierachy: []string{"Font"}},
		{hierachy: []string{"Document"}},

		{err: errItemNotFound, hierachy: []string{"Document", "Fake"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Document"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Rectangle"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Text"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Image"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Vector"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Container"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Paragraph"}},
		{hierachy: []string{"Document", "Font"}},
		{hierachy: []string{"Document", "Page"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Document"}},
		{hierachy: []string{"Document", "Page", "Rectangle"}},
		{hierachy: []string{"Document", "Page", "Text"}},
		{hierachy: []string{"Document", "Page", "Image"}},
		{hierachy: []string{"Document", "Page", "Vector"}},
		{hierachy: []string{"Document", "Page", "Container"}},
		{hierachy: []string{"Document", "Page", "Paragraph"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Page"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Rectangle", "Document"}},
		{hierachy: []string{"Document", "Page", "Rectangle", "Rectangle"}},
		{hierachy: []string{"Document", "Page", "Rectangle", "Text"}},
		{hierachy: []string{"Document", "Page", "Rectangle", "Image"}},
		{hierachy: []string{"Document", "Page", "Rectangle", "Vector"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Rectangle", "Container"}},
		{hierachy: []string{"Document", "Page", "Rectangle", "Paragraph"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Rectangle", "Page"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Rectangle", "Font"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Document"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Rectangle"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Text"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Image"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Vector"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Container"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Paragraph"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Page"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Text", "Font"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Document"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Rectangle"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Text"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Image"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Vector"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Container"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Paragraph"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Page"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Image", "Font"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Document"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Rectangle"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Text"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Image"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Vector"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Container"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Paragraph"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Page"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Vector", "Font"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Container", "Document"}},
		{hierachy: []string{"Document", "Page", "Container", "Rectangle"}},
		{hierachy: []string{"Document", "Page", "Container", "Text"}},
		{hierachy: []string{"Document", "Page", "Container", "Image"}},
		{hierachy: []string{"Document", "Page", "Container", "Vector"}},
		{hierachy: []string{"Document", "Page", "Container", "Container"}},
		{hierachy: []string{"Document", "Page", "Container", "Paragraph"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Container", "Page"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Container", "Font"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Paragraph", "Document"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Paragraph", "Rectangle"}},
		{hierachy: []string{"Document", "Page", "Paragraph", "Text"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Paragraph", "Image"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Paragraph", "Vector"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Paragraph", "Container"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Paragraph", "Paragraph"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Paragraph", "Page"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Paragraph", "Font"}},

		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Document"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Rectangle"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Text"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Image"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Vector"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Container"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Paragraph"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Page"}},
		{err: errChildrenNotAllowed, hierachy: []string{"Document", "Page", "Font", "Font"}},
	}
	for i, tt := range tests {
		program := strings.Join(tt.hierachy, " { ")
		program += " { "
		program += strings.Repeat(" } ", len(tt.hierachy))

		l := lexer.New(program)
		p := parser.New(l)
		item, err := p.Parse()
		if err != nil {
			t.Fatalf("[%d] parsing failed : %v on : \n%s", i, err, program)
		}
		_, err = generate(item)
		if !errors.Is(err, tt.err) {
			t.Fatalf("[%d] generator failed got %v expected %v on : \n%s", i, err, tt.err, program)
		}
	}
}

func TestGenerateFrom_Property(t *testing.T) {
	tests := []struct {
		err      error
		hierachy []string
		property string
		value    string
	}{

		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Rectangle"}, value: `"str"`, property: "x"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Rectangle"}, value: `"str"`, property: "y"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Rectangle"}, value: `"str"`, property: "width"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Rectangle"}, value: `"str"`, property: "height"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Rectangle"}, value: `"str"`, property: "color"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Rectangle"}, value: `"str"`, property: "xAnchor"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Rectangle"}, value: `"str"`, property: "yAnchor"},

		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"x"`, property: "x"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"x"`, property: "y"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"x"`, property: "width"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"x"`, property: "height"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"x"`, property: "color"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"x"`, property: "align"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `1.0`, property: "text"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `1.0`, property: "fontName"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"x"`, property: "fontSize"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"str"`, property: "xAnchor"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Text"}, value: `"str"`, property: "yAnchor"},

		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Image"}, value: `"str"`, property: "x"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Image"}, value: `"str"`, property: "y"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Image"}, value: `"str"`, property: "width"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Image"}, value: `"str"`, property: "height"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Image"}, value: `12334`, property: "file"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Image"}, value: `"str"`, property: "mode"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Image"}, value: `"str"`, property: "xAnchor"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Image"}, value: `"str"`, property: "yAnchor"},

		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Vector"}, value: `"str"`, property: "x"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Vector"}, value: `"str"`, property: "y"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Vector"}, value: `"str"`, property: "width"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Vector"}, value: `"str"`, property: "height"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Vector"}, value: `12334`, property: "file"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Vector"}, value: `"str"`, property: "xAnchor"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Vector"}, value: `"str"`, property: "yAnchor"},

		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Container"}, value: `"x"`, property: "x"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Container"}, value: `"x"`, property: "y"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Container"}, value: `"str"`, property: "xAnchor"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Container"}, value: `"str"`, property: "yAnchor"},

		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Paragraph"}, value: `"str"`, property: "lineHeight"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Paragraph"}, value: `"str"`, property: "x"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Paragraph"}, value: `"str"`, property: "y"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Paragraph"}, value: `"str"`, property: "width"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Paragraph"}, value: `"str"`, property: "height"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Paragraph"}, value: `"str"`, property: "xAnchor"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Page", "Paragraph"}, value: `"str"`, property: "yAnchor"},

		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Font"}, value: `1`, property: "file"},
		{err: ast.ErrInvalidTypeForProperty, hierachy: []string{"Document", "Font"}, value: `1`, property: "name"},
	}
	for i, tt := range tests {
		program := strings.Join(tt.hierachy, " { ")
		program += " { " + tt.property + " : " + tt.value
		program += strings.Repeat(" } ", len(tt.hierachy))

		l := lexer.New(program)
		p := parser.New(l)
		item, err := p.Parse()
		if err != nil {
			t.Fatalf("[%d] parsing failed : %v on : \n%s", i, err, program)
		}
		_, err = generate(item)
		if !errors.Is(err, tt.err) {
			t.Fatalf("[%d] generator failed got %v expected %v on : \n%s", i, err, tt.err, program)
		}
	}
}
