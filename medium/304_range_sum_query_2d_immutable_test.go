package medium

import (
	"reflect"
	"testing"
)

func TestConstructor2D(t *testing.T) {
	cases := []struct {
		matrix [][]int
		sum    []int
		want   int
	}{
		{
			[][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}},
			[]int{2, 1, 4, 3},
			8,
		},
		{
			[][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}},
			[]int{1, 1, 2, 2},
			11,
		},
		{
			[][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}},
			[]int{1, 2, 2, 4},
			12,
		},
	}
	for _, c := range cases {
		obj := Constructor2D(c.matrix)
		got := obj.SumRegion(c.sum[0], c.sum[1], c.sum[2], c.sum[3])

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Constructor2D(%d) == %d, want %d", c.matrix, got, c.want)
		}
	}
}
