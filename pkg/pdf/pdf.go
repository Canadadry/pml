package pdf

import (
	"fmt"
	"github.com/canadadry/pml/compiler/renderer"
	"github.com/canadadry/pml/pkg/svg/svgdrawer"
	"github.com/jung-kurt/gofpdf"
	"image/color"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const ptToMm = 25.4 / 72.0

var alignPossibleValue = map[renderer.PdfTextAlign]string{
	renderer.AlingTopLeft:        "TL",
	renderer.AlingTopCenter:      "TC",
	renderer.AlingTopRight:       "TR",
	renderer.AlingMiddleLeft:     "LM",
	renderer.AlingMiddleCenter:   "CM",
	renderer.AlingMiddleRight:    "RM",
	renderer.AlingBottomLeft:     "BL",
	renderer.AlingBottomCenter:   "BC",
	renderer.AlingBottomRight:    "BR",
	renderer.AlingBaselineLeft:   "AL",
	renderer.AlingBaselineCenter: "AC",
	renderer.AlingBaselineRight:  "AR",
}

type pdfbuilder struct {
	drawSvg svgdrawer.DrawFunc
}

type pdf struct {
	gopdf      *gofpdf.Fpdf
	drawSvg    svgdrawer.DrawFunc
	imageCount int
}

func New(d svgdrawer.DrawFunc) renderer.Pdf {
	return &pdfbuilder{
		drawSvg: d,
	}
}

func (pb *pdfbuilder) Init() renderer.PdfDrawer {
	return &pdf{
		gopdf:      gofpdf.New("P", "mm", "A4", ""),
		drawSvg:    pb.drawSvg,
		imageCount: 0,
	}
}

func (p *pdf) AddPage() {
	p.gopdf.AddPage()
}

func (p *pdf) SetFillColor(c color.RGBA) {
	p.gopdf.SetFillColor(int(c.R), int(c.G), int(c.B))
}

func (p *pdf) Rect(x float64, y float64, width float64, height float64) {
	p.gopdf.Rect(x, y, width, height, "F")
}

func (p *pdf) LoadFont(fontName string, fontFilePath string) error {
	dir := filepath.Dir(fontFilePath)
	base := filepath.Base(fontFilePath)
	namePart := strings.Split(base, ".")
	name := strings.Join(namePart[:len(namePart)-1], ".")

	if !fileExists(dir + "/" + name + ".json") {
		err := gofpdf.MakeFont(fontFilePath, dir+"/cp1258.map", dir, os.Stdout, true)
		if err != nil {
			return err
		}
	}
	p.gopdf.AddUTF8Font(fontName, "", fontFilePath)
	return nil
}

func (p *pdf) SetFont(fontName string, fontSizeMm float64) {

	fontSizePt := fontSizeMm / ptToMm
	p.gopdf.SetFont(fontName, "", fontSizePt)
}

func (p *pdf) GetDefaultFontName() string {
	return "Arial"
}

func (p *pdf) SetTextColor(c color.RGBA) {
	p.gopdf.SetTextColor(int(c.R), int(c.G), int(c.B))
}

func (p *pdf) Text(text string, x float64, y float64, width float64, height float64, align renderer.PdfTextAlign) {
	gopdfAlign, ok := alignPossibleValue[align]
	if !ok {
		gopdfAlign = "TL"
	}
	p.gopdf.SetXY(x, y)
	p.gopdf.CellFormat(width, height, text, "", 0, gopdfAlign, false, 0, "")
}

func (p *pdf) GetStringWidth(text string) float64 {
	return p.gopdf.GetStringWidth(text)
}

func (p *pdf) Image(image io.ReadSeeker, x float64, y float64, width float64, height float64) {

	mimetype := getFileContentType(image)
	imageType := p.gopdf.ImageTypeFromMime(mimetype)

	image.Seek(0, 0)

	imageRef := fmt.Sprintf("image%d", p.imageCount)
	p.imageCount = p.imageCount + 1

	_ = p.gopdf.RegisterImageOptionsReader(imageRef, gofpdf.ImageOptions{ImageType: imageType}, image)

	p.gopdf.ImageOptions(
		imageRef,
		x,
		y,
		width,
		height,
		false,
		gofpdf.ImageOptions{},
		0,
		"",
	)
}

func (p *pdf) Vector(vector io.Reader, x float64, y float64, width float64, height float64) {
	p.drawSvg(p, vector, x, y, width, height)
}

func (p *pdf) Output(out io.Writer) error {
	return p.gopdf.Output(out)
}

func (p *pdf) MoveTo(x float64, y float64) {
	p.gopdf.MoveTo(x, y)
}

func (p *pdf) LineTo(x float64, y float64) {
	p.gopdf.LineTo(x, y)
}

func (p *pdf) BezierTo(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) {

	p.gopdf.CurveBezierCubicTo(x1, y1, x2, y2, x3, y3)
}

func (p *pdf) CloseAndDraw(s svgdrawer.Style) {
	p.gopdf.ClosePath()

	p.gopdf.SetDrawColor(int(s.BorderColor.R), int(s.BorderColor.G), int(s.BorderColor.B))
	p.gopdf.SetFillColor(int(s.FillColor.R), int(s.FillColor.G), int(s.FillColor.B))
	p.gopdf.SetLineWidth(s.BorderSize)

	mode := ""

	if s.Fill {
		mode += "F"
	}
	if s.BorderSize > 0 {
		mode += "D"
	}

	p.gopdf.DrawPath(mode)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getFileContentType(file io.Reader) string {
	buf := make([]byte, 512, 512)
	file.Read(buf)
	contentType := http.DetectContentType(buf)
	return contentType
}
