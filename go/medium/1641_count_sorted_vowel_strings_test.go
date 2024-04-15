package medium

import "testing"

func TestCountVowelStrings(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{1, 5},
		{2, 15},
		{33, 66045},
	}
	for _, c := range cases {
		got := CountVowelStrings(c.in)
		if got != c.want {
			t.Errorf("CountVowelStrings(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
