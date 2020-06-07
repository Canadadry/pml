package svg

import (
	"fmt"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"pml/pkg/renderer/svg/svgpath"
)

func svgPath(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {
	sp := &svgNode{
		worldToLocal: worldToParent,
	}

	d, ok := element.Attributes["d"]
	if !ok {
		return nil, fmt.Errorf("No attribute 'd' found")
	}

	cmd, err := svgpath.Parse(d)
	if err != nil {
		return nil, err
	}

	sp.commands = cmd

	for i, c := range sp.commands {
		for j, pt := range c.Points {
			x, y, _ := worldToParent.Project(pt.X, pt.Y, 1)
			sp.commands[i].Points[j].X = x
			sp.commands[i].Points[j].Y = y
		}
	}

	return sp, nil
}
