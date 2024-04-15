package medium

import (
	"testing"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	cases := []struct {
		k    [][]int
		want int
	}{
		{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, -1},
	}
	for _, c := range cases {
		got := ShortestPathBinaryMatrix(c.k)
		if got != c.want {
			t.Errorf("ShortestPathBinaryMatrix(%d) == %d, want %d", c.k, got, c.want)
		}
	}
}
