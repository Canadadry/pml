package svg

import (
	"fmt"
	"github.com/canadadry/pml/pkg/renderer/svg/matrix"
	"github.com/canadadry/pml/pkg/renderer/svg/svgparser"
	"strconv"
	"strings"
)

type svgParserElement func(*svgparser.Element, matrix.Matrix) (*svgNode, error)

func svgGroup(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {
	var parserMap map[string]svgParserElement = map[string]svgParserElement{
		"g":      svgGroup,
		"path":   svgPath,
		"rect":   svgRectangle,
		"circle": svgCircle,
	}

	sg := &svgNode{
		worldToLocal: worldToParent,
		children:     []*svgNode{},
	}

	transformAttr, ok := element.Attributes["transform"]
	if ok {
		transformMatrix, err := matrixFromGAttributes(transformAttr)
		if err != nil {
			return nil, err
		}
		sg.worldToLocal = worldToParent.MultiplyBy(transformMatrix)
	}

	for _, child := range element.Children {
		parser, ok := parserMap[child.Name]
		if !ok {
			return nil, errCannotParseElement
		}

		child, err := parser(child, sg.worldToLocal)
		if err != nil {
			return nil, err
		}
		sg.children = append(sg.children, child)
	}
	return sg, nil
}

func matrixFromGAttributes(transformAttr string) (matrix.Matrix, error) {

	if len(transformAttr) < 7 {
		return matrix.Identity(), errCannotParseMainTransformAttr
	}

	if transformAttr[:6] != "matrix" {
		return matrix.Identity(), errCannotParseMainTransformAttr
	}

	param := transformAttr[7 : len(transformAttr)-1]
	coef := strings.Split(param, ",")
	if len(coef) != 6 {
		err := fmt.Errorf(
			"%w : group transform matrix (%s), dont have 6 part dont know what to do\n",
			errCannotParseMainTransformAttr,
			transformAttr,
		)
		return matrix.Identity(), err
	}

	n11, err := strconv.ParseFloat(coef[0], 64)
	n21, err := strconv.ParseFloat(coef[1], 64)
	n12, err := strconv.ParseFloat(coef[2], 64)
	n22, err := strconv.ParseFloat(coef[3], 64)
	n13, err := strconv.ParseFloat(coef[4], 64)
	n23, err := strconv.ParseFloat(coef[5], 64)

	if err != nil {
		return matrix.Identity(), fmt.Errorf("%w : %v", errCannotParseMainTransformAttr, err)
	}
	return matrix.New(
		n11, n12, n13,
		n21, n22, n23,
		0, 0, 1,
	), nil
}
