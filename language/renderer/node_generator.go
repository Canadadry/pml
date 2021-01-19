package renderer

import (
	"errors"
	"fmt"
	"github.com/canadadry/pml/language/ast"
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

func GenerateFrom(item *ast.Item) (Node, error) {

	if item.TokenType.Literal != itemDocument {
		return nil, fmt.Errorf("%w : got %s exp %s", rootMustBeDocumentItem, item.TokenType.Literal, itemDocument)
	}
	return generateFrom(item)
}

func generateFrom(item *ast.Item) (Node, error) {
	n, err := createNodeFrom(item)
	if err != nil {
		return nil, err
	}
	err = n.initFrom(item)
	if err != nil {
		return nil, err
	}
	for _, c := range item.Children {
		child, err := generateFrom(&c)
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

func (nd *NodeDocument) Children() []Node { return nd.children }
func (nd *NodeDocument) addChild(child Node) error {
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
	nd.children = append(nd.children, child)
	return nil
}
func (nd *NodeDocument) initFrom(*ast.Item) error { return nil }

type NodePage struct {
	children []Node
}

func (np *NodePage) Children() []Node { return np.children }
func (np *NodePage) addChild(child Node) error {
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
	np.children = append(np.children, child)
	return nil
}
func (np *NodePage) initFrom(*ast.Item) error { return nil }

type NodeRectangle struct {
	children []Node
	x        float64
	y        float64
	width    float64
	height   float64
	color    color.RGBA
}

func (nr *NodeRectangle) Children() []Node { return nr.children }
func (nr *NodeRectangle) addChild(child Node) error {
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
	nr.children = append(nr.children, child)
	return nil
}
func (nr *NodeRectangle) initFrom(item *ast.Item) error {
	var err error

	nr.color, err = item.GetPropertyAsColorWithDefault("color", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return err
	}
	nr.x, err = item.GetPropertyAsFloatWithDefault("x", 0)
	if err != nil {
		return err
	}
	nr.y, err = item.GetPropertyAsFloatWithDefault("y", 0)
	if err != nil {
		return err
	}
	nr.width, err = item.GetPropertyAsFloatWithDefault("width", 0)
	if err != nil {
		return err
	}
	nr.height, err = item.GetPropertyAsFloatWithDefault("height", 0)
	if err != nil {
		return err
	}
	return nil
}

type NodeText struct {
	text     string
	x        float64
	y        float64
	width    float64
	height   float64
	color    color.RGBA
	align    string
	fontName string
	fontSize float64
}

func (nt *NodeText) Children() []Node          { return nil }
func (nt *NodeText) addChild(child Node) error { return errChildrenNotAllowed }
func (nt *NodeText) initFrom(item *ast.Item) error {
	var err error
	nt.text, err = item.GetPropertyAsStringWithDefault("text", "")
	if err != nil {
		return err
	}
	nt.x, err = item.GetPropertyAsFloatWithDefault("x", 0)
	if err != nil {
		return err
	}
	nt.y, err = item.GetPropertyAsFloatWithDefault("y", 0)
	if err != nil {
		return err
	}
	nt.width, err = item.GetPropertyAsFloatWithDefault("width", 0)
	if err != nil {
		return err
	}
	nt.height, err = item.GetPropertyAsFloatWithDefault("height", 0)
	if err != nil {
		return err
	}
	nt.color, err = item.GetPropertyAsColorWithDefault("color", color.RGBA{0, 0, 0, 0})
	if err != nil {
		return err
	}
	nt.align, err = item.GetPropertyAsIdentifierWithDefault("align", "TopLeft")
	if err != nil {
		return err
	}
	nt.fontName, err = item.GetPropertyAsStringWithDefault("fontName", "")
	if err != nil {
		return err
	}
	nt.fontSize, err = item.GetPropertyAsFloatWithDefault("fontSize", 6)
	if err != nil {
		return err
	}
	return nil
}

type NodeFont struct {
	file string
	name string
}

func (nf *NodeFont) Children() []Node          { return nil }
func (nf *NodeFont) addChild(child Node) error { return errChildrenNotAllowed }
func (nf *NodeFont) initFrom(item *ast.Item) error {
	var err error
	nf.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return err
	}
	nf.name, err = item.GetPropertyAsStringWithDefault("name", "")
	if err != nil {
		return err
	}
	return nil
}

type NodeImage struct {
	file   string
	mode   string
	x      float64
	y      float64
	width  float64
	height float64
}

func (ni *NodeImage) Children() []Node          { return nil }
func (ni *NodeImage) addChild(child Node) error { return errChildrenNotAllowed }
func (ni *NodeImage) initFrom(item *ast.Item) error {
	var err error
	ni.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return err
	}
	ni.mode, err = item.GetPropertyAsIdentifierWithDefault("mode", ImgModeFile)
	if err != nil {
		return err
	}
	ni.x, err = item.GetPropertyAsFloatWithDefault("x", 0)
	if err != nil {
		return err
	}
	ni.y, err = item.GetPropertyAsFloatWithDefault("y", 0)
	if err != nil {
		return err
	}
	ni.width, err = item.GetPropertyAsFloatWithDefault("width", 0)
	if err != nil {
		return err
	}
	ni.height, err = item.GetPropertyAsFloatWithDefault("height", 0)
	if err != nil {
		return err
	}
	return nil
}

type NodeVector struct {
	file   string
	x      float64
	y      float64
	width  float64
	height float64
}

func (nv *NodeVector) Children() []Node          { return nil }
func (nv *NodeVector) addChild(child Node) error { return errChildrenNotAllowed }
func (nv *NodeVector) initFrom(item *ast.Item) error {
	var err error
	nv.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return err
	}
	nv.x, err = item.GetPropertyAsFloatWithDefault("x", 0)
	if err != nil {
		return err
	}
	nv.y, err = item.GetPropertyAsFloatWithDefault("y", 0)
	if err != nil {
		return err
	}
	nv.width, err = item.GetPropertyAsFloatWithDefault("width", 0)
	if err != nil {
		return err
	}
	nv.height, err = item.GetPropertyAsFloatWithDefault("height", 0)
	if err != nil {
		return err
	}
	return nil
}

