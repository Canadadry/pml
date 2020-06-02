package svg

func rectangle(element *Element, worldToParent matrix) (*svgNode, error) {

	sn := &svgNode{
		worldToLocal: worldToParent,
		commands:     []command{},
	}

	x, err := readAttributeAsFloat(element, "x")
	y, err := readAttributeAsFloat(element, "y")
	w, err := readAttributeAsFloat(element, "width")
	h, err := readAttributeAsFloat(element, "height")
	if err != nil {
		return nil, err
	}

	left, top, _ := sn.worldToLocal.multiplyPoint(x, y, 1.0)
	right, bottom, _ := sn.worldToLocal.multiplyPoint(x+w, y+h, 1.0)

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
