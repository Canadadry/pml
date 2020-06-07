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
			path:      "z",
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
			path:      "M1,2H3Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:3, y:2",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2V3Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:1, y:3",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2h3Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:4, y:2",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2v3Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:1, y:5",
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
		{
			path:      "M1,2c3,4 5,6 7,8Z",
			transform: matrix.Identity(),
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 8,10, anchor 1 4,6 anchor 2 6,8",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
	}

	for _, tt := range tests {
		pathElement := &svgparser.Element{
			Name: "Path",
			Attributes: map[string]string{
				"d": tt.path,
			},
		}

		node, err := svgPath(pathElement, tt.transform)
		if err != nil {
			t.Fatalf("[%s] failed : %v", tt.path, err)
		}

		result := &drawCallStack{callstack: []string{}}

		err = node.draw(result)

		if err != nil {
			t.Fatalf("[%s] failed : %v", tt.path, err)
		}
		if len(result.callstack) != len(tt.expected.callstack) {
			t.Errorf("expected stack %#v", tt.expected.callstack)
			t.Errorf("  result stack %#v", result.callstack)
			t.Fatalf("[%s] callstack wrong size got %d exp %d", tt.path, len(result.callstack), len(tt.expected.callstack))
		}
		for j := range tt.expected.callstack {
			if tt.expected.callstack[j] != result.callstack[j] {
				t.Fatalf("[%s] call %d got %s exp %s", tt.path, j, result.callstack[j], tt.expected.callstack[j])
			}
		}
	}
}
