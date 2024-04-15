package medium

import (
	"reflect"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	cases := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
	}
	for _, c := range cases {
		got := IsInterleave(c.s1, c.s2, c.s3)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("IsInterleave(%s, %s, %s) == %t, want %t", c.s1, c.s2, c.s3, got, c.want)
		}
	}
}
