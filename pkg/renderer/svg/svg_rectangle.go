package svg

import (
	"fmt"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
)

func svgRectangle(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {

	sn := &svgNode{
		worldToLocal: worldToParent,
		commands:     []command{},
	}

	x, err := element.ReadAttributeAsFloat("x")
	if err != nil {
		return nil, fmt.Errorf("error while reading rectangle arrtibute x :%w", err)
	}
	y, err := element.ReadAttributeAsFloat("y")
	if err != nil {
		return nil, fmt.Errorf("error while reading rectangle arrtibute y :%w", err)
	}
	w, err := element.ReadAttributeAsFloat("width")
	if err != nil {
		return nil, fmt.Errorf("error while reading rectangle arrtibute width :%w", err)
	}
	h, err := element.ReadAttributeAsFloat("height")
	if err != nil {
		return nil, fmt.Errorf("error while reading rectangle arrtibute height :%w", err)
	}

	left, top, _ := sn.worldToLocal.Project(x, y, 1.0)
	right, bottom, _ := sn.worldToLocal.Project(x+w, y+h, 1.0)

	sn.commands = append(sn.commands,
		command{'M', left, top, 0, 0, 0, 0},
		command{'L', right, top, 0, 0, 0, 0},
		command{'L', right, bottom, 0, 0, 0, 0},
		command{'L', left, bottom, 0, 0, 0, 0},
		command{'L', left, top, 0, 0, 0, 0},
		command{'Z', 0, 0, 0, 0, 0, 0},
	)
	return sn, nil
}
