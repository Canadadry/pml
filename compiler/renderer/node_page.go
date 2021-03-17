package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

type NodePage struct {
	children []Node
}

func (n *NodePage) Children() []Node      { return n.children }
func (n *NodePage) needToDrawChild() bool { return true }
func (n *NodePage) addChild(child Node) error {
	n.children = append(n.children, child)
	return nil
}
func (n *NodePage) new(*ast.Item) (Node, error) { return &NodePage{}, nil }
func (n *NodePage) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.AddPage()
	return renderBox{w: PdfWidthMm, h: PdfHeight, s: 1.0}, nil
}
