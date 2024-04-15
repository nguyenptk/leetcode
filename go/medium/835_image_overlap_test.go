package medium

import (
	"testing"
)

func TestLargestOverlap(t *testing.T) {
	cases := []struct {
		img1 [][]int
		img2 [][]int
		want int
	}{
		// test case 1
		{
			[][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			[][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
			3,
		},
		// test case 2
		{
			[][]int{{1}},
			[][]int{{1}},
			1,
		},
		// test case 3
		{
			[][]int{{0}},
			[][]int{{0}},
			0,
		},
	}
	for _, c := range cases {
		got := LargestOverlap(c.img1, c.img2)
		if got != c.want {
			t.Errorf("LargestOverlap(%d, %d) == %d, want %d", c.img1, c.img2, got, c.want)
		}
	}
}
