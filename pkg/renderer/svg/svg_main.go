package svg

import (
	"fmt"
	"math"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"strconv"
	"strings"
)

type viewBox struct {
	x float64
	y float64
	w float64
	h float64
}

func svgMain(element *svgparser.Element, targetVb viewBox) (*svgNode, error) {
	epsilon := 1e-8

	fileVb, err := parseSvgAttribute(element)
	if err != nil {
		return nil, err
	}

	if math.Abs(targetVb.w) < epsilon {
		targetVb.w = fileVb.w
		targetVb.h = fileVb.h
	}
	if math.Abs(targetVb.h) < epsilon {
		targetVb.h = fileVb.h / fileVb.w * targetVb.w
	}

	transform := matrix.Identity().Translate(targetVb.x, targetVb.y).Scale(targetVb.w, targetVb.h)
	transform = transform.Translate(fileVb.x, fileVb.y).Scale(1/fileVb.w, 1/fileVb.h)

	root, err := svgGroup(element, transform)
	if err != nil {
		return nil, err
	}
	return root, nil
}

func parseSvgAttribute(element *svgparser.Element) (*viewBox, error) {

	vb := &viewBox{0, 0, 0, 0}

	var err, errWidth, errHeight error
	vb.w, errWidth = element.ReadAttributeAsFloat("width")
	vb.h, errHeight = element.ReadAttributeAsFloat("height")
	if errWidth != nil || errHeight != nil {
		vb.w = 1
		vb.h = 1
	}

	viewBoxAttr, ok := element.Attributes["viewBox"]
	if !ok {
		return vb, nil
	}

	viewBoxParam := strings.Split(viewBoxAttr, " ")

	if len(viewBoxParam) != 4 {
		return nil, fmt.Errorf("viewBox (%s), dont have 4 part dont know what to do", viewBoxAttr)
	}

	vb.x, err = strconv.ParseFloat(viewBoxParam[0], 64)
	if err != nil {
		return nil, err
	}
	vb.y, err = strconv.ParseFloat(viewBoxParam[1], 64)
	if err != nil {
		return nil, err
	}
	vb.w, err = strconv.ParseFloat(viewBoxParam[2], 64)
	if err != nil {
		return nil, err
	}
	vb.h, err = strconv.ParseFloat(viewBoxParam[3], 64)
	if err != nil {
		return nil, err
	}

	return vb, nil
}
