package easy

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{2, 2},
		{3, 3},
	}
	for _, c := range cases {
		got := ClimbStairs(c.in)
		if got != c.want {
			t.Errorf("ClimbStairs(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
