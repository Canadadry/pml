package paragraph

type Align string

const (
	AlignLeft    Align = "left"
	AlignRight         = "right"
	AlignCenter        = "center"
	AlignJustify       = "justify"
)

func (a Align) GetDetailForLine(l Line) (float64, float64) {
	switch a {
	case AlignRight:
		return l.SpaceLeft, 0
	case AlignCenter:
		return l.SpaceLeft / 2, 0
	case AlignJustify:
		if l.Overflow {
			return 0, l.SpaceLeft / float64(len(l.Words)-1)
		}
	}
	return 0, 0
}
