package svg

import (
	"math"
)

func circle(element *Element, worldToParent matrix) (*svgNode, error) {

	sn := &svgNode{
		worldToLocal: worldToParent,
		commands:     []command{},
	}

	cx, err := readAttributeAsFloat(element, "cx")
	cy, err := readAttributeAsFloat(element, "cy")
	r, err := readAttributeAsFloat(element, "r")
	if err != nil {
		return nil, err
	}

	newOriginX, newOriginY, _ := sn.worldToLocal.multiplyPoint(cx, cy, 1.0)
	newRadiusX, _, _ := sn.worldToLocal.multiplyPoint(r, 0, 1.0)

	cx = newOriginX
	cy = newOriginY
	r = newRadiusX

	arc := 4 / 3 * (math.Sqrt2 - 1) * r

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
