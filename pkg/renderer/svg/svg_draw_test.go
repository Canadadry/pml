package svg

import (
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"testing"
)

func TestSvgNodeDraw(t *testing.T) {
	tests := []struct {
		path      string
		transform matrix.Matrix
		expected  *drawCallStack
	}{
		{
			path:      "Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{"CloseAndDraw"}}
				return dcs
			}(),
		},
	}

	for i, tt := range tests {
		pathElement := &svgparser.Element{
			Name: "Path",
			Attributes: map[string]string{
				"d": tt.path,
			},
		}

		node, err := svgPath(pathElement, tt.transform)
		if err != nil {
			t.Fatalf("[%d] failed : %v", i, err)
		}

		result := &drawCallStack{callstack: []string{}}

		err = node.draw(result)

		if err != nil {
			t.Fatalf("[%d] failed : %v", i, err)
		}
		if len(result.callstack) != len(tt.expected.callstack) {
			t.Errorf("expected stack %#v", tt.expected.callstack)
			t.Errorf("  result stack %#v", result.callstack)
			t.Fatalf("[%d] callstack wrong size got %d exp %d", i, len(result.callstack), len(tt.expected.callstack))
		}
		for j := range tt.expected.callstack {
			if tt.expected.callstack[j] != result.callstack[j] {
				t.Fatalf("[%d] call %d got %s exp %s", i, j, result.callstack[j], tt.expected.callstack[j])
			}
		}
	}
}
