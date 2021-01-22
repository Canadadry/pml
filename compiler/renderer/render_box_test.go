package renderer

import (
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		In    renderBox
		Frame Frame
		Exp   renderBox
	}{
		{
			In:    renderBox{0, 0, 100, 100},
			Frame: Frame{10, 10, 80, 80, Free, Free},
			Exp:   renderBox{10, 10, 80, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{15, 15, 80, 80, Free, Free},
			Exp:   renderBox{25, 25, 80, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{10, 10, 80, 80, Left, Free},
			Exp:   renderBox{10, 20, 80, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{10, 10, 80, 80, Center, Free},
			Exp:   renderBox{20, 20, 80, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{10, 10, 80, 80, Right, Free},
			Exp:   renderBox{30, 20, 80, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{10, 10, 10, 80, Fill, Free},
			Exp:   renderBox{10, 20, 100, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{10, 10, 80, 80, Free, Top},
			Exp:   renderBox{20, 10, 80, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{10, 10, 80, 80, Free, Middle},
			Exp:   renderBox{20, 20, 80, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{10, 10, 80, 80, Free, Bottom},
			Exp:   renderBox{20, 30, 80, 80},
		},
		{
			In:    renderBox{10, 10, 100, 100},
			Frame: Frame{10, 10, 80, 20, Free, Fill},
			Exp:   renderBox{20, 10, 80, 100},
		},
	}

	for i, tt := range tests {
		result := tt.In.Cut(tt.Frame)
		if result != tt.Exp {
			t.Fatalf("[%d] failed got %#v exp %#v", i, result, tt.Exp)
		}
	}
}
