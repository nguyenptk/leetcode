package medium

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{4, "1211"},
	}
	for _, c := range cases {
		got := CountAndSay(c.n)
		if got != c.want {
			t.Errorf("CountAndSay(%d) == %s, want %s", c.n, got, c.want)
		}
	}
}
