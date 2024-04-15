package easy

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for _, c := range cases {
		got := ContainsNearbyDuplicate(c.nums, c.k)
		if got != c.want {
			t.Errorf("ContainsNearbyDuplicate(%d, %d) == %t, want %t", c.nums, c.k, got, c.want)
		}
	}
}
