package easy

import (
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	cases := []struct {
		n    int
		want bool
	}{
		{27, true},
		{0, false},
		{9, true},
	}
	for _, c := range cases {
		got := IsPowerOfThree(c.n)
		if got != c.want {
			t.Errorf("IsPowerOfThree(%d) == %t, want %t", c.n, got, c.want)
		}
	}
}
