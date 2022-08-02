package medium

import "testing"

func TestKthSmallest(t *testing.T) {
	cases := []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
		{[][]int{{-5}}, 1, -5},
		{[][]int{{1, 2}, {1, 3}}, 2, 1},
	}
	for _, c := range cases {
		got := KthSmallest(c.matrix, c.k)
		if got != c.want {
			t.Errorf("KthSmallest(%d, %d) == %d, want %d", c.matrix, c.k, got, c.want)
		}
	}
}
