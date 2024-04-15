package easy

import (
	"testing"
)

func TestArrayStringsAreEqual(t *testing.T) {
	cases := []struct {
		word1 []string
		word2 []string
		want  bool
	}{
		{[]string{"ab", "c"}, []string{"a", "bc"}, true},
		{[]string{"a", "cb"}, []string{"ab", "c"}, false},
		{[]string{"abc", "d", "defg"}, []string{"abcddefg"}, true},
	}
	for _, c := range cases {
		got := ArrayStringsAreEqual(c.word1, c.word2)
		if got != c.want {
			t.Errorf("ArrayStringsAreEqual(%s, %s) == %t, want %t", c.word1, c.word2, got, c.want)
		}
	}
}
