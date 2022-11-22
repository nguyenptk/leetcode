package hard

import "testing"

func TestNumSquares(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{12, 3},
		{13, 2},
	}
	for _, c := range cases {
		got := NumSquares(c.n)
		if got != c.want {
			t.Errorf("NumSquares(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}
