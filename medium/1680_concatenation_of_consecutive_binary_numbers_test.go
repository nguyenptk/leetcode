package medium

import "testing"

func TestConcatenatedBinary(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{1, 1},
		{3, 27},
		{12, 505379714},
	}
	for _, c := range cases {
		got := ConcatenatedBinary(c.in)
		if got != c.want {
			t.Errorf("ConcatenatedBinary(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
