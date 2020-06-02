package svg

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	"strconv"
	"strings"
)

type command struct {
	kind byte
	x1   float64
	y1   float64
	x2   float64
	y2   float64
	x3   float64
	y3   float64
}

type svgNode struct {
	worldToLocal matrix
	commands     []command
	children     []*svgNode
}

func (sn *svgNode) draw(pdf *gofpdf.Fpdf) error {

	for _, cmd := range sn.commands {
		fmt.Printf("drawing %s, %g, %g\n", string(cmd.kind), cmd.x1, cmd.y2)
		switch cmd.kind {
		case 'M':
			pdf.MoveTo(cmd.x1, cmd.y1)
		case 'L':
			pdf.LineTo(cmd.x1, cmd.y1)
		case 'C':
			pdf.LineTo(cmd.x1, cmd.y1)
		case 'Z':
			pdf.ClosePath()
			pdf.DrawPath("FD")
		}
	}

	for _, child := range sn.children {
		err := child.draw(pdf)
		if err != nil {
			return err
		}
	}
	return nil
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

func readAttributeAsFloat(element *Element, attribute string) (float64, error) {
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
