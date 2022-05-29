package easy

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	cases := []struct {
		in   uint32
		want int
	}{
		{00000000000000000000000000001011, 3},
		{00000000000000000000000010000000, 1},
	}
	for _, c := range cases {
		got := HammingWeight(c.in)
		if got != c.want {
			t.Errorf("HammingWeight(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
