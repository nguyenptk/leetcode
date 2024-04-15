package medium

import "testing"

func TestWordBreak(t *testing.T) {
	cases := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, c := range cases {
		got := WordBreak(c.s, c.wordDict)
		if got != c.want {
			t.Errorf("WordBreak(%s, %s) == %t, want %t", c.s, c.wordDict, got, c.want)
		}
	}
}
