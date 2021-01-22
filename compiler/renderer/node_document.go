package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

type NodeDocument struct {
	children []Node
}

func (n *NodeDocument) Children() []Node { return n.children }
func (n *NodeDocument) addChild(child Node) error {
	n.children = append(n.children, child)
	return nil
}
func (n *NodeDocument) new(*ast.Item) (Node, error) { return &NodeDocument{}, nil }
func (n *NodeDocument) needToDrawChild() bool       { return true }
func (n *NodeDocument) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	return rb, nil
}
