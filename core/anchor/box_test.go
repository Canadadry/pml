package renderer

import (
	"testing"
)

func TestCutAlignment(t *testing.T) {
	tests := []struct {
		child  Box
		parent Box
		out    Box
	}{
		{
			parent: Box{0, 0, 100, 100, Free, Free},
			child:  Box{10, 10, 80, 80, Free, Free},
			out:    Box{10, 10, 80, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{15, 15, 80, 80, Free, Free},
			out:    Box{25, 25, 80, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{10, 10, 80, 80, Start, Free},
			out:    Box{10, 20, 80, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{10, 10, 80, 80, Center, Free},
			out:    Box{20, 20, 80, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{10, 10, 80, 80, End, Free},
			out:    Box{30, 20, 80, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{10, 10, 10, 80, Fill, Free},
			out:    Box{10, 20, 100, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{10, 10, 80, 80, Free, Start},
			out:    Box{20, 10, 80, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{10, 10, 80, 80, Free, Center},
			out:    Box{20, 20, 80, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{10, 10, 80, 80, Free, End},
			out:    Box{20, 30, 80, 80, Free, Free},
		},
		{
			parent: Box{10, 10, 100, 100, Free, Free},
			child:  Box{10, 10, 80, 20, Free, Fill},
			out:    Box{20, 10, 80, 100, Free, Free},
		},
	}

	for i, tt := range tests {
		result := tt.child.AlignIn(tt.parent)
		if result != tt.out {
			t.Fatalf("[%d] failed got %#v exp %#v", i, result, tt.out)
		}
	}
}
