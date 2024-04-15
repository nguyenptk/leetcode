package easy

import (
	"testing"
)

func TestWordPattern(t *testing.T) {
	cases := []struct {
		pattern string
		s       string
		want    bool
	}{
		{"abba", "dog cat cat dog", true},
	}
	for _, c := range cases {
		got := WordPattern(c.pattern, c.s)
		if got != c.want {
			t.Errorf("WordPattern(%s, %s) == %t, want %t", c.pattern, c.s, got, c.want)
		}
	}
}
