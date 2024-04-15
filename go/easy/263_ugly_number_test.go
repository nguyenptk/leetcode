package easy

import (
	"testing"
)

func TestIsUgly(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{6, true},
		{1, true},
		{14, false},
	}
	for _, c := range cases {
		got := IsUgly(c.in)
		if got != c.want {
			t.Errorf("IsUgly(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}
