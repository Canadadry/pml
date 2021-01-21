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
