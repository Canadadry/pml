package paragraph

import (
	"image/color"
	"strings"
)

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

func BlocsToWords(blocs []Bloc) []Word {
	words := []Word{}
	for _, b := range r {
		splitted := strings.Split(node.text, "\n")
		for _, s := range splitted {
			words = append(words, Word{
				Text:     cleanText(s),
				FontSize: b.fontSize,
				FontName: b.fontName,
				Color:    b.color,
			})
		}
	}
}

type Line struct {
	Words          []Word
	WordSpacing    float64
	StartingOffset float64
	MaxWidth       float64
}

type TextSizer interface {
	GetStringWidth(str string, fontName string, font float64) float64
}

func WordsToLines(words []Word, maxWidth float64, sizer TextSizer) []Line {
	x := 0.0
	lines := []Line{}
	line := Line{}
	for _, w := range words {
		if (x+w.width) > width || w.text == "\n" {
			if x == 0 && w.text != "\n" {
				continue
			}
			line.lineWasBreaked = w.text == "\n"
			lines = append(lines, line)
			line = Line{}
			x = 0
		}
		x = x + w.width + w.spaceWidth
		if w.text != "\n" {
			line.words = append(line.words, w)
		}
	}
	if len(line.words) > 0 {
		line.lineWasBreaked = true
		lines = append(lines, line)
	}
	return lines
}
