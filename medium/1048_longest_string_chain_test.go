package medium

import "testing"

func TestLongestStrChain(t *testing.T) {
	cases := []struct {
		words []string
		want  int
	}{
		{[]string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
		{[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
		{[]string{"abcd", "dbqca"}, 1},
	}
	for _, c := range cases {
		got := LongestStrChain(c.words)
		if got != c.want {
			t.Errorf("LongestStrChain(%s) == %d, want %d", c.words, got, c.want)
		}
	}
}
