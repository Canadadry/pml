package renderer

import (
	"errors"
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"image/color"
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

func (n *NodePage) Children() []Node { return n.children }
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

type NodeRectangle struct {
	Frame
	children []Node
	color    color.RGBA
}

func (n *NodeRectangle) Children() []Node { return n.children }
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

type NodeFont struct {
	file string
	name string
}

func (n *NodeFont) Children() []Node          { return nil }
func (n *NodeFont) addChild(child Node) error { return errChildrenNotAllowed }
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

type NodeImage struct {
	Frame
	file string
	mode string
}

func (n *NodeImage) Children() []Node          { return nil }
func (n *NodeImage) addChild(child Node) error { return errChildrenNotAllowed }
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

type NodeVector struct {
	Frame
	file string
}

func (n *NodeVector) Children() []Node          { return nil }
func (n *NodeVector) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeVector) initFrom(item *ast.Item) error {
	var err error
	n.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return err
	}

	return n.Frame.initFrom(item)
}

type NodeParagraph struct {
	Frame
	children   []Node
	lineHeight float64
}

func (n *NodeParagraph) Children() []Node { return n.children }
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

type NodeContainer struct {
	Frame
	children []Node
}

func (n *NodeContainer) Children() []Node { return n.children }
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
