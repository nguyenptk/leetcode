package medium

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}
	for _, c := range cases {
		got := LengthOfLIS(c.nums)
		if got != c.want {
			t.Errorf("LengthOfLIS(%d) == %d, want %d", c.nums, got, c.want)
		}
	}
}
