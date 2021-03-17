package renderer

type renderBox struct {
	x float64
	y float64
	w float64
	h float64
	s float64
}

func (rb renderBox) Cut(f Frame) renderBox {

	f.width *= f.scale * rb.s
	f.height *= f.scale * rb.s

	out := renderBox{
		x: rb.x + f.x*rb.s,
		y: rb.y + f.y*rb.s,
		w: f.width,
		h: f.height,
		s: f.scale * rb.s,
	}
	switch f.xAnchor {
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
	switch f.yAnchor {
	case Top:
		out.y = rb.y
	case Center:
		out.y = rb.y + (rb.h-f.height)/2
	case Bottom:
		out.y = rb.y + rb.h - f.height
	case Fill:
		out.y = rb.y
		out.h = rb.h
	}
	return out
}
