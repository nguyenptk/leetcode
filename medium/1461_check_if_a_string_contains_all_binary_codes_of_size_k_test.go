package medium

import (
	"testing"
)

func TestHasAllCodes(t *testing.T) {
	cases := []struct {
		s    string
		k    int
		want bool
	}{
		{"00110110", 2, true},
		{"0110", 1, true},
		{"0110", 2, false},
		{"00110", 2, true},
	}
	for _, c := range cases {
		got := HasAllCodes(c.s, c.k)
		if got != c.want {
			t.Errorf("HasAllCodes(%s, %d) == %t, want %t", c.s, c.k, got, c.want)
		}
	}
}
