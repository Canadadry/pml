package svgpath

type Point struct {
	X float64
	Y float64
}

type Command struct {
	Kind   byte
	Points []Point
}

func (c Command) ToString() string {
	return "test" //fmt.Sprintf("%s, 1: %g,%g 2: %g,%g 3: %g,%g", string(c.kind), c.x1, c.y1, c.x2, c.y2, c.x3, c.y3)
}
