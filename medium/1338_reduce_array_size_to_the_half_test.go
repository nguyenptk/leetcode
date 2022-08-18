package medium

import "testing"

func TestMinSetSize(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
		{[]int{7, 7, 7, 7, 7}, 1},
	}
	for _, c := range cases {
		got := MinSetSize(c.in)
		if got != c.want {
			t.Errorf("MinSetSize(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
