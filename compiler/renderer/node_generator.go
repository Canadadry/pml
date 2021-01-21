package renderer

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"image/color"
	"io"
	"os"
)

const (
	itemDocument  = "Document"
	itemPage      = "Page"
	itemRectangle = "Rectangle"
	itemText      = "Text"
	itemFont      = "Font"
	itemImage     = "Image"
	itemVector    = "Vector"
	itemParagraph = "Paragraph"
	itemContainer = "Container"
)

const (
	Left   = "left"
	Center = "center"
	Right  = "right"
	Top    = "top"
	Middle = "middle"
	Bottom = "bottom"
	Fill   = "fill"
	Layout = "layout"
	Free   = "free"

	ImgModeFile = "file"
	ImgModeB64  = "b64"
)

var (
	errItemNotFound        = errors.New("errItemNotFound")
	rootMustBeDocumentItem = errors.New("rootMustBeDocumentItem")
	errChildrenNotAllowed  = errors.New("errChildrenNotAllowed")
)

type Node interface {
	Children() []Node
	addChild(child Node) error
	initFrom(*ast.Item) error
	draw(pdf PdfDrawer, rb renderBox) (renderBox, error)
	needToDrawChild() bool
}

type renderBox struct {
	x float64
	y float64
	w float64
	h float64
}

func (rb renderBox) Cut(f Frame) renderBox {
	rb.x = rb.x + f.x
	rb.y = rb.y + f.y
	rb.w = f.width
	rb.h = f.height
	return rb
}

func generate(item *ast.Item) (Node, error) {

	if item.TokenType.Literal != itemDocument {
		return nil, fmt.Errorf("%w : got %s exp %s", rootMustBeDocumentItem, item.TokenType.Literal, itemDocument)
	}
	return generateItem(item)
}

func generateItem(item *ast.Item) (Node, error) {
	n, err := createNodeFrom(item)
	if err != nil {
		return nil, err
	}
	err = n.initFrom(item)
	if err != nil {
		return nil, err
	}
	for _, c := range item.Children {
		child, err := generateItem(&c)
		if err != nil {
			return nil, err
		}
		err = n.addChild(child)
		if err != nil {
			return nil, err
		}
	}
	return n, nil
}

func createNodeFrom(item *ast.Item) (Node, error) {
	switch item.TokenType.Literal {
	case itemDocument:
		return &NodeDocument{children: []Node{}}, nil
	case itemPage:
		return &NodePage{children: []Node{}}, nil
	case itemRectangle:
		return &NodeRectangle{children: []Node{}}, nil
	case itemText:
		return &NodeText{}, nil
	case itemFont:
		return &NodeFont{}, nil
	case itemImage:
		return &NodeImage{}, nil
	case itemVector:
		return &NodeVector{}, nil
	case itemParagraph:
		return &NodeParagraph{}, nil
	case itemContainer:
		return &NodeContainer{}, nil
	}
	return nil, errItemNotFound
}

type NodeDocument struct {
	children []Node
}

func (n *NodeDocument) Children() []Node { return n.children }
func (n *NodeDocument) addChild(child Node) error {
	switch child.(type) {
	case *NodeDocument:
		return errChildrenNotAllowed
	case *NodePage:
	case *NodeRectangle:
		return errChildrenNotAllowed
	case *NodeText:
		return errChildrenNotAllowed
	case *NodeFont:
	case *NodeImage:
		return errChildrenNotAllowed
	case *NodeVector:
		return errChildrenNotAllowed
	case *NodeParagraph:
		return errChildrenNotAllowed
	case *NodeContainer:
		return errChildrenNotAllowed
	}
	n.children = append(n.children, child)
	return nil
}
func (n *NodeDocument) initFrom(*ast.Item) error { return nil }
func (n *NodeDocument) needToDrawChild() bool    { return true }
func (n *NodeDocument) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	return rb, nil
}

type Frame struct {
	x      float64
	y      float64
	width  float64
	height float64
	xAlign string
	yAlign string
}

func (f *Frame) initFrom(item *ast.Item) error {
	var err error
	f.x, err = item.GetPropertyAsFloatWithDefault("x", 0)
	if err != nil {
		return err
	}
	f.y, err = item.GetPropertyAsFloatWithDefault("y", 0)
	if err != nil {
		return err
	}
	f.width, err = item.GetPropertyAsFloatWithDefault("width", 0)
	if err != nil {
		return err
	}
	f.height, err = item.GetPropertyAsFloatWithDefault("height", 0)
	if err != nil {
		return err
	}
	xvalues := []string{Left, Center, Right, Fill, Layout, Free}
	f.xAlign, err = item.GetPropertyAsIdentifierFromListWithDefault("xAlign", xvalues[5], xvalues)
	if err != nil {
		return err
	}
	yvalues := []string{Top, Middle, Bottom, Fill, Layout, Free}
	f.yAlign, err = item.GetPropertyAsIdentifierFromListWithDefault("yAlign", yvalues[5], yvalues)
	if err != nil {
		return err
	}
	return nil
}

