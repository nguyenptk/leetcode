package medium

import "testing"

func TestMinOperations(t *testing.T) {
	cases := []struct {
		nums []int
		x    int
		want int
	}{
		{[]int{1, 1, 4, 2, 3}, 5, 2},
		{[]int{5, 6, 7, 8, 9}, 4, -1},
		{[]int{3, 2, 20, 1, 1, 3}, 10, 5},
	}
	for _, c := range cases {
		got := MinOperations(c.nums, c.x)
		if got != c.want {
			t.Errorf("MinOperations(%d) == %d, want %d", c.nums, got, c.want)
		}
	}
}
