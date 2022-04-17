package medium

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	}
	for _, c := range cases {
		got := MaxArea(c.in)
		if got != c.want {
			t.Errorf("MaxArea(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
