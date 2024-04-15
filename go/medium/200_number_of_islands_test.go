package medium

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	cases := []struct {
		grid [][]byte
		want int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
	}
	for _, c := range cases {
		got := NumIslands(c.grid)
		if got != c.want {
			t.Errorf("NumIslands(%s) == %d, want %d", c.grid, got, c.want)
		}
	}
}
