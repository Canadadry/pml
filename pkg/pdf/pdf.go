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
	p := gofpdf.New("P", "mm", "A4", "")
	p.SetAutoPageBreak(false, 0)
	return &pdf{
		gopdf:      p,
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

func (p *pdf) SetStrokeColor(c color.RGBA) {
	p.gopdf.SetDrawColor(int(c.R), int(c.G), int(c.B))
}

func (p *pdf) SetStrokeWidth(w float64) {
	p.gopdf.SetLineWidth(w)
}

func (p *pdf) Rect(x float64, y float64, width float64, height float64, radius float64) {
	mode := "F"
	if p.gopdf.GetLineWidth() > 0 {
		mode = "B"
	}
	if radius <= 0 {
		p.gopdf.Rect(x, y, width, height, mode)
		return
	}

	xs := []float64{x, x + radius, x + width - radius, x + width}
	ys := []float64{y, y + radius, y + height - radius, y + height}

	p.gopdf.MoveTo(xs[1], ys[0])

	p.gopdf.LineTo(xs[2], ys[0])
	p.gopdf.ArcTo(xs[2], ys[1], radius, radius, -90, 180, 90)

	p.gopdf.LineTo(xs[3], ys[2])
	p.gopdf.ArcTo(xs[2], ys[2], radius, radius, -90, 90, 0)

	p.gopdf.LineTo(xs[1], ys[3])
	p.gopdf.ArcTo(xs[1], ys[2], radius, radius, -90, 0, -90)

	p.gopdf.LineTo(xs[0], ys[1])
	p.gopdf.ArcTo(xs[1], ys[1], radius, radius, -90, -90, -180)

	p.gopdf.ClosePath()
	p.gopdf.DrawPath(mode)
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

	switch s.PathStyle {
	case svgdrawer.None:
		return
	case svgdrawer.Fill:
		mode = "f"
	case svgdrawer.Stroke:
		mode = "S"
	default:
		mode = "B"
	}

	if s.EvenOddRule {
		mode += "*"
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
