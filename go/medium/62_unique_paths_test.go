package medium

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	cases := []struct {
		m    int
		n    int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}
	for _, c := range cases {
		got := UniquePaths(c.m, c.n)
		if got != c.want {
			t.Errorf("UniquePaths(%d, %d) == %d, want %d", c.m, c.n, got, c.want)
		}
	}
}
