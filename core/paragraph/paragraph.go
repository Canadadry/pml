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
	WordSpacing    float64
	StartingOffset float64
	MaxWidth       float64
	LineWasBreaked bool
}

func wordsToLines(words []Word, maxWidth float64, sizer TextSizer) []Line {
	x := 0.0
	lines := []Line{}
	line := Line{}
	for _, w := range words {
		if (x+w.Width) > maxWidth || w.Text == "\n" {
			if x == 0 && w.Text != "\n" {
				continue
			}
			line.LineWasBreaked = w.Text == "\n"
			lines = append(lines, line)
			line = Line{MaxWidth: maxWidth}
			x = 0
		}
		x = x + w.Width + sizer.GetStringWidth(" ", w.FontName, w.FontSize)
		if w.Text != "\n" {
			line.Words = append(line.Words, w)
		}
	}
	if len(line.Words) > 0 {
		line.LineWasBreaked = true
		lines = append(lines, line)
	}
	return lines
}
