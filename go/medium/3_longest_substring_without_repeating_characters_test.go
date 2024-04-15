package medium

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, c := range cases {
		got := LengthOfLongestSubstring(c.in)
		if got != c.want {
			t.Errorf("LengthOfLongestSubstring(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}
