package svg

import (
	"math"
	"testing"
)

const epsilon = 1e-6

func TestSet(t *testing.T) {
	coef := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	matrix := newMatrix(coef[0], coef[1], coef[2], coef[3], coef[4], coef[5], coef[6], coef[7], coef[8])
	if matrix.n11 != coef[0] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[0])
	}
	if matrix.n12 != coef[1] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[1])
	}
	if matrix.n13 != coef[2] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[2])
	}
	if matrix.n21 != coef[3] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[3])
	}
	if matrix.n22 != coef[4] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[4])
	}
	if matrix.n23 != coef[5] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[5])
	}
	if matrix.n31 != coef[6] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[6])
	}
	if matrix.n32 != coef[7] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[7])
	}
	if matrix.n33 != coef[8] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[8])
	}
}

func TestClone(t *testing.T) {
	coef := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	matrix := newMatrix(coef[0], coef[1], coef[2], coef[3], coef[4], coef[5], coef[6], coef[7], coef[8]).clone()
	if matrix.n11 != coef[0] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[0])
	}
	if matrix.n12 != coef[1] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[1])
	}
	if matrix.n13 != coef[2] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[2])
	}
	if matrix.n21 != coef[3] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[3])
	}
	if matrix.n22 != coef[4] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[4])
	}
	if matrix.n23 != coef[5] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[5])
	}
	if matrix.n31 != coef[6] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[6])
	}
	if matrix.n32 != coef[7] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[7])
	}
	if matrix.n33 != coef[8] {
		t.Fatalf("invalid coef m11 got %g, exp %g", matrix.n11, coef[8])
	}
}

func TestMultiplyMatrix(t *testing.T) {
	tests := []struct {
		left   matrix
		right  matrix
		result matrix
	}{
		{
			left:   newMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9),
			right:  newMatrix(9, 2, 2, 8, 5, 0, 2, 7, 2),
			result: newMatrix(31, 33, 8, 88, 75, 20, 145, 117, 32),
		},
	}

	for _, tt := range tests {

		result := tt.left.multiplyMatrix(tt.right)

		if result.n11 != tt.result.n11 {
			t.Fatalf("invalid coef m11 got %g, exp %g", result.n11, tt.result.n11)
		}
		if result.n12 != tt.result.n12 {
			t.Fatalf("invalid coef m12 got %g, exp %g", result.n12, tt.result.n12)
		}
		if result.n13 != tt.result.n13 {
			t.Fatalf("invalid coef m13 got %g, exp %g", result.n13, tt.result.n13)
		}
		if result.n21 != tt.result.n21 {
			t.Fatalf("invalid coef m21 got %g, exp %g", result.n21, tt.result.n21)
		}
		if result.n22 != tt.result.n22 {
			t.Fatalf("invalid coef m22 got %g, exp %g", result.n22, tt.result.n22)
		}
		if result.n23 != tt.result.n23 {
			t.Fatalf("invalid coef m23 got %g, exp %g", result.n23, tt.result.n23)
		}
		if result.n31 != tt.result.n31 {
			t.Fatalf("invalid coef m31 got %g, exp %g", result.n31, tt.result.n31)
		}
		if result.n32 != tt.result.n32 {
			t.Fatalf("invalid coef m32 got %g, exp %g", result.n32, tt.result.n32)
		}
		if result.n33 != tt.result.n33 {
			t.Fatalf("invalid coef m33 got %g, exp %g", result.n33, tt.result.n33)
		}
	}
}

func TestMultiplyPoint(t *testing.T) {
	tests := []struct {
		mat    matrix
		points []float64
		result []float64
	}{
		{
			mat:    newMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9),
			points: []float64{9, 8, 2},
			result: []float64{31, 88, 145},
		},
	}

	for _, tt := range tests {

		x, y, z := tt.mat.multiplyPoint(tt.points[0], tt.points[1], tt.points[2])

		if x != tt.result[0] {
			t.Fatalf("invalid x got %g, exp %g", x, tt.result[0])
		}
		if y != tt.result[1] {
			t.Fatalf("invalid y got %g, exp %g", y, tt.result[1])
		}
		if z != tt.result[2] {
			t.Fatalf("invalid y got %g, exp %g", z, tt.result[2])
		}
	}
}

