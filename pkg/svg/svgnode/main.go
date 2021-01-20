package svgnode

import (
	"errors"
	"fmt"
	"github.com/canadadry/pml/pkg/matrix"
	"github.com/canadadry/pml/pkg/svg/svgdrawer"
	"github.com/canadadry/pml/pkg/svg/svgparser"
	"github.com/canadadry/pml/pkg/svg/svgpath"
	"math"
	"strconv"
	"strings"
)

var (
	errCannotParseMainTransformAttr = errors.New("errCannotParseMainTransformAttr")
	errCannotParseElement           = errors.New("errCannotParseElement")
)

type SvgNode struct {
	worldToLocal matrix.Matrix
	commands     []svgpath.Command
	style        svgdrawer.Style
	children     []*SvgNode
}

type ViewBox struct {
	X float64
	Y float64
	W float64
	H float64
}

func SvgMain(element *svgparser.Element, targetVb ViewBox) (*SvgNode, error) {
	epsilon := 1e-8

	fileVb, err := parseSvgAttribute(element)
	if err != nil {
		return nil, err
	}

	if math.Abs(targetVb.W) < epsilon {
		targetVb.W = fileVb.W
		targetVb.H = fileVb.H
	}
	if math.Abs(targetVb.H) < epsilon {
		targetVb.H = fileVb.H / fileVb.W * targetVb.W
	}

	transform := matrix.Identity().Translate(targetVb.X, targetVb.Y).Scale(targetVb.W, targetVb.H)
	transform = transform.Translate(fileVb.X, fileVb.Y).Scale(1/fileVb.W, 1/fileVb.H)

	root, err := svgGroup(element, transform)
	if err != nil {
		return nil, err
	}
	return root, nil
}

func parseSvgAttribute(element *svgparser.Element) (*ViewBox, error) {

	vb := &ViewBox{0, 0, 0, 0}

	var err, errWidth, errHeight error
	vb.W, errWidth = element.ReadAttributeAsFloat("width")
	vb.H, errHeight = element.ReadAttributeAsFloat("height")
	if errWidth != nil || errHeight != nil {
		vb.W = 1
		vb.H = 1
	}

	viewBoxAttr, ok := element.Attributes["viewBox"]
	if !ok {
		return vb, nil
	}

	viewBoxParam := strings.Split(viewBoxAttr, " ")

	if len(viewBoxParam) != 4 {
		return nil, fmt.Errorf("viewBox (%s), dont have 4 part dont know what to do", viewBoxAttr)
	}

	vb.X, err = strconv.ParseFloat(viewBoxParam[0], 64)
	if err != nil {
		return nil, err
	}
	vb.Y, err = strconv.ParseFloat(viewBoxParam[1], 64)
	if err != nil {
		return nil, err
	}
	vb.W, err = strconv.ParseFloat(viewBoxParam[2], 64)
	if err != nil {
		return nil, err
	}
	vb.H, err = strconv.ParseFloat(viewBoxParam[3], 64)
	if err != nil {
		return nil, err
	}

	return vb, nil
}
