package easy

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	cases := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for _, c := range cases {
		got := CanConstruct(c.ransomNote, c.magazine)
		if got != c.want {
			t.Errorf("CanConstruct(%s, %s) == %t, want %t", c.ransomNote, c.magazine, got, c.want)
		}
	}
}
