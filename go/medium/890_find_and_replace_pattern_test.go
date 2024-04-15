package medium

import (
	"reflect"
	"testing"
)

func TestFindAndReplacePattern(t *testing.T) {
	cases := []struct {
		words   []string
		pattern string
		want    []string
	}{
		{[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb", []string{"mee", "aqq"}},
		{[]string{"a", "b", "c"}, "a", []string{"a", "b", "c"}},
	}
	for _, c := range cases {
		got := FindAndReplacePattern(c.words, c.pattern)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FindAndReplacePattern(%s, %s) == %s, want %s", c.words, c.pattern, got, c.want)
		}
	}
}
