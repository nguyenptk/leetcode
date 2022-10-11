package medium

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
	}
	for _, c := range cases {
		got := IncreasingTriplet(c.nums)
		if got != c.want {
			t.Errorf("IncreasingTriplet(%d) == %t, want %t", c.nums, got, c.want)
		}
	}
}