type NodePage struct {
	children []Node
}

func (n *NodePage) Children() []Node      { return n.children }
func (n *NodePage) needToDrawChild() bool { return true }
func (n *NodePage) addChild(child Node) error {
	switch child.(type) {
	case *NodeDocument:
		return errChildrenNotAllowed
	case *NodePage:
		return errChildrenNotAllowed
	case *NodeRectangle:
	case *NodeText:
	case *NodeFont:
		return errChildrenNotAllowed
	case *NodeImage:
	case *NodeVector:
	case *NodeParagraph:
	case *NodeContainer:
	}
	n.children = append(n.children, child)
	return nil
}
func (n *NodePage) initFrom(*ast.Item) error { return nil }
func (n *NodePage) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.AddPage()
	return renderBox{w: PdfWidthMm, h: PdfHeight}, nil
}

type NodeRectangle struct {
	Frame
	children []Node
	color    color.RGBA
}

func (n *NodeRectangle) Children() []Node      { return n.children }
func (n *NodeRectangle) needToDrawChild() bool { return true }
func (n *NodeRectangle) addChild(child Node) error {
	switch child.(type) {
	case *NodeDocument:
		return errChildrenNotAllowed
	case *NodePage:
		return errChildrenNotAllowed
	case *NodeRectangle:
	case *NodeText:
	case *NodeFont:
		return errChildrenNotAllowed
	case *NodeImage:
	case *NodeVector:
	case *NodeParagraph:
	case *NodeContainer:
		return errChildrenNotAllowed
	}
	n.children = append(n.children, child)
	return nil
}
func (n *NodeRectangle) initFrom(item *ast.Item) error {
	var err error

	n.color, err = item.GetPropertyAsColorWithDefault("color", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return err
	}
	return n.Frame.initFrom(item)
}
func (n *NodeRectangle) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.SetFillColor(n.color)
	rb = rb.Cut(n.Frame)
	pdf.Rect(rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}

type NodeText struct {
	Frame
	text     string
	color    color.RGBA
	align    string
	fontName string
	fontSize float64
}

func (n *NodeText) Children() []Node          { return nil }
func (n *NodeText) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeText) needToDrawChild() bool     { return true }
func (n *NodeText) initFrom(item *ast.Item) error {
	var err error
	n.text, err = item.GetPropertyAsStringWithDefault("text", "")
	if err != nil {
		return err
	}
	n.color, err = item.GetPropertyAsColorWithDefault("color", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return err
	}
	values := []string{
		string(AlingTopLeft),
		string(AlingTopCenter),
		string(AlingTopRight),
		string(AlingMiddleLeft),
		string(AlingMiddleCenter),
		string(AlingMiddleRight),
		string(AlingBottomLeft),
		string(AlingBottomCenter),
		string(AlingBottomRight),
		string(AlingBaselineLeft),
		string(AlingBaselineCenter),
		string(AlingBaselineRight),
	}
	n.align, err = item.GetPropertyAsIdentifierFromListWithDefault("align", values[0], values)
	if err != nil {
		return err
	}
	n.fontName, err = item.GetPropertyAsStringWithDefault("fontName", "")
	if err != nil {
		return err
	}
	n.fontSize, err = item.GetPropertyAsFloatWithDefault("fontSize", 6)
	if err != nil {
		return err
	}
	return n.Frame.initFrom(item)
}
func (n *NodeText) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {

	if len(n.fontName) == 0 {
		n.fontName = pdf.GetDefaultFontName()
	}
	pdf.SetFont(n.fontName, n.fontSize)
	pdf.SetTextColor(n.color)
	rb = rb.Cut(n.Frame)
	pdf.Text(n.text, rb.x, rb.y, rb.w, rb.h, PdfTextAlign(n.align))
	return rb, nil
}

type NodeFont struct {
	file string
	name string
}

func (n *NodeFont) Children() []Node          { return nil }
func (n *NodeFont) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeFont) needToDrawChild() bool     { return true }
func (n *NodeFont) initFrom(item *ast.Item) error {
	var err error
	n.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return err
	}
	n.name, err = item.GetPropertyAsStringWithDefault("name", "")
	if err != nil {
		return err
	}
	return nil
}
func (n *NodeFont) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.LoadFont(n.name, n.file)
	return rb, nil
}

type NodeImage struct {
	Frame
	file string
	mode string
}

