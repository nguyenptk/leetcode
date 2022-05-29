package medium

import "testing"

func TestCountSubstrings(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"abc", 3},
		{"aaa", 6},
	}
	for _, c := range cases {
		got := CountSubstrings(c.in)
		if got != c.want {
			t.Errorf("CountSubstrings(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}
