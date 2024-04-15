package medium

import (
	"testing"
)

func TestFindPaths(t *testing.T) {
	cases := []struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
		want        int
	}{
		{
			2,
			2,
			2,
			0,
			0,
			6,
		},
		{
			1,
			3,
			3,
			0,
			1,
			12,
		},
	}
	for _, c := range cases {
		got := FindPaths(c.m, c.n, c.maxMove, c.startRow, c.startColumn)

		if got != c.want {
			t.Errorf("FindPaths(%d, %d, %d, %d, %d) == %d, want %d", c.m, c.n, c.maxMove, c.startRow, c.startColumn, got, c.want)
		}
	}
}
