package medium

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, c := range cases {
		got := FindKthLargest(c.nums, c.k)
		if got != c.want {
			t.Errorf("FindKthLargest(%d, %d) == %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
