package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

type NodeContainer struct {
	Frame
	children []Node
}

func (n *NodeContainer) Children() []Node      { return n.children }
func (n *NodeContainer) needToDrawChild() bool { return true }
func (n *NodeContainer) addChild(child Node) error {
	n.children = append(n.children, child)
	return nil
}
func (*NodeContainer) new(item *ast.Item) (Node, error) {
	n := &NodeContainer{}
	err := n.Frame.initFrom(item)
	return n, err
}

func (n *NodeContainer) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	return rb.Cut(n.Frame), nil
}
