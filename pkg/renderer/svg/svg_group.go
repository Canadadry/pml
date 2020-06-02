package svg

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"strconv"
	"strings"
)

type svgGroup struct {
	worldToLocal matrix
	children     []object
}

func group(element *Element, worldToParent matrix) (object, error) {
	sg := &svgGroup{
		worldToLocal: worldToParent,
		children:     []object{},
	}

	transformAttr, ok := element.Attributes["transform"]
	if ok {
		transformMatrix, err := matrixFromGAttributes(transformAttr)
		if err != nil {
			return nil, err
		}
		sg.worldToLocal = worldToParent.multiplyMatrix(transformMatrix)
	}

	for _, child := range element.Children {
		switch child.Name {
		case "g":
			child, err := group(child, sg.worldToLocal)
			if err != nil {
				return nil, err
			}
			sg.children = append(sg.children, child)
		}
	}
	return sg, nil
}

func (sg *svgGroup) draw(pdf *gofpdf.Fpdf) error {
	for _, child := range sg.children {
		err := child.draw(pdf)
		if err != nil {
			return err
		}
	}
	return nil
}

func matrixFromGAttributes(transformAttr string) (matrix, error) {

	if len(transformAttr) < 7 {
		return identityMatrix(), errCannotParseMainTransformAttr
	}

	if transformAttr[:6] != "matrix" {
		return identityMatrix(), errCannotParseMainTransformAttr
	}

	param := transformAttr[7 : len(transformAttr)-1]
	coef := strings.Split(param, ",")
	if len(coef) != 6 {
		err := fmt.Errorf(
			"%w : group transform matrix (%s), dont have 6 part dont know what to do\n",
			errCannotParseMainTransformAttr,
			transformAttr,
		)
		return identityMatrix(), err
	}

	out := identityMatrix()
	var err error
	out.n11, err = strconv.ParseFloat(coef[0], 64)
	out.n12, err = strconv.ParseFloat(coef[1], 64)
	out.n13, err = strconv.ParseFloat(coef[2], 64)
	out.n21, err = strconv.ParseFloat(coef[3], 64)
	out.n22, err = strconv.ParseFloat(coef[4], 64)
	out.n23, err = strconv.ParseFloat(coef[5], 64)

	if err != nil {
		return identityMatrix(), fmt.Errorf("%w : %v", errCannotParseMainTransformAttr, err)
	}
	return out, nil
}
