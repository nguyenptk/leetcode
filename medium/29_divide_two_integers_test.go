package medium

import (
	"testing"
)

func TestDivide(t *testing.T) {
	cases := []struct {
		dividend int
		divisor  int
		want     int
	}{
		{10, 3, 3},
		{7, -3, -2},
	}
	for _, c := range cases {
		got := Divide(c.dividend, c.divisor)
		if got != c.want {
			t.Errorf("Divide(%d, %d) == %d, want %d", c.dividend, c.divisor, got, c.want)
		}
	}
}
