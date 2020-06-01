package svg

import (
	"github.com/jung-kurt/gofpdf"
)

type svgGroup struct {
	worldToLocal matrix
}

func group(element *Element, transform matrix) object {
	return &svgGroup{
		worldToLocal: transform,
	}

}

func (sg *svgGroup) draw(pdf *gofpdf.Fpdf) error {
	return nil
}
