package medium

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	cases := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		// {[][]int{{0, 1}, {0, 0}}, 1},
	}
	for _, c := range cases {
		got := UniquePathsWithObstacles(c.in)
		if got != c.want {
			t.Errorf("UniquePathsWithObstacles(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
