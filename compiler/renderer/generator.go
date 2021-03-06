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
	itemPath      = "Path"
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
		AllowedChild: []string{itemRectangle, itemText, itemImage, itemVector, itemParagraph, itemContainer, itemPath},
		Type:         &NodePage{},
	},
	itemContainer: {
		AllowedChild: []string{itemRectangle, itemText, itemImage, itemVector, itemParagraph, itemContainer, itemPath},
		Type:         &NodeContainer{},
	},
	itemRectangle: {
		AllowedChild: []string{itemRectangle, itemText, itemImage, itemVector, itemParagraph, itemPath},
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
	itemPath: {
		AllowedChild: []string{},
		Type:         &NodePath{},
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
			return nil, fmt.Errorf("%w : in %s cannot have %s", errChildrenNotAllowed, item.TokenType.Literal, c.TokenType.Literal)
		}
		err = n.addChild(child)
		if err != nil {
			return nil, err
		}
	}
	return n, nil
}
