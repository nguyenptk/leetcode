package hard

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	cases := []struct {
		words  []string
		prefix string
		suffix string
		want   int
	}{
		// Test case 1
		{
			[]string{"apple"},
			"a",
			"e",
			0,
		},
		// Test case 2
		{
			[]string{"apple"},
			"a",
			"l",
			-1,
		},
	}
	for _, c := range cases {
		myWordFilter := Constructor(c.words)
		// call f function
		got := myWordFilter.F(c.prefix, c.suffix)
		if got != c.want {
			t.Errorf("Constructor(words=%s, prefix=%s, suffix=%s) == %d, want %d", c.words, c.prefix, c.suffix, got, c.want)
		}
	}
}
