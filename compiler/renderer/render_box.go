package renderer

type renderBox struct {
	x float64
	y float64
	w float64
	h float64
}

func (rb renderBox) Cut(f Frame) renderBox {
	out := renderBox{
		x: rb.x + f.x,
		y: rb.y + f.y,
		w: f.width,
		h: f.height,
	}
	switch f.xAlign {
	case Left:
		out.x = rb.x
	case Center:
		out.x = rb.x + (rb.w-f.width)/2
	case Right:
		out.x = rb.x + rb.w - f.width
	case Fill:
		out.x = rb.x
		out.w = rb.w
	}
	switch f.yAlign {
	case Top:
		out.y = rb.y
	case Middle:
		out.y = rb.y + (rb.h-f.height)/2
	case Bottom:
		out.y = rb.y + rb.h - f.height
	case Fill:
		out.y = rb.y
		out.h = rb.h
	}
	return out
}
