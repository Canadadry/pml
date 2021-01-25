package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

const (
	Left     = "left"
	Right    = "right"
	Top      = "top"
	Bottom   = "bottom"
	Center   = "center"
	Fill     = "fill"
	Layout   = "layout"
	Relative = "relative"
)

type Frame struct {
	x       float64
	y       float64
	width   float64
	height  float64
	xAnchor string
	yAnchor string
}

func (f *Frame) initFrom(item *ast.Item) error {
	var err error
	f.x, err = item.GetPropertyAsFloatWithDefault("x", 0)
	if err != nil {
		return err
	}
	f.y, err = item.GetPropertyAsFloatWithDefault("y", 0)
	if err != nil {
		return err
	}
	f.width, err = item.GetPropertyAsFloatWithDefault("width", 0)
	if err != nil {
		return err
	}
	f.height, err = item.GetPropertyAsFloatWithDefault("height", 0)
	if err != nil {
		return err
	}
	xvalues := []string{Left, Center, Right, Fill, Layout, Relative}
	f.xAnchor, err = item.GetPropertyAsIdentifierFromListWithDefault("xAnchor", xvalues[5], xvalues)
	if err != nil {
		return err
	}
	yvalues := []string{Top, Center, Bottom, Fill, Layout, Relative}
	f.yAnchor, err = item.GetPropertyAsIdentifierFromListWithDefault("yAnchor", yvalues[5], yvalues)
	if err != nil {
		return err
	}

	values := []string{Center, Fill, Relative}
	anchor, err := item.GetPropertyAsIdentifierFromListWithDefault("anchor", values[2], values)
	if err != nil {
		return err
	}
	switch anchor {
	case Center:
		f.xAnchor = Center
		f.yAnchor = Center
	case Fill:
		f.xAnchor = Fill
		f.yAnchor = Fill
	}

	return nil
}