func (n *NodeImage) Children() []Node          { return nil }
func (n *NodeImage) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeImage) needToDrawChild() bool     { return true }
func (n *NodeImage) initFrom(item *ast.Item) error {
	var err error
	n.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return err
	}
	values := []string{ImgModeFile, ImgModeB64}
	n.mode, err = item.GetPropertyAsIdentifierFromListWithDefault("mode", ImgModeFile, values)
	if err != nil {
		return err
	}

	return n.Frame.initFrom(item)
}
func (n *NodeImage) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	if len(n.file) == 0 {
		return rb, ErrEmptyImageFileProperty
	}
	var rs io.ReadSeeker
	if n.mode == ImgModeFile {
		file, err := os.Open(n.file)
		if err != nil {
			return rb, fmt.Errorf("%w : %v", ErrCantOpenFile, err)
		}
		defer file.Close()
		rs = file
	} else {
		decoded, err := base64.StdEncoding.DecodeString(n.file)
		if err != nil {
			return rb, fmt.Errorf("%w : %v", ErrB64Read, err)
		}
		rs = bytes.NewReader(decoded)
	}
	rb = rb.Cut(n.Frame)
	pdf.Image(rs, rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}

type NodeVector struct {
	Frame
	file string
}

func (n *NodeVector) Children() []Node          { return nil }
func (n *NodeVector) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeVector) needToDrawChild() bool     { return true }
func (n *NodeVector) initFrom(item *ast.Item) error {
	var err error
	n.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return err
	}

	return n.Frame.initFrom(item)
}
func (n *NodeVector) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	if len(n.file) == 0 {
		return rb, ErrEmptyImageFileProperty
	}
	file, err := os.Open(n.file)
	if err != nil {
		return rb, fmt.Errorf("%w : %v", ErrCantOpenFile, err)
	}
	defer file.Close()
	rb = rb.Cut(n.Frame)
	pdf.Vector(file, rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}

type NodeParagraph struct {
	Frame
	children   []Node
	lineHeight float64
}

func (n *NodeParagraph) Children() []Node      { return n.children }
func (n *NodeParagraph) needToDrawChild() bool { return false }
func (n *NodeParagraph) addChild(child Node) error {
	switch child.(type) {
	case *NodeDocument:
		return errChildrenNotAllowed
	case *NodePage:
		return errChildrenNotAllowed
	case *NodeRectangle:
		return errChildrenNotAllowed
	case *NodeText:
	case *NodeFont:
		return errChildrenNotAllowed
	case *NodeImage:
		return errChildrenNotAllowed
	case *NodeVector:
		return errChildrenNotAllowed
	case *NodeParagraph:
		return errChildrenNotAllowed
	case *NodeContainer:
		return errChildrenNotAllowed
	}
	n.children = append(n.children, child)
	return nil
}
func (n *NodeParagraph) initFrom(item *ast.Item) error {
	var err error

	n.lineHeight, err = item.GetPropertyAsFloatWithDefault("lineHeight", 6)
	if err != nil {
		return err
	}
	xvalues := []string{Left, Center, Right, Fill, Layout, Free}
	n.xAlign, err = item.GetPropertyAsIdentifierFromListWithDefault("xAlign", xvalues[5], xvalues)
	if err != nil {
		return err
	}
	yvalues := []string{Top, Middle, Bottom, Fill, Layout, Free}
	n.yAlign, err = item.GetPropertyAsIdentifierFromListWithDefault("yAlign", yvalues[5], yvalues)
	if err != nil {
		return err
	}

	return n.Frame.initFrom(item)
}
func (n *NodeParagraph) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	x := 0.0
	y := 0.0
	rb = rb.Cut(n.Frame)

	for _, child := range n.children {
		offset := 0
		textChild, ok := child.(*NodeText)
		if !ok {
			return rb, fmt.Errorf("Unexpected node in paragraph")
		}
		if len(textChild.fontName) == 0 {
			textChild.fontName = pdf.GetDefaultFontName()
		}

		pdf.SetFont(textChild.fontName, textChild.fontSize)
		pdf.SetTextColor(textChild.color)

		for offset < len(textChild.text) {
			maxSize, textWidth := pdf.GetTextMaxLength(textChild.text[offset:], n.width-x)
			text := textChild.text[offset : offset+maxSize]
			pdf.Text(text, x+rb.x, y+rb.y, rb.w, n.lineHeight, "BaselineLeft")
			offset = offset + maxSize
			x = x + textWidth
			if x > rb.w {
				x = 0
				y = y + n.lineHeight
			}
		}
	}
	return rb, nil
}

type NodeContainer struct {
	Frame
	children []Node
}

func (n *NodeContainer) Children() []Node      { return n.children }
func (n *NodeContainer) needToDrawChild() bool { return true }
func (n *NodeContainer) addChild(child Node) error {
	switch child.(type) {
	case *NodeDocument:
		return errChildrenNotAllowed
	case *NodePage:
		return errChildrenNotAllowed
	case *NodeRectangle:
	case *NodeText:
	case *NodeFont:
		return errChildrenNotAllowed
	case *NodeImage:
	case *NodeVector:
	case *NodeParagraph:
	case *NodeContainer:
	}
	n.children = append(n.children, child)
	return nil
}
func (n *NodeContainer) initFrom(item *ast.Item) error {
	return n.Frame.initFrom(item)
}

func (n *NodeContainer) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	return rb.Cut(n.Frame), nil
}
