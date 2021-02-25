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

const (
	ImgDrawKeep = "keepAspectRatio"
	ImgDrawFill = "fill"
)

type NodeImage struct {
	Frame
	file     string
	srcMode  string
	drawMode string
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
	modeValues := []string{ImgModeFile, ImgModeB64}
	n.srcMode, err = item.GetPropertyAsIdentifierFromListWithDefault("mode", ImgModeFile, modeValues)
	if err != nil {
		return nil, err
	}
	drawValues := []string{ImgDrawKeep, ImgDrawFill}
	n.drawMode, err = item.GetPropertyAsIdentifierFromListWithDefault("draw", ImgDrawFill, drawValues)
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
	if n.srcMode == ImgModeFile {
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
	pdf.Image(rs, rb.x, rb.y, rb.w, rb.h, n.drawMode == ImgDrawKeep)
	return rb, nil
}
