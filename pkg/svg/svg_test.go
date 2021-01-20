package svg

import (
	"github.com/canadadry/pml/pkg/svg/svgdrawer"
	"strings"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	tests := []struct {
		svg      string
		x        float64
		y        float64
		w        float64
		h        float64
		expected *svgdrawer.ForTesting
	}{
		{
			svg: `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg width="100%" height="100%" viewBox="0 0 256 256" version="1.1" xmlns="http://www.w3.org/2000/svg">
</svg>`,
			x: 0, y: 0, w: 100, h: 100,
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{}}
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
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
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
			expected: func() *svgdrawer.ForTesting {
				dcs := &svgdrawer.ForTesting{Callstack: []string{
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
		result := &svgdrawer.ForTesting{Callstack: []string{}}
		err := Draw(result, strings.NewReader(tt.svg), tt.x, tt.y, tt.w, tt.h)
		if err != nil {
			t.Fatalf("[%d] failed : %v", i, err)
		}
		if len(result.Callstack) != len(tt.expected.Callstack) {
			t.Errorf("expected stack %#v", tt.expected.Callstack)
			t.Errorf("  result stack %#v", result.Callstack)
			t.Fatalf("[%d] Callstack wrong size got %d exp %d", i, len(result.Callstack), len(tt.expected.Callstack))
		}
		for j := range tt.expected.Callstack {
			if tt.expected.Callstack[j] != result.Callstack[j] {
				t.Fatalf("[%d] call %d got %s exp %s", i, j, result.Callstack[j], tt.expected.Callstack[j])
			}
		}
	}
}
