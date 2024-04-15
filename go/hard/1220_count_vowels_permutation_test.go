package hard

import (
	"testing"
)

func TestCountVowelPermutation(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{1, 5},
		{2, 10},
		{5, 68},
	}
	for _, c := range cases {
		got := CountVowelPermutation(c.n)
		if got != c.want {
			t.Errorf("CountVowelPermutation(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}
