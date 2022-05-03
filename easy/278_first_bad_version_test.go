package easy

import (
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		// {2, 1},
		{5, 4},
	}
	for _, c := range cases {
		got := FirstBadVersion(c.in)
		if got != c.want {
			t.Errorf("FirstBadVersion(%q) == %d, want %d", c.in, got, c.want)
		}
	}

}
