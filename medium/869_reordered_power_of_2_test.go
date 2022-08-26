package medium

import "testing"

func TestReorderedPowerOf2(t *testing.T) {
	cases := []struct {
		n    int
		want bool
	}{
		{1, true},
		{10, false},
	}
	for _, c := range cases {
		got := ReorderedPowerOf2(c.n)
		if got != c.want {
			t.Errorf("ReorderedPowerOf2(%d) == %t, want %t", c.n, got, c.want)
		}
	}
}
