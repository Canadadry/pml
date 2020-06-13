package renderer

import (
	"fmt"
	"github.com/canadadry/pml/pkg/abstract/abstractsvg"
	"github.com/canadadry/pml/pkg/domain/ast"
	"github.com/jung-kurt/gofpdf"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const ptToMm = 25.4 / 72.0

var alignPossibleValue = map[string]string{
	alingTopLeft:        "TL",
	alingTopCenter:      "TC",
	alingTopRight:       "TR",
	alingMiddleLeft:     "LM",
	alingMiddleCenter:   "CM",
	alingMiddleRight:    "RM",
	alingBottomLeft:     "BL",
	alingBottomCenter:   "BC",
	alingBottomRight:    "BR",
	alingBaselineLeft:   "AL",
	alingBaselineCenter: "AC",
	alingBaselineRight:  "AR",
}

type renderer struct {
	output      io.Writer
	svgRenderer abstractsvg.Svg
}

func New(output io.Writer, svg abstractsvg.Svg) renderer {
	return renderer{
		output:      output,
		svgRenderer: svg,
	}
}

func (r *renderer) Render(tree *ast.Item) error {

	rt, err := GenerateFrom(tree)
	if err != nil {
		return err
	}
	return r.draw(rt, nil)
}

func (r *renderer) draw(node Node, pdf *gofpdf.Fpdf) error {

	initialized := false
	renderChild := true

	switch n := node.(type) {
	case *NodeDocument:
		initialized = true
		pdf = gofpdf.New("P", "mm", "A4", "")
	case *NodePage:
		pdf.AddPage()
	case *NodeRectangle:
		pdf.SetFillColor(int(n.color.R), int(n.color.G), int(n.color.B))
		pdf.Rect(n.x, n.y, n.width, n.height, "F")
	case *NodeText:
		align, ok := alignPossibleValue[n.align]
		if !ok {
			return fmt.Errorf("%s is not a valid value for align property of text", n.align)
		}
		if len(n.fontName) == 0 {
			n.fontName = "Arial"
		}
		fontSizePt := n.fontSize / ptToMm
		pdf.SetFont(n.fontName, "", fontSizePt)
		pdf.SetTextColor(int(n.color.R), int(n.color.G), int(n.color.B))
		pdf.SetXY(n.x, n.y)
		pdf.CellFormat(n.width, n.height, n.text, "", 0, align, false, 0, "")
	case *NodeFont:
		dir := filepath.Dir(n.file)
		base := filepath.Base(n.file)
		namePart := strings.Split(base, ".")
		name := strings.Join(namePart[:len(namePart)-1], ".")

		if !fileExists(dir + "/" + name + ".json") {
			err := gofpdf.MakeFont(n.file, dir+"/cp1258.map", dir, os.Stdout, true)
			if err != nil {
				return err
			}
		}
		pdf.AddUTF8Font(n.name, "", n.file)
	case *NodeImage:

		if len(n.file) == 0 {
			return fmt.Errorf("in image item, you must specify a property file")
		}
		pdf.ImageOptions(
			n.file,
			n.x,
			n.y,
			n.width,
			n.height,
			false,
			gofpdf.ImageOptions{},
			0,
			"",
		)
	case *NodeVector:
		if len(n.file) == 0 {
			return fmt.Errorf("in vector item, you must specify a property file")
		}
		svgFile, err := os.Open(n.file)
		if err != nil {
			return err
		}
		defer svgFile.Close()

		r.svgRenderer.Draw(NewSvgToPdf(pdf), svgFile, n.x, n.y, n.width, n.height)
	case *NodeParagraph:
		renderChild = false
		x := 0.0
		y := 0.0

		for _, child := range node.Chilrend() {
			offset := 0
			textChild, ok := child.(*NodeText)
			if !ok {
				return fmt.Errorf("Unexpected node in paragraph")
			}
			if len(textChild.fontName) == 0 {
				textChild.fontName = "Arial"
			}

			fontSizePt := textChild.fontSize / ptToMm
			pdf.SetFont(textChild.fontName, "", fontSizePt)
			pdf.SetTextColor(int(textChild.color.R), int(textChild.color.G), int(textChild.color.B))

			for offset < len(textChild.text) {
				maxSize, textWidth := getTextMaxLength(pdf, textChild.text[offset:], n.width-x)
				pdf.SetXY(n.x+x, n.y+y)
				text := textChild.text[offset : offset+maxSize]
				align := "AL"
				pdf.CellFormat(n.width, n.lineHeight, text, "", 0, align, false, 0, "")
				offset = offset + maxSize
				x = x + textWidth
				if x > n.width {
					x = 0
					y = y + n.lineHeight
				}
			}

		}

	default:
		return fmt.Errorf("cannot render node type")
	}

	if renderChild {
		for _, child := range node.Chilrend() {
			err := r.draw(child, pdf)
			if err != nil {
				return err
			}
		}
	}

	if initialized {
		return pdf.Output(r.output)
	}

	return nil
}

func getTextMaxLength(pdf *gofpdf.Fpdf, text string, maxWidth float64) (int, float64) {
	splitted := strings.Split(text, " ")
	tmp := ""
	textWidth := 0.0
	for _, part := range splitted {
		textWidth = pdf.GetStringWidth(tmp + part + " ")
		if textWidth > maxWidth {
			return len(tmp), textWidth
		}
		tmp = tmp + part + " "
	}
	return len(text), textWidth
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
