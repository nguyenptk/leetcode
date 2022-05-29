package easy

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{-2, -5, 6, -2, -3, 1, 5, -6}, 7},
	}
	for _, c := range cases {
		got := MaxSubArray(c.in)
		if got != c.want {
			t.Errorf("MaxSubArray(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
