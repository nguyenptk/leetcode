package easy

import (
	"testing"
)

func TestIsAlienSorted(t *testing.T) {
	cases := []struct {
		words []string
		order string
		want  bool
	}{
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
	}
	for _, c := range cases {
		got := IsAlienSorted(c.words, c.order)
		if got != c.want {
			t.Errorf("IsAlienSorted(%s, %s) == %t, want %t", c.words, c.order, got, c.want)
		}
	}
}
