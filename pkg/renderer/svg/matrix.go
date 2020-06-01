package svg

import (
	"math"
)

type matrix struct {
	n11 float64
	n12 float64
	n13 float64
	n21 float64
	n22 float64
	n23 float64
	n31 float64
	n32 float64
	n33 float64
}

func newMatrix(
	n11 float64, n12 float64, n13 float64,
	n21 float64, n22 float64, n23 float64,
	n31 float64, n32 float64, n33 float64,
) matrix {
	return matrix{
		n11, n12, n13, n21, n22, n23, n31, n32, n33,
	}
}

func (m matrix) clone() matrix {
	return newMatrix(m.n11, m.n12, m.n13, m.n21, m.n22, m.n23, m.n31, m.n32, m.n33)
}

func (a matrix) multiplyMatrix(b matrix) matrix {
	return newMatrix(
		a.n11*b.n11+a.n12*b.n21+a.n13*b.n31,
		a.n11*b.n12+a.n12*b.n22+a.n13*b.n32,
		a.n11*b.n13+a.n12*b.n23+a.n13*b.n33,
		a.n21*b.n11+a.n22*b.n21+a.n23*b.n31,
		a.n21*b.n12+a.n22*b.n22+a.n23*b.n32,
		a.n21*b.n13+a.n22*b.n23+a.n23*b.n33,
		a.n31*b.n11+a.n32*b.n21+a.n33*b.n31,
		a.n31*b.n12+a.n32*b.n22+a.n33*b.n32,
		a.n31*b.n13+a.n32*b.n23+a.n33*b.n33,
	)
}

func (m matrix) multiplyPoint(x float64, y float64, z float64) (float64, float64, float64) {
	return m.n11*x + m.n12*y + m.n13*z, m.n21*x + m.n22*y + m.n23*z, m.n31*x + m.n32*y + m.n33*z
}

func (m matrix) scale(sx float64, sy float64) matrix {
	scaleMatrix := newMatrix(
		sx, 0, 0,
		0, sy, 0,
		0, 0, 1,
	)
	return m.multiplyMatrix(scaleMatrix)
}

func (m matrix) rotate(theta float64) matrix {
	c := math.Cos(theta)
	s := math.Sin(theta)

	rotationMatrix := newMatrix(
		c, -s, 0,
		s, c, 0,
		0, 0, 1,
	)
	return m.multiplyMatrix(rotationMatrix)
}

func (m matrix) translate(tx float64, ty float64) matrix {

	translationMatrix := newMatrix(
		1, 0, tx,
		0, 1, ty,
		0, 0, 1,
	)
	return m.multiplyMatrix(translationMatrix)
}
