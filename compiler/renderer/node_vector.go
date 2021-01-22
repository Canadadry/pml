package renderer

import (
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"os"
)

type NodeVector struct {
	Frame
	file string
}

func (n *NodeVector) Children() []Node          { return nil }
func (n *NodeVector) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeVector) needToDrawChild() bool     { return true }
func (*NodeVector) new(item *ast.Item) (Node, error) {
	n := &NodeVector{}
	var err error
	n.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return nil, err
	}

	err = n.Frame.initFrom(item)
	return n, err
}
func (n *NodeVector) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	if len(n.file) == 0 {
		return rb, ErrEmptyImageFileProperty
	}
	file, err := os.Open(n.file)
	if err != nil {
		return rb, fmt.Errorf("%w : %v", ErrCantOpenFile, err)
	}
	defer file.Close()
	rb = rb.Cut(n.Frame)
	pdf.Vector(file, rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}
