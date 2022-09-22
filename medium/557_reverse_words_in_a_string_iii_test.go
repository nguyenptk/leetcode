package medium

import "testing"

func TestReverseWords(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
	}
	for _, c := range cases {
		got := ReverseWords(c.s)
		if got != c.want {
			t.Errorf("ReverseWords(%s) == %s, want %s", c.s, got, c.want)
		}
	}
}
