package renderer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/canadadry/pml/compiler/ast"
	"io"
	"os"
)

const (
	ImgModeFile = "file"
	ImgModeB64  = "b64"
)

type NodeImage struct {
	Frame
	file string
	mode string
}

func (n *NodeImage) Children() []Node          { return nil }
func (n *NodeImage) addChild(child Node) error { return errChildrenNotAllowed }
func (n *NodeImage) needToDrawChild() bool     { return true }
func (*NodeImage) new(item *ast.Item) (Node, error) {
	n := &NodeImage{}
	var err error
	n.file, err = item.GetPropertyAsStringWithDefault("file", "")
	if err != nil {
		return nil, err
	}
	values := []string{ImgModeFile, ImgModeB64}
	n.mode, err = item.GetPropertyAsIdentifierFromListWithDefault("mode", ImgModeFile, values)
	if err != nil {
		return nil, err
	}
	err = n.Frame.initFrom(item)
	return n, err
}
func (n *NodeImage) draw(pdf PdfDrawer, rb renderBox) (renderBox, error) {
	if len(n.file) == 0 {
		return rb, ErrEmptyImageFileProperty
	}
	var rs io.ReadSeeker
	if n.mode == ImgModeFile {
		file, err := os.Open(n.file)
		if err != nil {
			return rb, fmt.Errorf("%w : %v", ErrCantOpenFile, err)
		}
		defer file.Close()
		rs = file
	} else {
		decoded, err := base64.StdEncoding.DecodeString(n.file)
		if err != nil {
			return rb, fmt.Errorf("%w : %v", ErrB64Read, err)
		}
		rs = bytes.NewReader(decoded)
	}
	rb = rb.Cut(n.Frame)
	pdf.Image(rs, rb.x, rb.y, rb.w, rb.h)
	return rb, nil
}
