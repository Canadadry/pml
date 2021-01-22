package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

type NodeFont struct {
	file string
	name string
}

func (n *NodeFont) Children() []Node          { return nil }
func (n *NodeFont) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeFont) needToDrawChild() bool     { return true }
func (*NodeFont) new(item *ast.Item) (Node, error) {
	n := &NodeFont{}
	var err error
	n.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return nil, err
	}
	n.name, err = item.GetPropertyAsStringWithDefault("name", "")
	if err != nil {
		return nil, err
	}
	return n, nil
}
func (n *NodeFont) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	pdf.LoadFont(n.name, n.file)
	return rb, nil
}
