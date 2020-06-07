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

	left, top, _ := sn.worldToLocal.Project(x, y, 1.0)
	right, bottom, _ := sn.worldToLocal.Project(x+w, y+h, 1.0)

	sn.commands = append(sn.commands,
		svgpath.Command{'M', []svgpath.Point{{left, top}}},
		svgpath.Command{'L', []svgpath.Point{{right, top}}},
		svgpath.Command{'L', []svgpath.Point{{right, bottom}}},
		svgpath.Command{'L', []svgpath.Point{{left, bottom}}},
		svgpath.Command{'L', []svgpath.Point{{left, top}}},
		svgpath.Command{'Z', []svgpath.Point{}},
	)
	return sn, nil
}
