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
			path:      "",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{}}
				return dcs
			}(),
		},
		{
			path:      "Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{"CloseAndDraw"}}
				return dcs
			}(),
		},
		{
			path:      "M1,2Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2M1,2Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"MoveTo x:1, y:2",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2m1,2Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"MoveTo x:2, y:4",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2l2,2Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:3, y:4",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2L3,4Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:3, y:4",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2C3,4 5,6 7,8Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 7,8, anchor 1 3,4 anchor 2 5,6",
					"CloseAndDraw",
				}}
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
