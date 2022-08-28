package medium

import (
	"reflect"
	"testing"
)

func TestDiagonalSort(t *testing.T) {
	cases := []struct {
		mat  [][]int
		want [][]int
	}{
		{
			[][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}},
			[][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}}},
		{
			[][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}},
			[][]int{{5, 17, 4, 1, 52, 7}, {11, 11, 25, 45, 8, 69}, {14, 23, 25, 44, 58, 15}, {22, 27, 31, 36, 50, 66}, {84, 28, 75, 33, 55, 68}},
		},
	}
	for _, c := range cases {
		got := DiagonalSort(c.mat)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("DiagonalSort(%d) == %d, want %d", c.mat, got, c.want)
		}
	}
}
