package svgnode

import (
	"errors"
	"github.com/canadadry/pml/pkg/matrix"
	"github.com/canadadry/pml/pkg/svg/svgparser"
	"github.com/canadadry/pml/pkg/svg/svgpath"
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
		elem     *svgparser.Element
		commands []svgpath.Command
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
			commands: []svgpath.Command{
				{'M', []svgpath.Point{{0, 0}}},
				{'h', []svgpath.Point{{1, 0}}},
				{'v', []svgpath.Point{{1, 0}}},
				{'h', []svgpath.Point{{-1, 0}}},
				{'v', []svgpath.Point{{-1, 0}}},
				{'Z', []svgpath.Point{}},
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
			commands: []svgpath.Command{
				{'M', []svgpath.Point{{0, 0}}},
				{'h', []svgpath.Point{{25, 0}}},
				{'v', []svgpath.Point{{10, 0}}},
				{'h', []svgpath.Point{{-25, 0}}},
				{'v', []svgpath.Point{{0 - 10, 0}}},
				{'Z', []svgpath.Point{}},
			},
		},
	}

	for i, tt := range tests {
		node, err := svgRectangle(tt.elem, matrix.Identity())

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
