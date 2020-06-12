package svg

import (
	"fmt"
	"github.com/canadadry/pml/pkg/renderer/svg/matrix"
	"github.com/canadadry/pml/pkg/renderer/svg/svgparser"
	"github.com/canadadry/pml/pkg/renderer/svg/svgpath"
)

func svgPath(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {
	sp := &svgNode{
		worldToLocal: worldToParent,
		style:        parseStyleAttribute(element, worldToParent),
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

	return sp, nil
}
