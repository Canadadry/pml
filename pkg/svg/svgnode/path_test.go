package svgnode

import (
	"github.com/canadadry/pml/pkg/matrix"
	"github.com/canadadry/pml/pkg/svg/svgparser"
	"github.com/canadadry/pml/pkg/svg/svgpath"
	"testing"
)

func TestPathParsing(t *testing.T) {

	tests := []struct {
		path     string
		expected []svgpath.Command
	}{
		{
			path: "M148,148C148,88.448 196.449,40 256,40C315.551,40 364,88.448 364,148L10,12.33Z",
			expected: []svgpath.Command{
				{'M', []svgpath.Point{{148, 148}}},
				{'C', []svgpath.Point{{148, 88.448}, {196.449, 40}, {256, 40}}},
				{'C', []svgpath.Point{{315.551, 40}, {364, 88.448}, {364, 148}}},
				{'L', []svgpath.Point{{10, 12.33}}},
				{'Z', []svgpath.Point{}},
			},
		},
	}

	for i, tt := range tests {
		pathElement := &svgparser.Element{
			Name: "Path",
			Attributes: map[string]string{
				"d": tt.path,
			},
		}

		result, err := svgPath(pathElement, matrix.Identity().Scale(1, 2))

		if err != nil {
			t.Fatalf("[%s] should not failed but got %v", tt.path, err)
		}

		if len(result.commands) != len(tt.expected) {
			t.Fatalf("[%d] result path has not the right number of commands exp=%d got=%d", i, len(tt.expected), len(result.commands))
		}

		for j := range tt.expected {
			if tt.expected[j].ToString() != result.commands[j].ToString() {
				t.Fatalf("[%d] command %d got %s exp %s", i, j, result.commands[j].ToString(), tt.expected[j].ToString())
			}
		}
	}
}
