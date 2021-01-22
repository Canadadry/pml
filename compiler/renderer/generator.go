package renderer

import (
	"errors"
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
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

type definition struct {
	AllowedChild []string
	Type         Node
}

func (d *definition) AllowChild(name string) bool {
	for _, c := range d.AllowedChild {
		if c == name {
			return true
		}
	}
	return false
}

var nodeDefinition = map[string]definition{
	itemDocument: {
		AllowedChild: []string{itemPage, itemFont},
		Type:         &NodeDocument{},
	},
	itemFont: {
		AllowedChild: []string{},
		Type:         &NodeFont{},
	},
	itemPage: {
		AllowedChild: []string{itemRectangle, itemText, itemImage, itemVector, itemParagraph, itemContainer},
		Type:         &NodePage{},
	},
	itemContainer: {
		AllowedChild: []string{itemRectangle, itemText, itemImage, itemVector, itemParagraph, itemContainer},
		Type:         &NodeContainer{},
	},
	itemRectangle: {
		AllowedChild: []string{itemRectangle, itemText, itemImage, itemVector, itemParagraph},
		Type:         &NodeRectangle{},
	},
	itemParagraph: {
		AllowedChild: []string{itemText},
		Type:         &NodeParagraph{},
	},
	itemText: {
		AllowedChild: []string{},
		Type:         &NodeText{},
	},
	itemImage: {
		AllowedChild: []string{},
		Type:         &NodeImage{},
	},
	itemVector: {
		AllowedChild: []string{},
		Type:         &NodeVector{},
	},
}

var (
	errItemNotFound        = errors.New("errItemNotFound")
	rootMustBeDocumentItem = errors.New("rootMustBeDocumentItem")
	errChildrenNotAllowed  = errors.New("errChildrenNotAllowed")
)

type Node interface {
	Children() []Node
	addChild(child Node) error
	new(*ast.Item) (Node, error)
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
	def, ok := nodeDefinition[item.TokenType.Literal]
	if !ok {
		return nil, errItemNotFound
	}
	n, err := def.Type.new(item)
	if err != nil {
		return nil, err
	}
	for _, c := range item.Children {
		child, err := generateItem(&c)
		if err != nil {
			return nil, err
		}
		if def.AllowChild(c.TokenType.Literal) == false {
			return nil, errChildrenNotAllowed
		}
		err = n.addChild(child)
		if err != nil {
			return nil, err
		}
	}
	return n, nil
}
