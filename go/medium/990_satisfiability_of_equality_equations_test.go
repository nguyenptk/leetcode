package medium

import "testing"

func TestEquationsPossible(t *testing.T) {
	cases := []struct {
		words []string
		want  bool
	}{
		{[]string{"a==b", "b!=a"}, false},
		{[]string{"b==a", "a==b"}, true},
	}
	for _, c := range cases {
		got := EquationsPossible(c.words)
		if got != c.want {
			t.Errorf("EquationsPossible(%s) == %t, want %t", c.words, got, c.want)
		}
	}
}
