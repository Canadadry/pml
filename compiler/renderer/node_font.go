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
