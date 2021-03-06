package svgnode

import (
	"github.com/canadadry/pml/pkg/matrix"
	"github.com/canadadry/pml/pkg/svg/svgdrawer"
	"github.com/canadadry/pml/pkg/svg/svgparser"
	"testing"
)

func TestSvgNodeDraw(t *testing.T) {
	tests := []struct {
		path      string
		transform matrix.Matrix
		expected  *svgdrawer.ForTesting
	}{
		{
			path:      "",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{}}
				return dcs
			}(),
		},
		{
			path:      "Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{"CloseAndDraw"}}
				return dcs
			}(),
		},
		{
			path:      "z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{"CloseAndDraw"}}
				return dcs
			}(),
		},
		{
			path:      "M1,2Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2M1,2Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"MoveTo x:1, y:2",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2m1,2Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"MoveTo x:12, y:14",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2m1,2Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
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
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:3, y:4",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2l2,2Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"LineTo x:13, y:14",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2L3,4Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"LineTo x:13, y:14",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2H3Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:3, y:2",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2H3Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"LineTo x:13, y:12",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2V3Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:1, y:3",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2V3Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"LineTo x:11, y:13",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2h3Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:4, y:2",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2h3Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"LineTo x:14, y:12",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2v3Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:1, y:5",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2v3Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"LineTo x:11, y:15",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2L3,4Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
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
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 7,8, anchor 1 3,4 anchor 2 5,6",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2C3,4 5,6 7,8 9,10 11,12 13,14Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 7,8, anchor 1 3,4 anchor 2 5,6",
					"BezierTo 13,14, anchor 1 9,10 anchor 2 11,12",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2C3,4 5,6 7,8Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"BezierTo 17,18, anchor 1 13,14 anchor 2 15,16",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2c3,4 5,6 7,8Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 8,10, anchor 1 4,6 anchor 2 6,8",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2c3,4 5,6 7,8 9,10 11,12 13,14Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 8,10, anchor 1 4,6 anchor 2 6,8",
					"BezierTo 21,24, anchor 1 17,20 anchor 2 19,22",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2c3,4 5,6 7,8Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"BezierTo 18,20, anchor 1 14,16 anchor 2 16,18",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2Q3,4 5,6Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 5,6, anchor 1 3,4 anchor 2 3,4",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2Q3,4 5,6 7,8 9,10Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 5,6, anchor 1 3,4 anchor 2 3,4",
					"BezierTo 9,10, anchor 1 7,8 anchor 2 7,8",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2Q3,4 5,6Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"BezierTo 15,16, anchor 1 13,14 anchor 2 13,14",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2q3,4 5,6Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 6,8, anchor 1 4,6 anchor 2 4,6",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2q3,4 5,6 7,8 9,10Z",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"BezierTo 6,8, anchor 1 4,6 anchor 2 4,6",
					"BezierTo 15,18, anchor 1 13,16 anchor 2 13,16",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2q3,4 5,6Z",
			transform: matrix.Identity().Translate(10, 10),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:11, y:12",
					"BezierTo 16,18, anchor 1 14,16 anchor 2 14,16",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2l2,2",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:3, y:4",
					"CloseAndDraw",
				}}
				return dcs
			}(),
		},
		{
			path:      "M1,2l2,2ZM3,4l2,2",
			transform: matrix.Identity(),
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
					"MoveTo x:1, y:2",
					"LineTo x:3, y:4",
					"CloseAndDraw",
					"MoveTo x:3, y:4",
					"LineTo x:5, y:6",
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

		result := &svgdrawer.ForTesting{Callstack: []string{}}

		err = node.Draw(result)

		if err != nil {
			t.Fatalf("[%s] failed : %v", tt.path, err)
		}
		if len(result.Callstack) != len(tt.expected.Callstack) {
			t.Errorf("expected stack %#v", tt.expected.Callstack)
			t.Errorf("  result stack %#v", result.Callstack)
			t.Fatalf("[%s] Callstack wrong size got %d exp %d", tt.path, len(result.Callstack), len(tt.expected.Callstack))
		}
		for j := range tt.expected.Callstack {
			if tt.expected.Callstack[j] != result.Callstack[j] {
				t.Fatalf("[%s] call %d got %s exp %s", tt.path, j, result.Callstack[j], tt.expected.Callstack[j])
			}
		}
	}
}
