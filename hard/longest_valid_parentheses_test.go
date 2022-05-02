package hard

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"()(())", 6},
		{"()(()))))", 6},
		{"((()()))", 8},
	}
	for _, c := range cases {
		got := LongestValidParentheses(c.in)
		if got != c.want {
			t.Errorf("IsValid(%q) == %d, want %d", c.in, got, c.want)
		}
	}

}
