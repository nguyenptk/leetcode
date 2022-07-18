package hard

import (
	"testing"
)

func TestNumSubmatrixSumTarget(t *testing.T) {
	cases := []struct {
		matrix [][]int
		target int
		want   int
	}{
		{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0, 4},
		{[][]int{{1, -1}, {-1, 1}}, 0, 5},
		{[][]int{{904}}, 0, 0},
	}
	for _, c := range cases {
		got := NumSubmatrixSumTarget(c.matrix, c.target)
		if got != c.want {
			t.Errorf("NumSubmatrixSumTarget(%d, %d) == %d, want %d", c.matrix, c.target, got, c.want)
		}
	}
}
