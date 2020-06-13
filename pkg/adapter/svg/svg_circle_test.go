package svg

import (
	"errors"
	"github.com/canadadry/pml/pkg/renderer/svg/matrix"
	"github.com/canadadry/pml/pkg/renderer/svg/svgparser"
	"github.com/canadadry/pml/pkg/renderer/svg/svgpath"
	"testing"
)

func TestSvgCircleErrors(t *testing.T) {
	tests := []struct {
		elem *svgparser.Element
		err  error
	}{
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{},
			},
			err: svgparser.ErrMissingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"cx": "0.0",
				},
			},
			err: svgparser.ErrMissingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"cx": "a",
				},
			},
			err: svgparser.ErrParsingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"cx": "0.0",
					"cy": "0.0",
				},
			},
			err: svgparser.ErrMissingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"cx": "0.0",
					"cy": "a",
				},
			},
			err: svgparser.ErrParsingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"cx": "0.0",
					"cy": "0.0",
					"r":  "1.0",
				},
			},
			err: nil,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"cx": "0.0",
					"cy": "0.0",
					"r":  "a",
				},
			},
			err: svgparser.ErrParsingAttr,
		},
	}

	for i, tt := range tests {
		_, err := svgCircle(tt.elem, matrix.Identity())

		if !errors.Is(err, tt.err) {
			t.Fatalf("[%d] wrong err returned got %v, exp %v", i, err, tt.err)
		}
	}
}

func TestSvgCircleCommands(t *testing.T) {
	tests := []struct {
		elem     *svgparser.Element
		commands []svgpath.Command
	}{
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"cx": "5.0",
					"cy": "5.0",
					"r":  "5.0",
				},
			},
			// circle with bezier curve param : 2.761423749153967
			//     r (5) - bezier curve param : 2.238576250846033
			commands: []svgpath.Command{
				{'M', []svgpath.Point{{5, 0}}},
				{'C', []svgpath.Point{{7.761423749153967, 0}, {10, 2.238576250846033}, {10, 5}}},
				{'C', []svgpath.Point{{10, 7.761423749153967}, {7.761423749153967, 10}, {5, 10}}},
				{'C', []svgpath.Point{{2.238576250846033, 10}, {0, 7.761423749153967}, {0, 5}}},
				{'C', []svgpath.Point{{0, 2.238576250846033}, {2.238576250846033, 0}, {5, 0}}},
				{'Z', []svgpath.Point{}},
			},
		},
	}

	for i, tt := range tests {
		node, err := svgCircle(tt.elem, matrix.Identity())

		if err != nil {
			t.Fatalf("[%d] failed : %v", i, err)
		}
		if len(node.commands) != len(tt.commands) {
			t.Errorf("expected commands %#v", tt.commands)
			t.Errorf("  result commands %#v", node.commands)
			t.Fatalf("[%d] commands wrong size got %d exp %d", i, len(node.commands), len(tt.commands))
		}
		for j := range tt.commands {
			if tt.commands[j].ToString() != node.commands[j].ToString() {
				t.Fatalf("[%d] command %d got %s exp %s", i, j, node.commands[j].ToString(), tt.commands[j].ToString())
			}
		}
	}
}
