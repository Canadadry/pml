package svg

import (
	"fmt"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"strconv"
	"strings"
)

func svgMain(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {

	worldToParent, err := applyViewBoxAttr(worldToParent, element)
	if err != nil {
		return nil, err
	}

	root, err := svgGroup(element, worldToParent)
	if err != nil {
		return nil, err
	}
	return root, nil
}

func applyViewBoxAttr(transform matrix.Matrix, element *svgparser.Element) (matrix.Matrix, error) {

	viewBoxAttr, ok := element.Attributes["viewBox"]
	if !ok {
		return transform, fmt.Errorf("cant find viewBox, dont know what to do")
	}

	viewBoxParam := strings.Split(viewBoxAttr, " ")

	if len(viewBoxParam) != 4 {
		return transform, fmt.Errorf("viewBox (%s), dont have 4 part dont know what to do", viewBoxAttr)
	}

	viewBox := struct {
		x float64
		y float64
		w float64
		h float64
	}{0, 0, 0, 0}

	var err error
	viewBox.x, err = strconv.ParseFloat(viewBoxParam[0], 64)
	viewBox.y, err = strconv.ParseFloat(viewBoxParam[1], 64)
	viewBox.w, err = strconv.ParseFloat(viewBoxParam[2], 64)
	viewBox.h, err = strconv.ParseFloat(viewBoxParam[3], 64)
	if err != nil {
		return transform, err
	}

	transform = transform.Translate(viewBox.x, viewBox.y)
	transform = transform.Scale(1/viewBox.w, 1/viewBox.h)

	return transform, nil
}
