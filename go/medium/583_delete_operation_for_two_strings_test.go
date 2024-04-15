package medium

import "testing"

func TestMinDistance(t *testing.T) {
	cases := []struct {
		word1 string
		word2 string
		want  int
	}{
		{"sea", "eat", 2},
		{"leetcode", "etco", 4},
		{"a", "a", 0},
	}
	for _, c := range cases {
		got := MinDistance(c.word1, c.word2)
		if got != c.want {
			t.Errorf("MinDistance(%s, %s) == %d, want %d", c.word1, c.word2, got, c.want)
		}
	}
}
