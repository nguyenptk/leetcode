package medium

import "testing"

func TestNumRollsToTarget(t *testing.T) {
	cases := []struct {
		n      int
		k      int
		target int
		want   int
	}{
		{1, 6, 3, 1},
		{2, 6, 7, 6},
		{30, 30, 500, 222616187},
	}
	for _, c := range cases {
		got := NumRollsToTarget(c.n, c.k, c.target)
		if got != c.want {
			t.Errorf("NumRollsToTarget(%d, %d, %d) == %d, want %d", c.n, c.k, c.target, got, c.want)
		}
	}
}
