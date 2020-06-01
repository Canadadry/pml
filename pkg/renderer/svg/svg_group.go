package svg

import (
	"github.com/jung-kurt/gofpdf"
	"strconv"
)

type svgGroup struct {
	worldToLocal matrix
	children     []object
}

func group(element *Element, worldToParent matrix) object {
	sg = &svgGroup{
		worldToLocal: worldToParent,
		children:     []object{},
	}

	transformAttr, ok := element.Attributes["transform"]
	if ok {
		transformMatrix := matrixFromGAttributes(transformAttr)
		sg.worldToLocal = worldToParent.multiplyMatrix(transformMatrix)
	}

	for _, child := range element.Children {
		switch child.Name {
		case "g":
			sg.children = append(sg.children, group(child, sg.worldToLocal))
		}
	}

}

func (sg *svgGroup) draw(pdf *gofpdf.Fpdf) error {
	for _, child := range sg.children {
		err := child.draw(pdf)
		if err != nil {
			return err
		}
	}
}

func matrixFromGAttributes(transformAttr string) matrix {

	if transform[:6] != "matrix" {
		return identityMatrix()
	}

	param := transform[7 : len(transform)-1]
	coef = strings.Split(param, ",")
	if len(coef) == 6 {
		fmt.Printf("group transform matrix (%s), dont have 6 part dont know what to do\n", transform)
		return identityMatrix()
	}

	out = identityMatrix()
	out.n11, err = strconv.ParseFloat(coef[0], 64)
	out.n12, err = strconv.ParseFloat(coef[1], 64)
	out.n13, err = strconv.ParseFloat(coef[2], 64)
	out.n21, err = strconv.ParseFloat(coef[3], 64)
	out.n22, err = strconv.ParseFloat(coef[4], 64)
	out.n23, err = strconv.ParseFloat(coef[5], 64)

	if err != nil {
		return identityMatrix()
	}
	return out
}
