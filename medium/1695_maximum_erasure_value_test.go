package medium

import "testing"

func TestMaximumUniqueSubarray(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{4, 2, 4, 5, 6}, 17},
		{[]int{5, 2, 1, 2, 5, 2, 1, 2, 5}, 8},
	}
	for _, c := range cases {
		got := MaximumUniqueSubarray(c.in)
		if got != c.want {
			t.Errorf("MaximumUniqueSubarray(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
