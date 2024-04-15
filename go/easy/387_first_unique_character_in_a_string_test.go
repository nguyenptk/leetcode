package easy

import (
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}
	for _, c := range cases {
		got := FirstUniqChar(c.in)
		if got != c.want {
			t.Errorf("FirstUniqChar(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}
