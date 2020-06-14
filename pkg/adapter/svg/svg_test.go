package svg

import (
	"fmt"
	"github.com/canadadry/pml/pkg/abstract/abstractsvg"
	"strings"
	"testing"
)

type drawCallStack struct {
	callstack []string
}

func (dcs *drawCallStack) MoveTo(x float64, y float64) {
	dcs.callstack = append(dcs.callstack,
		fmt.Sprintf("MoveTo x:%g, y:%g", x, y),
	)
}

func (dcs *drawCallStack) LineTo(x float64, y float64) {
	dcs.callstack = append(dcs.callstack,
		fmt.Sprintf("LineTo x:%g, y:%g", x, y),
	)
}

func (dcs *drawCallStack) BezierTo(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) {
	dcs.callstack = append(dcs.callstack,
		fmt.Sprintf("BezierTo %g,%g, anchor 1 %g,%g anchor 2 %g,%g", x3, y3, x1, y1, x2, y2),
	)
}

func (dcs *drawCallStack) CloseAndDraw(s abstractsvg.Style) {
	dcs.callstack = append(dcs.callstack,
		fmt.Sprintf("CloseAndDraw"),
	)
}

func TestEndToEnd(t *testing.T) {
	tests := []struct {
		svg      string
		x        float64
		y        float64
		w        float64
		h        float64
		expected *drawCallStack
	}{
		{
			svg: `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg width="100%" height="100%" viewBox="0 0 256 256" version="1.1" xmlns="http://www.w3.org/2000/svg">
</svg>`,
			x: 0, y: 0, w: 100, h: 100,
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{}}
				return dcs
			}(),
		},
		{
			svg: `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg width="100%" height="100%" viewBox="0 0 210 297"  style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:1.41421;">
    <rect x="100" y="50" width="32" height="32" style="fill:rgb(51,153,102);fill-rule:nonzero;"/>
</svg>`,
			x: 0, y: 0, w: 210, h: 297,
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:100, y:50",
					"LineTo x:132, y:50",
					"LineTo x:132, y:82",
					"LineTo x:100, y:82",
					"LineTo x:100, y:50",
					"CloseAndDraw",
				}}

				return dcs
			}(),
		},
		{
			svg: `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg width="100%" height="100%" viewBox="0 0 210 297"  style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:1.41421;">
    <rect x="100" y="50" width="32" height="32" style="fill:rgb(51,153,102);fill-rule:nonzero;"/>
</svg>`,
			x: 0, y: 0, w: 210, h: 0,
			expected: func() *drawCallStack {
				dcs := &drawCallStack{callstack: []string{
					"MoveTo x:100, y:50",
					"LineTo x:132, y:50",
					"LineTo x:132, y:82",
					"LineTo x:100, y:82",
					"LineTo x:100, y:50",
					"CloseAndDraw",
				}}

				return dcs
			}(),
		},
	}

	for i, tt := range tests {
		result := &drawCallStack{callstack: []string{}}
		err := New().Draw(result, strings.NewReader(tt.svg), tt.x, tt.y, tt.w, tt.h)
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
