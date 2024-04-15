package medium

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	cases := []struct {
		s    string
		p    string
		want []int
	}{
		{"abab", "ab", []int{0, 1, 2}},
		{"cbaebabacd", "abc", []int{0, 6}},
		{"ceffec", "efc", []int{0, 3}},
		{"hello", "olleh", []int{0}},
		{"bro", "bro", []int{0}},
	}
	for _, c := range cases {
		got := FindAnagrams(c.s, c.p)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FindAnagrams(%s, %s) == %d, want %d", c.s, c.p, got, c.want)
		}
	}
}
