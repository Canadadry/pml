package renderer

type Anchor string

const (
	Start  Anchor = "start"
	End           = "end"
	Center        = "center"
	Fill          = "fill"
	Free          = "free"
)

type Box struct {
	X       float64
	Y       float64
	W       float64
	H       float64
	XAnchor Anchor
	YAnchor Anchor
}

func (child Box) AlignIn(parent Box) Box {
	x, w := child.XAnchor.align(child.X, child.W, parent.X, parent.W)
	y, h := child.YAnchor.align(child.Y, child.H, parent.Y, parent.H)
	return Box{X: x, Y: y, W: w, H: h, XAnchor: Free, YAnchor: Free}
}

func (a Anchor) align(startChild, sizeChild, startParent, sizeParent float64) (float64, float64) {
	outStart := startChild + startParent
	outSize := sizeChild
	switch a {
	case Start:
		outStart = startParent
	case Center:
		outStart = startParent + (sizeParent-sizeChild)/2
	case End:
		outStart = startParent + (sizeParent - sizeChild)
	case Fill:
		outStart = startParent
		outSize = sizeParent
	}
	return outStart, outSize
}
