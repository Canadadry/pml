package svg

import (
	"fmt"
	"math"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"pml/pkg/renderer/svg/svgpath"
)

func svgCircle(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {

	sn := &svgNode{
		worldToLocal: worldToParent,
		commands:     []svgpath.Command{},
		style:        parseStyleAttribute(element),
	}

	cx, err := element.ReadAttributeAsFloat("cx")
	if err != nil {
		return nil, fmt.Errorf("error while reading circle arrtibute cx :%w", err)
	}
	cy, err := element.ReadAttributeAsFloat("cy")
	if err != nil {
		return nil, fmt.Errorf("error while reading circle arrtibute cy :%w", err)
	}
	r, err := element.ReadAttributeAsFloat("r")
	if err != nil {
		return nil, fmt.Errorf("error while reading circle arrtibute r :%w", err)
	}

	// circle with bezier curve param
	arc := 4.0 / 3.0 * (math.Sqrt2 - 1) * r

	sn.commands = append(sn.commands,
		svgpath.Command{'M', []svgpath.Point{{cx, cy - r}}},
		svgpath.Command{'C', []svgpath.Point{{cx + arc, cy - r}, {cx + r, cy - arc}, {cx + r, cy}}},
		svgpath.Command{'C', []svgpath.Point{{cx + r, cy + arc}, {cx + arc, cy + r}, {cx, cy + r}}},
		svgpath.Command{'C', []svgpath.Point{{cx - arc, cy + r}, {cx - r, cy + arc}, {cx - r, cy}}},
		svgpath.Command{'C', []svgpath.Point{{cx - r, cy - arc}, {cx - arc, cy - r}, {cx, cy - r}}},
		svgpath.Command{'Z', []svgpath.Point{}},
	)
	return sn, nil
}
