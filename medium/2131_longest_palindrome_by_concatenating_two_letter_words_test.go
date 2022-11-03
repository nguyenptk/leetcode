package medium

import (
	"testing"
)

func TestLongestPalindromeConcatenating(t *testing.T) {
	cases := []struct {
		words []string
		want  int
	}{
		{[]string{"lc", "cl", "gg"}, 6},
		{[]string{"ab", "ty", "yt", "lc", "cl", "ab"}, 8},
		{[]string{"cc", "ll", "xx"}, 2},
	}
	for _, c := range cases {
		got := LongestPalindromeConcatenating(c.words)
		if got != c.want {
			t.Errorf("LongestPalindromeConcatenating(%s) == %d, want %d", c.words, got, c.want)
		}
	}
}
