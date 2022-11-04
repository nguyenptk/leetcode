package easy

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
	}
	for _, c := range cases {
		got := ReverseVowels(c.s)
		if got != c.want {
			t.Errorf("ReverseVowels(%s) == %s, want %s", c.s, got, c.want)
		}
	}
}
