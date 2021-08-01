package paragraph

import (
	"image/color"
	"strings"
)

type TextSizer interface {
	GetStringWidth(str string, fontName string, fontSize float64) float64
}

func Format(in []Block, maxWidth float64, sizer TextSizer) []Line {
	return wordsToLines(blocsToWords(in), maxWidth, sizer)
}

type Block struct {
	Text     string
	FontSize float64
	FontName string
	Color    color.RGBA
}

type Word struct {
	Text     string
	Width    float64
	FontSize float64
	FontName string
	Color    color.RGBA
}

func blocsToWords(blocs []Block) []Word {
	words := []Word{}
	for _, b := range blocs {
		lines := strings.Split(b.Text, "\n")
		for i, l := range lines {
			w := strings.Split(l, " ")
			for _, s := range w {
				if len(s) == 0 {
					continue
				}
				words = append(words, Word{
					Text:     cleanText(s),
					FontSize: b.FontSize,
					FontName: b.FontName,
					Color:    b.Color,
				})
			}
			if i < (len(lines) - 1) {
				words = append(words, Word{Text: "\n"})
			}
		}
	}
	return words
}

func cleanText(in string) string {
	return in
}

type Line struct {
	Words          []Word
	StartingOffset float64
	MaxWidth       float64
	Overflow       bool
}

type lines struct {
	l         []Line
	w         float64
	lastLineX float64
}

func newLines(maxWidth float64) lines {
	return lines{
		l: []Line{{MaxWidth: maxWidth}},
		w: maxWidth,
	}
}

func (l *lines) AddWordInLastLine(w Word, sizer TextSizer) {
	w.Width = sizer.GetStringWidth(w.Text, w.FontName, w.FontSize)
	spacing := sizer.GetStringWidth(" ", w.FontName, w.FontSize)

	if l.lastLineX+w.Width > l.w {
		l.l[len(l.l)-1].Overflow = true
		l.BreakLine()
	}

	l.lastLineX += w.Width + spacing

	l.l[len(l.l)-1].Words = append(l.l[len(l.l)-1].Words, w)
}

func (l *lines) BreakLine() {
	l.lastLineX = 0
	l.l = append(l.l, Line{MaxWidth: l.w})
}

func wordsToLines(words []Word, maxWidth float64, sizer TextSizer) []Line {
	l := newLines(maxWidth)
	for _, w := range words {
		if w.Text == "\n" {
			l.BreakLine()
			continue
		}
		l.AddWordInLastLine(w, sizer)
	}
	return l.l
}
