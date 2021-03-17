package renderer

import (
	"testing"
)

func TestCutAlignment(t *testing.T) {
	tests := []struct {
		In    renderBox
		Frame Frame
		Exp   renderBox
	}{
		{
			In:    renderBox{0, 0, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 80, Relative, Relative, 1.0},
			Exp:   renderBox{10, 10, 80, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{15, 15, 80, 80, Relative, Relative, 1.0},
			Exp:   renderBox{25, 25, 80, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 80, Left, Relative, 1.0},
			Exp:   renderBox{10, 20, 80, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 80, Center, Relative, 1.0},
			Exp:   renderBox{20, 20, 80, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 80, Right, Relative, 1.0},
			Exp:   renderBox{30, 20, 80, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{10, 10, 10, 80, Fill, Relative, 1.0},
			Exp:   renderBox{10, 20, 100, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 80, Relative, Top, 1.0},
			Exp:   renderBox{20, 10, 80, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 80, Relative, Center, 1.0},
			Exp:   renderBox{20, 20, 80, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 80, Relative, Bottom, 1.0},
			Exp:   renderBox{20, 30, 80, 80, 1.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 20, Relative, Fill, 1.0},
			Exp:   renderBox{20, 10, 80, 100, 1.0},
		},
	}

	for i, tt := range tests {
		result := tt.In.Cut(tt.Frame)
		if result != tt.Exp {
			t.Fatalf("[%d] failed got %#v exp %#v", i, result, tt.Exp)
		}
	}
}

func TestCutScaling(t *testing.T) {
	tests := []struct {
		In    renderBox
		Frame Frame
		Exp   renderBox
	}{
		{
			In:    renderBox{0, 0, 100, 100, 1.0},
			Frame: Frame{10, 10, 80, 80, Relative, Relative, 2.0},
			Exp:   renderBox{10, 10, 160, 160, 2.0},
		},
		{
			In:    renderBox{0, 0, 100, 100, 2.0},
			Frame: Frame{10, 10, 80, 80, Relative, Relative, 1.0},
			Exp:   renderBox{20, 20, 160, 160, 2.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 2.0},
			Frame: Frame{10, 10, 80, 80, Relative, Relative, 1.0},
			Exp:   renderBox{30, 30, 160, 160, 2.0},
		},
		{
			In:    renderBox{10, 10, 100, 100, 2.0},
			Frame: Frame{10, 10, 80, 80, Relative, Relative, 2.0},
			Exp:   renderBox{30, 30, 320, 320, 4.0},
		},
	}

	for i, tt := range tests {
		result := tt.In.Cut(tt.Frame)
		if result != tt.Exp {
			t.Fatalf("[%d] failed got %#v exp %#v", i, result, tt.Exp)
		}
	}
}
