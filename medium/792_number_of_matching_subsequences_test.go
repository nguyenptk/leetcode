package medium

import "testing"

func TestNumMatchingSubseq(t *testing.T) {
	cases := []struct {
		s     string
		words []string
		want  int
	}{
		{"abcde", []string{"a", "bb", "acd", "ace"}, 3},
	}
	for _, c := range cases {
		got := NumMatchingSubseq(c.s, c.words)
		if got != c.want {
			t.Errorf("NumMatchingSubseq(%s) == %d, want %d", c.words, got, c.want)
		}
	}
}
