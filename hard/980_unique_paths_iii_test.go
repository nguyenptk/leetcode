package hard

import (
	"testing"
)

func TestUniquePathsIII(t *testing.T) {
	cases := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}, 2},
		{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}, 4},
		{[][]int{{0, 1}, {2, 0}}, 0},
	}
	for _, c := range cases {
		got := UniquePathsIII(c.matrix)
		if got != c.want {
			t.Errorf("UniquePathsIII(%d) == %d, want %d", c.matrix, got, c.want)
		}
	}
}
