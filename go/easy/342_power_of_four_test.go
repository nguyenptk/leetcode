package easy

import (
	"testing"
)

func TestIsPowerOfFour(t *testing.T) {
	cases := []struct {
		n    int
		want bool
	}{
		{16, true},
		{5, false},
		{1, true},
	}
	for _, c := range cases {
		got := IsPowerOfFour(c.n)
		if got != c.want {
			t.Errorf("IsPowerOfFour(%d) == %t, want %t", c.n, got, c.want)
		}
	}
}
