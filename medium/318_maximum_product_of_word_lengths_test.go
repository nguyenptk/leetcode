package medium

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	cases := []struct {
		words []string
		want  int
	}{
		{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}, 16},
		{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
		{[]string{"a", "aa", "aaa", "aaaa"}, 0},
	}
	for _, c := range cases {
		got := MaxProduct(c.words)
		if got != c.want {
			t.Errorf("MaxProduct(%s) == %d, want %d", c.words, got, c.want)
		}
	}
}
