package flowable

type Area struct {
	left, right, top, bottom float64
	spacingV, spacingH       float64
	x, y                     float64
	lineHeight               float64
}

func New(x, y, width, height, spacingH, spacingV float64) Area {
	return Area{
		left:     x,
		right:    x + width,
		top:      y,
		bottom:   y + height,
		spacingH: spacingH,
		spacingV: spacingV,
	}
}

func (a Area) GetX() float64 {
	return a.x
}

func (a Area) GetY() float64 {
	return a.y
}

func (a *Area) Add(width, height float64) bool {
	if a.x+width > a.right {
		if a.y+a.lineHeight+height > a.bottom {
			return false
		}
		a.y += a.lineHeight + a.spacingV
		a.x = a.left
		a.lineHeight = 0
	}
	if height > a.lineHeight {
		a.lineHeight = height
	}
	a.x += width + a.spacingH
	return true
}
