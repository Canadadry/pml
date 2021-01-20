package svg

import (
	"github.com/canadadry/pml/pkg/svg/svgdrawer"
	"github.com/canadadry/pml/pkg/svg/svgnode"
	"github.com/canadadry/pml/pkg/svg/svgparser"
	"io"
)

func Draw(d svgdrawer.Drawer, svg io.Reader, x float64, y float64, w float64, h float64) error {

	element, err := svgparser.Parse(svg)
	if err != nil {
		return err
	}

	root, err := svgnode.SvgMain(element, svgnode.ViewBox{x, y, w, h})
	if err != nil {
		return err
	}

	return root.Draw(d)
}
