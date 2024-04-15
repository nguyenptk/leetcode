package medium

import (
	"reflect"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	cases := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			5,
			true,
		},
		{
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			20,
			false,
		},
		{
			[][]int{{-5}},
			-5,
			true,
		},
	}
	for _, c := range cases {
		got := SearchMatrixII(c.matrix, c.target)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SearchMatrix(%d, %d) == %t, want %t", c.matrix, c.target, got, c.want)
		}
	}
}
