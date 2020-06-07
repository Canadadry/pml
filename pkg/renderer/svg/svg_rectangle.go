package svg

import (
	"fmt"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"pml/pkg/renderer/svg/svgpath"
)

func svgRectangle(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {

	sn := &svgNode{
		worldToLocal: worldToParent,
		commands:     []svgpath.Command{},
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

	sn.commands = append(sn.commands,
		svgpath.Command{'M', []svgpath.Point{{x, y}}},
		svgpath.Command{'h', []svgpath.Point{{w, 0}}},
		svgpath.Command{'v', []svgpath.Point{{h, 0}}},
		svgpath.Command{'h', []svgpath.Point{{-w, 0}}},
		svgpath.Command{'v', []svgpath.Point{{-h, 0}}},
		svgpath.Command{'Z', []svgpath.Point{}},
	)
	return sn, nil
}
