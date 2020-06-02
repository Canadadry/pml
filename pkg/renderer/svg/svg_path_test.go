package svg

import (
	"math"
	"testing"
)

func TestPathParsing(t *testing.T) {
	epsilon := 1e-6

	tests := []struct {
		path     string
		tranform matrix
		expected []command
	}{
		{
			path:     "M148,148C148,88.448 196.449,40 256,40C315.551,40 364,88.448 364,148L10,12.33 21,23Z",
			tranform: identityMatrix(),
			expected: []command{
				{'M', 148, 148, 0, 0, 0, 0},
				{'C', 148, 88.448, 196.449, 40, 256, 40},
				{'C', 315.551, 40, 364, 88.448, 364, 148},
				{'L', 10, 12.33, 21, 23, 0, 0},
				{'Z', 0, 0, 0, 0, 0, 0},
			},
		},
		{
			path:     "M123,456Z",
			tranform: identityMatrix().scale(2, 3),
			expected: []command{
				{'M', 246, 1368, 0, 0, 0, 0},
				{'Z', 0, 0, 0, 0, 0, 0},
			},
		},
	}

	for i, tt := range tests {
		pathElement := &Element{
			Name: "Path",
			Attributes: map[string]string{
				"d": tt.path,
			},
		}

		result, err := path(pathElement, tt.tranform)

		if err != nil {
			t.Fatalf("[%d] should not failed but got %v", i, err)
		}

		if len(result.commands) != len(tt.expected) {
			t.Fatalf("[%d] result path has not the right number of commands exp=%d got=%d", i, len(tt.expected), len(result.commands))
		}

		for j, gotCmd := range result.commands {
			expCmd := tt.expected[j]
			if expCmd.kind != gotCmd.kind {
				t.Fatalf("[%d] cmd %d not the right kind exp=%s got=%s", i, j, string(expCmd.kind), string(gotCmd.kind))
			}
			if math.Abs(expCmd.x1-gotCmd.x1) > epsilon {
				t.Fatalf("[%d] cmd %d not the right x1 exp=%g got=%g", i, j, expCmd.x1, gotCmd.x1)
			}
			if math.Abs(expCmd.y1-gotCmd.y1) > epsilon {
				t.Fatalf("[%d] cmd %d not the right y1 exp=%g got=%g", i, j, expCmd.y1, gotCmd.y1)
			}
			if math.Abs(expCmd.x2-gotCmd.x2) > epsilon {
				t.Fatalf("[%d] cmd %d not the right x2 exp=%g got=%g", i, j, expCmd.x2, gotCmd.x2)
			}
			if math.Abs(expCmd.y2-gotCmd.y2) > epsilon {
				t.Fatalf("[%d] cmd %d not the right y2 exp=%g got=%g", i, j, expCmd.y2, gotCmd.y2)
			}
			if math.Abs(expCmd.x3-gotCmd.x3) > epsilon {
				t.Fatalf("[%d] cmd %d not the right x3 exp=%g got=%g", i, j, expCmd.x3, gotCmd.x3)
			}
			if math.Abs(expCmd.y3-gotCmd.y3) > epsilon {
				t.Fatalf("[%d] cmd %d not the right y3 exp=%g got=%g", i, j, expCmd.y3, gotCmd.y3)
			}
		}
	}

}
