package renderer

import (
	"github.com/canadadry/pml/compiler/ast"
)

type Frame struct {
	x      float64
	y      float64
	width  float64
	height float64
	xAlign string
	yAlign string
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
	xvalues := []string{Left, Center, Right, Fill, Layout, Free}
	f.xAlign, err = item.GetPropertyAsIdentifierFromListWithDefault("xAlign", xvalues[5], xvalues)
	if err != nil {
		return err
	}
	yvalues := []string{Top, Middle, Bottom, Fill, Layout, Free}
	f.yAlign, err = item.GetPropertyAsIdentifierFromListWithDefault("yAlign", yvalues[5], yvalues)
	if err != nil {
		return err
	}
	return nil
}
