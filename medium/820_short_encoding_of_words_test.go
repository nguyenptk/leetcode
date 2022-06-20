package medium

import "testing"

func TestMinimumLengthEncoding(t *testing.T) {
	cases := []struct {
		words []string
		want  int
	}{
		{[]string{"time", "me", "bell"}, 10},
		{[]string{"t"}, 2},
	}
	for _, c := range cases {
		got := MinimumLengthEncoding(c.words)
		if got != c.want {
			t.Errorf("MinimumLengthEncoding(%s) == %d, want %d", c.words, got, c.want)
		}
	}
}
