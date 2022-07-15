package medium

import (
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	cases := []struct {
		matrix [][]int
		want   int
	}{
		{
			[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
			6,
		},
		{
			[][]int{{0, 0, 0, 0, 0, 0, 0, 0}},
			0,
		},
	}
	for _, c := range cases {
		got := MaxAreaOfIsland(c.matrix)
		if got != c.want {
			t.Errorf("MaxAreaOfIsland(%d) == %d, want %d", c.matrix, got, c.want)
		}
	}
}
