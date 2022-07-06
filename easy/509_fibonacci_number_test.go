package easy

import (
	"testing"
)

func TestFib(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for _, c := range cases {
		got := Fib(c.in)
		if got != c.want {
			t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
