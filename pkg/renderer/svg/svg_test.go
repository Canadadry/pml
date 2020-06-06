package svg

import (
	"fmt"
	"strings"
	"testing"
)

type drawCallStack struct {
	callstack []string
}

func (dcs *drawCallStack) MoveTo(x float64, y float64) {
	dcs.callstack = append(dcs.callstack,
		fmt.Sprintf("MoveTo x:%g, y: %g", x, y),
	)
}

func (dcs *drawCallStack) LineTo(x float64, y float64) {
	dcs.callstack = append(dcs.callstack,
		fmt.Sprintf("LineTo x:%g, y: %g", x, y),
	)
}

func (dcs *drawCallStack) ClosePath() {
	dcs.callstack = append(dcs.callstack,
		fmt.Sprintf("ClosePath"),
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
	}

	for i, tt := range tests {
		result := &drawCallStack{callstack: []string{}}
		err := Draw(result, strings.NewReader(tt.svg), tt.x, tt.y, tt.w, tt.h)
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
