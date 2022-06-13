package medium

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	cases := []struct {
		triangle [][]int
		want     int
	}{
		{
			[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			11,
		},
		{
			[][]int{{-10}},
			-10,
		},
		{
			[][]int{{-1}, {2, 3}, {1, -1, -3}},
			-1,
		},
	}
	for _, c := range cases {
		got := MinimumTotal(c.triangle)

		if got != c.want {
			t.Errorf("MinimumTotal(%d) == %d, want %d", c.triangle, got, c.want)
		}
	}
}
