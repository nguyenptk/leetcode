package easy

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}
	for _, c := range cases {
		got := IsPalindrome(c.in)
		if got != c.want {
			t.Errorf("IsPalindrome(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}
