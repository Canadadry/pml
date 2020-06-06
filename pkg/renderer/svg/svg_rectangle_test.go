package svg

import (
	"errors"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"testing"
)

func TestSvgRectangleErrors(t *testing.T) {
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
					"x": "0.0",
				},
			},
			err: svgparser.ErrMissingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x": "a",
				},
			},
			err: svgparser.ErrParsingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x": "0.0",
					"y": "0.0",
				},
			},
			err: svgparser.ErrMissingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x": "0.0",
					"y": "a",
				},
			},
			err: svgparser.ErrParsingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x":     "0.0",
					"y":     "0.0",
					"width": "1.0",
				},
			},
			err: svgparser.ErrMissingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x":     "0.0",
					"y":     "0.0",
					"width": "a",
				},
			},
			err: svgparser.ErrParsingAttr,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x":      "0.0",
					"y":      "0.0",
					"width":  "1.0",
					"height": "1.0",
				},
			},
			err: nil,
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x":      "0.0",
					"y":      "0.0",
					"width":  "1.0",
					"height": "a",
				},
			},
			err: svgparser.ErrParsingAttr,
		},
	}

	for i, tt := range tests {
		_, err := svgRectangle(tt.elem, matrix.Identity())

		if !errors.Is(err, tt.err) {
			t.Fatalf("[%d] wrong err returned got %v, exp %v", i, err, tt.err)
		}
	}
}

func TestSvgRectangleCommands(t *testing.T) {
	tests := []struct {
		elem      *svgparser.Element
		transform matrix.Matrix
		commands  []command
	}{
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x":      "0.0",
					"y":      "0.0",
					"width":  "1.0",
					"height": "1.0",
				},
			},
			transform: matrix.Identity(),
			commands: []command{
				{'M', 0, 0, 0, 0, 0, 0},
				{'L', 1, 0, 0, 0, 0, 0},
				{'L', 1, 1, 0, 0, 0, 0},
				{'L', 0, 1, 0, 0, 0, 0},
				{'L', 0, 0, 0, 0, 0, 0},
				{'Z', 0, 0, 0, 0, 0, 0},
			},
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x":      "0.0",
					"y":      "0.0",
					"width":  "25.0",
					"height": "10.0",
				},
			},
			transform: matrix.Identity(),
			commands: []command{
				{'M', 0, 0, 0, 0, 0, 0},
				{'L', 25, 0, 0, 0, 0, 0},
				{'L', 25, 10, 0, 0, 0, 0},
				{'L', 0, 10, 0, 0, 0, 0},
				{'L', 0, 0, 0, 0, 0, 0},
				{'Z', 0, 0, 0, 0, 0, 0},
			},
		},
		{
			elem: &svgparser.Element{
				Attributes: map[string]string{
					"x":      "0.0",
					"y":      "0.0",
					"width":  "25.0",
					"height": "10.0",
				},
			},
			transform: matrix.Identity().Translate(5, 12),
			commands: []command{
				{'M', 5, 12, 0, 0, 0, 0},
				{'L', 30, 12, 0, 0, 0, 0},
				{'L', 30, 22, 0, 0, 0, 0},
				{'L', 5, 22, 0, 0, 0, 0},
				{'L', 5, 12, 0, 0, 0, 0},
				{'Z', 0, 0, 0, 0, 0, 0},
			},
		},
	}

	for i, tt := range tests {
		node, err := svgRectangle(tt.elem, tt.transform)

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
