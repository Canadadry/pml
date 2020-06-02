package svg

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"strconv"
	"strings"
)

func Draw(pdf *gofpdf.Fpdf, svg string, x float64, y float64, w float64, h float64) error {

	content, err := os.Open(svg)
	if err != nil {
		return err
	}
	element, err := svgparser.Parse(content)
	if err != nil {
		return err
	}

	viewBoxAttr, ok := element.Attributes["viewBox"]
	if !ok {
		return fmt.Errorf("cant find viewBox, dont know what to do")
	}

	viewBoxParam := strings.Split(viewBoxAttr, " ")

	if len(viewBoxParam) != 4 {
		return fmt.Errorf("viewBox (%s), dont have 4 part dont know what to do", viewBoxAttr)
	}

	viewBox := struct {
		x float64
		y float64
		w float64
		h float64
	}{0, 0, 0, 0}

	viewBox.x, err = strconv.ParseFloat(viewBoxParam[0], 64)
	viewBox.y, err = strconv.ParseFloat(viewBoxParam[1], 64)
	viewBox.w, err = strconv.ParseFloat(viewBoxParam[2], 64)
	viewBox.h, err = strconv.ParseFloat(viewBoxParam[3], 64)
	if err != nil {
		return err
	}

	transform := matrix.Identity().Translate(x, y).Scale(w, h)
	transform = transform.Translate(viewBox.x, viewBox.y)
	transform = transform.Scale(1/viewBox.w, 1/viewBox.h)

	root, err := group(element, transform)
	if err != nil {
		return err
	}

	return root.draw(pdf)
}

func readAttributeAsFloat(element *svgparser.Element, attribute string) (float64, error) {
	attr, ok := element.Attributes[attribute]
	if !ok {
		return 0, errMissingAttr
	}

	parsed, err := strconv.ParseFloat(attr, 64)
	if err != nil {
		return 0, err
	}

	return parsed, nil
}
