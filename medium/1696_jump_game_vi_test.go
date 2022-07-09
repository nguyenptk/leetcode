package medium

import "testing"

func TestMaxResult(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, -1, -2, 4, -7, 3}, 2, 7},
		{[]int{10, -5, -2, 4, 0, 3}, 3, 17},
		{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2, 0},
	}
	for _, c := range cases {
		got := MaxResult(c.nums, c.k)
		if got != c.want {
			t.Errorf("MaxResult(%d, %d) == %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
