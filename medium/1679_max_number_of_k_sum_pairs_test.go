package medium

import "testing"

func TestMaxOperations(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		{[]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2, 2},
	}
	for _, c := range cases {
		got := MaxOperations(c.in, c.target)
		if got != c.want {
			t.Errorf("MaxOperations(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
