package matrix

import (
	"math"
	"testing"
)

func TestSet(t *testing.T) {
	coef := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	matrix := New(coef[0], coef[1], coef[2], coef[3], coef[4], coef[5], coef[6], coef[7], coef[8])
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

	matrix := New(coef[0], coef[1], coef[2], coef[3], coef[4], coef[5], coef[6], coef[7], coef[8]).Clone()
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

func TestMultiplyBy(t *testing.T) {
	tests := []struct {
		left   Matrix
		right  Matrix
		result Matrix
	}{
		{
			left:   New(1, 2, 3, 4, 5, 6, 7, 8, 9),
			right:  New(9, 2, 2, 8, 5, 0, 2, 7, 2),
			result: New(31, 33, 8, 88, 75, 20, 145, 117, 32),
		},
	}

	for _, tt := range tests {

		result := tt.left.MultiplyBy(tt.right)

		err := AreEquales(result, tt.result)
		if err != nil {
			t.Fatalf("%v", err)
		}

	}
}

func TestProject(t *testing.T) {
	tests := []struct {
		mat    Matrix
		points []float64
		result []float64
	}{
		{
			mat:    New(1, 2, 3, 4, 5, 6, 7, 8, 9),
			points: []float64{9, 8, 2},
			result: []float64{31, 88, 145},
		},
	}

	for _, tt := range tests {

		x, y, z := tt.mat.Project(tt.points[0], tt.points[1], tt.points[2])

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

func TestProjectPoint(t *testing.T) {
	tests := []struct {
		mat    Matrix
		points []float64
		result []float64
	}{
		{
			mat:    New(1, 2, 3, 4, 5, 6, 0, 0, 1),
			points: []float64{9, 8},
			result: []float64{28, 82},
		},
	}

	for _, tt := range tests {

		x, y := tt.mat.ProjectPoint(tt.points[0], tt.points[1])

		if x != tt.result[0] {
			t.Fatalf("invalid x got %g, exp %g", x, tt.result[0])
		}
		if y != tt.result[1] {
			t.Fatalf("invalid y got %g, exp %g", y, tt.result[1])
		}
	}
}

func TestScale(t *testing.T) {
	tests := []struct {
		mat    Matrix
		scaleX float64
		scaleY float64
		result Matrix
	}{
		{
			mat:    New(1, 0, 0, 0, 1, 0, 0, 0, 1),
			scaleX: 2,
			scaleY: 3,
			result: New(2, 0, 0, 0, 3, 0, 0, 0, 1),
		},
	}

	for _, tt := range tests {

		result := tt.mat.Scale(tt.scaleX, tt.scaleY)

		err := AreEquales(result, tt.result)
		if err != nil {
			t.Fatalf("%v", err)
		}

	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		mat    Matrix
		theta  float64
		result Matrix
	}{
		{
			mat:    New(1, 0, 0, 0, 1, 0, 0, 0, 1),
			theta:  math.Pi,
			result: New(-1, 0, 0, 0, -1, 0, 0, 0, 1),
		},
		{
			mat:    New(1, 0, 0, 0, 1, 0, 0, 0, 1),
			theta:  math.Pi / 2,
			result: New(0, -1, 0, 1, 0, 0, 0, 0, 1),
		},
	}

	for _, tt := range tests {

		result := tt.mat.Rotate(tt.theta)

		err := AreEquales(result, tt.result)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		mat    Matrix
		tx     float64
		ty     float64
		result Matrix
	}{
		{
			mat:    New(1, 0, 0, 0, 1, 0, 0, 0, 1),
			tx:     3,
			ty:     5,
			result: New(1, 0, 3, 0, 1, 5, 0, 0, 1),
		},
	}

	for _, tt := range tests {

		result := tt.mat.Translate(tt.tx, tt.ty)

		err := AreEquales(result, tt.result)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}
}

func TestAreEquales(t *testing.T) {
	tests := []struct {
		left        Matrix
		right       Matrix
		shouldMatch bool
	}{
		{
			left:        New(1, 0, 0, 0, 1, 0, 0, 0, 1),
			right:       New(1, 0, 3, 0, 1, 5, 0, 0, 1),
			shouldMatch: false,
		},
		{
			left:        New(1, 0, 3, 0, 1, 5, 0, 0, 1),
			right:       New(1, 0, 3, 0, 1, 5, 0, 0, 1),
			shouldMatch: true,
		},
		{
			left:        Identity(),
			right:       Identity().Scale(2, 3),
			shouldMatch: false,
		},
	}

	for i, tt := range tests {

		err := AreEquales(tt.left, tt.right)

		if tt.shouldMatch {
			if err != nil {
				t.Fatalf("[%d] matrix should be equlas but :%v", i, err)
			}
		}
		if !tt.shouldMatch {
			if err == nil {
				t.Fatalf("[%d] matrix should not be equlas but \"AreEquales\" say they are", i)
			}
		}

	}
}
