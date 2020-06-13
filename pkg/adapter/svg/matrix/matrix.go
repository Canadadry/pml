package matrix

import (
	"fmt"
	"math"
)

type Matrix struct {
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

func New(
	n11 float64, n12 float64, n13 float64,
	n21 float64, n22 float64, n23 float64,
	n31 float64, n32 float64, n33 float64,
) Matrix {
	return Matrix{
		n11, n12, n13, n21, n22, n23, n31, n32, n33,
	}
}

func Identity() Matrix {
	return New(1, 0, 0, 0, 1, 0, 0, 0, 1)
}

func (m Matrix) Clone() Matrix {
	return New(m.n11, m.n12, m.n13, m.n21, m.n22, m.n23, m.n31, m.n32, m.n33)
}

func (a Matrix) MultiplyBy(b Matrix) Matrix {
	return New(
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

func (m Matrix) Project(x float64, y float64, z float64) (float64, float64, float64) {
	return m.n11*x + m.n12*y + m.n13*z, m.n21*x + m.n22*y + m.n23*z, m.n31*x + m.n32*y + m.n33*z
}

func (m Matrix) ProjectPoint(x float64, y float64) (float64, float64) {
	z := 1.0
	return m.n11*x + m.n12*y + m.n13*z, m.n21*x + m.n22*y + m.n23*z
}

func (m Matrix) Scale(sx float64, sy float64) Matrix {
	scaleMatrix := New(
		sx, 0, 0,
		0, sy, 0,
		0, 0, 1,
	)
	return m.MultiplyBy(scaleMatrix)
}

func (m Matrix) Rotate(theta float64) Matrix {
	c := math.Cos(theta)
	s := math.Sin(theta)

	rotationMatrix := New(
		c, -s, 0,
		s, c, 0,
		0, 0, 1,
	)
	return m.MultiplyBy(rotationMatrix)
}

func (m Matrix) Translate(tx float64, ty float64) Matrix {

	translationMatrix := New(
		1, 0, tx,
		0, 1, ty,
		0, 0, 1,
	)
	return m.MultiplyBy(translationMatrix)
}

func AreEquales(left Matrix, right Matrix) error {
	epsilon := 1e-6
	if math.Abs(left.n11-right.n11) > epsilon {
		return fmt.Errorf("invalid coef m11 got %g, exp %g", left.n11, right.n11)
	}
	if math.Abs(left.n12-right.n12) > epsilon {
		return fmt.Errorf("invalid coef m12 got %g, exp %g", left.n12, right.n12)
	}
	if math.Abs(left.n13-right.n13) > epsilon {
		return fmt.Errorf("invalid coef m13 got %g, exp %g", left.n13, right.n13)
	}
	if math.Abs(left.n21-right.n21) > epsilon {
		return fmt.Errorf("invalid coef m21 got %g, exp %g", left.n21, right.n21)
	}
	if math.Abs(left.n22-right.n22) > epsilon {
		return fmt.Errorf("invalid coef m22 got %g, exp %g", left.n22, right.n22)
	}
	if math.Abs(left.n23-right.n23) > epsilon {
		return fmt.Errorf("invalid coef m23 got %g, exp %g", left.n23, right.n23)
	}
	if math.Abs(left.n31-right.n31) > epsilon {
		return fmt.Errorf("invalid coef m31 got %g, exp %g", left.n31, right.n31)
	}
	if math.Abs(left.n32-right.n32) > epsilon {
		return fmt.Errorf("invalid coef m32 got %g, exp %g", left.n32, right.n32)
	}
	if math.Abs(left.n33-right.n33) > epsilon {
		return fmt.Errorf("invalid coef m33 got %g, exp %g", left.n33, right.n33)
	}
	return nil
}
