package easy

import "testing"

func TestRemovePalindromeSub(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"ababa", 1},
		{"abb", 2},
		{"baabb", 2},
	}
	for _, c := range cases {
		got := RemovePalindromeSub(c.in)
		if got != c.want {
			t.Errorf("RemovePalindromeSub(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
