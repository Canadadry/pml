package svg

import (
	"fmt"
	"io"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
)

type Drawer interface {
	MoveTo(x float64, y float64)
	LineTo(x float64, y float64)
	ClosePath()
}

type command struct {
	kind byte
	x1   float64
	y1   float64
	x2   float64
	y2   float64
	x3   float64
	y3   float64
}

func (c command) ToString() string {
	return fmt.Sprintf("%s, 1: %g,%g 2: %g,%g 3: %g,%g", string(c.kind), c.x1, c.y1, c.x2, c.y2, c.x3, c.y3)
}

type svgNode struct {
	worldToLocal matrix.Matrix
	commands     []command
	children     []*svgNode
}

func (sn *svgNode) draw(d Drawer) error {

	for _, cmd := range sn.commands {
		fmt.Printf("drawing %s, %g, %g\n", string(cmd.kind), cmd.x1, cmd.y2)
		switch cmd.kind {
		case 'M':
			d.MoveTo(cmd.x1, cmd.y1)
		case 'L':
			d.LineTo(cmd.x1, cmd.y1)
		case 'C':
			d.LineTo(cmd.x1, cmd.y1)
		case 'Z':
			d.ClosePath()
		}
	}

	for _, child := range sn.children {
		err := child.draw(d)
		if err != nil {
			return err
		}
	}
	return nil
}

func Draw(d Drawer, svg io.Reader, x float64, y float64, w float64, h float64) error {

	element, err := svgparser.Parse(svg)
	if err != nil {
		return err
	}

	transform := matrix.Identity().Translate(x, y).Scale(w, h)
	root, err := svgMain(element, transform)
	if err != nil {
		return err
	}

	return root.draw(d)
}
