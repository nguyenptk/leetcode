package medium

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
