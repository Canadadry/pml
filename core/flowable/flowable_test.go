package flowable

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type Size struct{ w, h float64 }
	tests := map[string]struct {
		A                  Area
		Objs               []Size
		expectedX          []float64
		expectedY          []float64
		expectedLH         []float64
		shouldAllObjectFit bool
	}{
		"one object": {
			A:                  New(0, 0, 100, 100, 0, 0),
			Objs:               []Size{{20, 20}},
			expectedX:          []float64{20},
			expectedY:          []float64{0},
			expectedLH:         []float64{20},
			shouldAllObjectFit: true,
		},
		"one line of object": {
			A:                  New(0, 0, 100, 100, 0, 0),
			Objs:               []Size{{50, 20}, {50, 20}, {50, 20}},
			expectedX:          []float64{50, 100, 50},
			expectedY:          []float64{0, 0, 20},
			expectedLH:         []float64{20, 20, 20},
			shouldAllObjectFit: true,
		},
		"one line of different height object": {
			A:                  New(0, 0, 100, 100, 0, 0),
			Objs:               []Size{{50, 40}, {50, 20}, {50, 20}},
			expectedX:          []float64{50, 100, 50},
			expectedY:          []float64{0, 0, 40},
			expectedLH:         []float64{40, 40, 20},
			shouldAllObjectFit: true,
		},
		"classic overflow": {
			A:                  New(0, 0, 100, 100, 0, 0),
			Objs:               []Size{{45, 45}, {45, 45}, {45, 45}, {45, 45}, {45, 45}},
			expectedX:          []float64{45, 90, 45, 90, 90},
			expectedY:          []float64{0, 0, 45, 45, 45},
			expectedLH:         []float64{45, 45, 45, 45, 45},
			shouldAllObjectFit: false,
		},
		"classic overflow with margin": {
			A:                  New(0, 0, 100, 100, 6, 2),
			Objs:               []Size{{45, 45}, {45, 45}, {45, 45}, {45, 45}, {45, 45}},
			expectedX:          []float64{51, 102, 51, 102, 102},
			expectedY:          []float64{0, 0, 47, 47, 47},
			expectedLH:         []float64{45, 45, 45, 45, 45},
			shouldAllObjectFit: false,
		},
	}

	for title, tt := range tests {
		a := tt.A
		allObjectFit := true
		for i, o := range tt.Objs {
			ok := a.Add(o.w, o.h)
			allObjectFit = allObjectFit && ok
			if a.GetX() != tt.expectedX[i] {
				t.Fatalf("%s : after Add obj %d for GetX got %v wanted %v", title, i, a.GetX(), tt.expectedX[i])
			}
			if a.GetY() != tt.expectedY[i] {
				t.Fatalf("%s : after Add obj %d for GetY got %v wanted %v", title, i, a.GetY(), tt.expectedY[i])
			}
			if a.GetLineHeight() != tt.expectedLH[i] {
				t.Fatalf("%s : after Add obj %d for GetLineHeight got %v wanted %v", title, i, a.GetLineHeight(), tt.expectedLH[i])
			}
		}
		if allObjectFit != tt.shouldAllObjectFit {
			t.Fatalf("%s : for AllObjectFit got %v wanted %v", title, allObjectFit, tt.shouldAllObjectFit)
		}
	}
}
