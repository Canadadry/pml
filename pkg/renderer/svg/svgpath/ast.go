package svgpath

import (
	"fmt"
	"strings"
)

type Point struct {
	X float64
	Y float64
}

type Command struct {
	Kind   byte
	Points []Point
}

func (c Command) ToString() string {

	points := []string{}

	for i, p := range c.Points {
		points = append(points, fmt.Sprintf("%d : (%g,%g)", i, p.X, p.Y))
	}

	return string(c.Kind) + " " + strings.Join(points, ", ")
}
