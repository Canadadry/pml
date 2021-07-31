package paragraph

const (
	AlignLeft    = "left"
	AlignRight   = "right"
	AlignCenter  = "center"
	AlignJustify = "justify"
)

func Align(lines *Line, align string) {
	for _, l := range lines {
		ls := getLineSize(l)
		x, spacing := getOffsetStartAndSpacingalign(align, ls)
		l.WordSpacing += spacing
		l.StartingOffset = x
	}
}

type lineSize struct {
	spaceLeft float64
	wordCount int
	lineBreak bool
}

func getLineSize(l Line) lineSize {
	ls := lineSize{
		wordCount: len(l.words),
		lineBreak: l.LineBreak,
	}
	if ls.wordCount > 0 {
		width := 0
		for _, word := range l.Words {
			width += word.Width
		}
		width += (ls.wordCount - 1) * l.WordSpacing
		ls.spaceLeft = l.MaxWidth - width
	}
	return ls
}

func getOffsetStartAndSpacing(align string, ls lineSize) (float64, float64) {
	switch align {
	case AlignRight:
		return ls.spaceLeft, 0
	case AlignCenter:
		return ls.spaceLeft / 2, 0
	case AlignJustify:
		if ls.lineBreak {
			return 0, 0
		}
		return 0, ls.spaceLeft / float64(ls.wordCount-1)
	}
	return 0, 0
}
