package svgpath

import (
	"testing"
)

func TestToString(t *testing.T) {

	tests := []struct {
		command  []Command
		expected []string
	}{
		{
			[]Command{
				{'m', []Point{{234.43804, 111.69821}}},
				{'c', []Point{{50.21866, 26.50627}, {126.75595, -3.87395}, {151.46369, -35.941621}}},
			},
			[]string{
				"m 0 : (234.43804,111.69821)",
				"c 0 : (50.21866,26.50627), 1 : (126.75595,-3.87395), 2 : (151.46369,-35.941621)",
			},
		},
	}

	for i, tt := range tests {
		for j := range tt.command {
			if tt.command[j].ToString() != tt.expected[j] {
				t.Fatalf("[%d] command %d got %s exp %s", i, j, tt.command[j].ToString(), tt.expected[j])
			}
		}
	}
}
