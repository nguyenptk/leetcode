package medium

import (
	"testing"
)

func TestMaximumScore(t *testing.T) {
	cases := []struct {
		nums        []int
		multipliers []int
		want        int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, 14},
		{[]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}, 102},
	}
	for _, c := range cases {
		got := MaximumScore(c.nums, c.multipliers)
		if got != c.want {
			t.Errorf("MaximumScore(%d, %d) == %d, want %d", c.nums, c.multipliers, got, c.want)
		}
	}
}