type NodeParagraph struct {
	children   []Node
	x          float64
	y          float64
	width      float64
	height     float64
	lineHeight float64
}

func (np *NodeParagraph) Children() []Node { return np.children }
func (np *NodeParagraph) addChild(child Node) error {
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
	np.children = append(np.children, child)
	return nil
}
func (np *NodeParagraph) initFrom(item *ast.Item) error {
	var err error
	np.x, err = item.GetPropertyAsFloatWithDefault("x", 0)
	if err != nil {
		return err
	}
	np.y, err = item.GetPropertyAsFloatWithDefault("y", 0)
	if err != nil {
		return err
	}
	np.width, err = item.GetPropertyAsFloatWithDefault("width", 0)
	if err != nil {
		return err
	}
	np.height, err = item.GetPropertyAsFloatWithDefault("height", 0)
	if err != nil {
		return err
	}
	np.lineHeight, err = item.GetPropertyAsFloatWithDefault("lineHeight", 6)
	if err != nil {
		return err
	}
	return nil
}

type NodeContainer struct {
	children []Node
	x        float64
	y        float64
}

func (nc *NodeContainer) Children() []Node { return nc.children }
func (nc *NodeContainer) addChild(child Node) error {
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
	nc.children = append(nc.children, child)
	return nil
}
func (nc *NodeContainer) initFrom(item *ast.Item) error {
	var err error
	nc.x, err = item.GetPropertyAsFloatWithDefault("x", 0)
	if err != nil {
		return err
	}
	nc.y, err = item.GetPropertyAsFloatWithDefault("y", 0)
	if err != nil {
		return err
	}
	return nil
}
