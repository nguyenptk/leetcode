package easy

import "testing"

func TestMakeGood(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"leEeetcode", "leetcode"},
		{"abBAcC", ""},
		{"s", "s"},
	}
	for _, c := range cases {
		got := MakeGood(c.s)
		if got != c.want {
			t.Errorf("MakeGood(%s) == %s, want %s", c.s, got, c.want)
		}
	}
}
