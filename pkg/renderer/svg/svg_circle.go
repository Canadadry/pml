package svg

import (
	"fmt"
	"math"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
)

func svgCircle(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {

	sn := &svgNode{
		worldToLocal: worldToParent,
		commands:     []command{},
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

	cx, cy, _ = sn.worldToLocal.Project(cx, cy, 1.0)
	r, _, _ = sn.worldToLocal.Project(r, 0, 1.0)

	// circle with bezier curve param
	arc := 4.0 / 3.0 * (math.Sqrt2 - 1) * r

	sn.commands = append(sn.commands,
		command{'M', cx, cy - r, 0, 0, 0, 0},
		command{'C', cx + arc, cy - r, cx + r, cy - arc, cx + r, cy},
		command{'C', cx + r, cy + arc, cx + arc, cy + r, cx, cy + r},
		command{'C', cx - arc, cy + r, cx - r, cy + arc, cx - r, cy},
		command{'C', cx - r, cy - arc, cx - arc, cy - r, cx, cy - r},
		command{'Z', 0, 0, 0, 0, 0, 0},
	)
	return sn, nil
}
