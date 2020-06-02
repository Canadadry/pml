package svg

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	"strconv"
	"strings"
)

type object interface {
	draw(pdf *gofpdf.Fpdf) error
}

func Draw(pdf *gofpdf.Fpdf, svg string, x float64, y float64, w float64, h float64) error {

	content, err := os.Open(svg)
	if err != nil {
		return err
	}
	element, err := Parse(content)
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

	transform := identityMatrix().translate(x, y).scale(w, h)
	transform = transform.translate(viewBox.x, viewBox.y)
	transform = transform.scale(1/viewBox.w, 1/viewBox.h)

	root, err := group(element, transform)
	if err != nil {
		return err
	}

	return root.draw(pdf)
}
