package medium

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		// {"babad", "bab"},
		// {"cbbd", "bb"},
		{"forgeeksskeegfor", "geeksskeeg"},
	}
	for _, c := range cases {
		got := LongestPalindrome(c.in)
		if got != c.want {
			t.Errorf("LongestPalindrome(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}
