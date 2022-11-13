package medium

import "testing"

func TestReverseWords(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
	}
	for _, c := range cases {
		got := ReverseWords(c.s)
		if got != c.want {
			t.Errorf("ReverseWords(%s) == %s, want %s", c.s, got, c.want)
		}
	}
}