func TestScale(t *testing.T) {
	tests := []struct {
		mat    matrix
		scaleX float64
		scaleY float64
		result matrix
	}{
		{
			mat:    newMatrix(1, 0, 0, 0, 1, 0, 0, 0, 1),
			scaleX: 2,
			scaleY: 3,
			result: newMatrix(2, 0, 0, 0, 3, 0, 0, 0, 1),
		},
	}

	for _, tt := range tests {

		result := tt.mat.scale(tt.scaleX, tt.scaleY)

		if result.n11 != tt.result.n11 {
			t.Fatalf("invalid coef m11 got %g, exp %g", result.n11, tt.result.n11)
		}
		if result.n12 != tt.result.n12 {
			t.Fatalf("invalid coef m12 got %g, exp %g", result.n12, tt.result.n12)
		}
		if result.n13 != tt.result.n13 {
			t.Fatalf("invalid coef m13 got %g, exp %g", result.n13, tt.result.n13)
		}
		if result.n21 != tt.result.n21 {
			t.Fatalf("invalid coef m21 got %g, exp %g", result.n21, tt.result.n21)
		}
		if result.n22 != tt.result.n22 {
			t.Fatalf("invalid coef m22 got %g, exp %g", result.n22, tt.result.n22)
		}
		if result.n23 != tt.result.n23 {
			t.Fatalf("invalid coef m23 got %g, exp %g", result.n23, tt.result.n23)
		}
		if result.n31 != tt.result.n31 {
			t.Fatalf("invalid coef m31 got %g, exp %g", result.n31, tt.result.n31)
		}
		if result.n32 != tt.result.n32 {
			t.Fatalf("invalid coef m32 got %g, exp %g", result.n32, tt.result.n32)
		}
		if result.n33 != tt.result.n33 {
			t.Fatalf("invalid coef m33 got %g, exp %g", result.n33, tt.result.n33)
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		mat    matrix
		theta  float64
		result matrix
	}{
		{
			mat:    newMatrix(1, 0, 0, 0, 1, 0, 0, 0, 1),
			theta:  math.Pi,
			result: newMatrix(-1, 0, 0, 0, -1, 0, 0, 0, 1),
		},
		{
			mat:    newMatrix(1, 0, 0, 0, 1, 0, 0, 0, 1),
			theta:  math.Pi / 2,
			result: newMatrix(0, -1, 0, 1, 0, 0, 0, 0, 1),
		},
	}

	for _, tt := range tests {

		result := tt.mat.rotate(tt.theta)

		if math.Abs(result.n11-tt.result.n11) > epsilon {
			t.Fatalf("invalid coef m11 got %g, exp %g", result.n11, tt.result.n11)
		}
		if math.Abs(result.n12-tt.result.n12) > epsilon {
			t.Fatalf("invalid coef m12 got %g, exp %g", result.n12, tt.result.n12)
		}
		if math.Abs(result.n13-tt.result.n13) > epsilon {
			t.Fatalf("invalid coef m13 got %g, exp %g", result.n13, tt.result.n13)
		}
		if math.Abs(result.n21-tt.result.n21) > epsilon {
			t.Fatalf("invalid coef m21 got %g, exp %g", result.n21, tt.result.n21)
		}
		if math.Abs(result.n22-tt.result.n22) > epsilon {
			t.Fatalf("invalid coef m22 got %g, exp %g", result.n22, tt.result.n22)
		}
		if math.Abs(result.n23-tt.result.n23) > epsilon {
			t.Fatalf("invalid coef m23 got %g, exp %g", result.n23, tt.result.n23)
		}
		if math.Abs(result.n31-tt.result.n31) > epsilon {
			t.Fatalf("invalid coef m31 got %g, exp %g", result.n31, tt.result.n31)
		}
		if math.Abs(result.n32-tt.result.n32) > epsilon {
			t.Fatalf("invalid coef m32 got %g, exp %g", result.n32, tt.result.n32)
		}
		if math.Abs(result.n33-tt.result.n33) > epsilon {
			t.Fatalf("invalid coef m33 got %g, exp %g", result.n33, tt.result.n33)
		}
	}
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		mat    matrix
		tx     float64
		ty     float64
		result matrix
	}{
		{
			mat:    newMatrix(1, 0, 0, 0, 1, 0, 0, 0, 1),
			tx:     3,
			ty:     5,
			result: newMatrix(1, 0, 3, 0, 1, 5, 0, 0, 1),
		},
	}

	for _, tt := range tests {

		result := tt.mat.translate(tt.tx, tt.ty)

		if math.Abs(result.n11-tt.result.n11) > epsilon {
			t.Fatalf("invalid coef m11 got %g, exp %g", result.n11, tt.result.n11)
		}
		if math.Abs(result.n12-tt.result.n12) > epsilon {
			t.Fatalf("invalid coef m12 got %g, exp %g", result.n12, tt.result.n12)
		}
		if math.Abs(result.n13-tt.result.n13) > epsilon {
			t.Fatalf("invalid coef m13 got %g, exp %g", result.n13, tt.result.n13)
		}
		if math.Abs(result.n21-tt.result.n21) > epsilon {
			t.Fatalf("invalid coef m21 got %g, exp %g", result.n21, tt.result.n21)
		}
		if math.Abs(result.n22-tt.result.n22) > epsilon {
			t.Fatalf("invalid coef m22 got %g, exp %g", result.n22, tt.result.n22)
		}
		if math.Abs(result.n23-tt.result.n23) > epsilon {
			t.Fatalf("invalid coef m23 got %g, exp %g", result.n23, tt.result.n23)
		}
		if math.Abs(result.n31-tt.result.n31) > epsilon {
			t.Fatalf("invalid coef m31 got %g, exp %g", result.n31, tt.result.n31)
		}
		if math.Abs(result.n32-tt.result.n32) > epsilon {
			t.Fatalf("invalid coef m32 got %g, exp %g", result.n32, tt.result.n32)
		}
		if math.Abs(result.n33-tt.result.n33) > epsilon {
			t.Fatalf("invalid coef m33 got %g, exp %g", result.n33, tt.result.n33)
		}
	}
}
