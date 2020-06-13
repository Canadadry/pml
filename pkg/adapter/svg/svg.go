package svg

import (
	"github.com/canadadry/pml/pkg/abstract"
	"github.com/canadadry/pml/pkg/adapter/svg/matrix"
	"github.com/canadadry/pml/pkg/adapter/svg/svgparser"
	"github.com/canadadry/pml/pkg/adapter/svg/svgpath"
	"io"
)

type svgNode struct {
	worldToLocal matrix.Matrix
	commands     []svgpath.Command
	style        abstract.Style
	children     []*svgNode
}

type svgDrawer struct{}

func New() *svgDrawer {
	return &svgDrawer{}
}

func (sd *svgDrawer) Draw(d abstract.Drawer, svg io.Reader, x float64, y float64, w float64, h float64) error {

	element, err := svgparser.Parse(svg)
	if err != nil {
		return err
	}

	root, err := svgMain(element, viewBox{x, y, w, h})
	if err != nil {
		return err
	}

	return root.draw(d)
}
