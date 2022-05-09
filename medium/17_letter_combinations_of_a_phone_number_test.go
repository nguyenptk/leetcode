package medium

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, c := range cases {
		got := LetterCombinations(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("LetterCombinations(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}
