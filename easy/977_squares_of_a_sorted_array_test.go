package easy

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, c := range cases {
		got := SortedSquares(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SortedSquares(%d) == %d, want %d", c.in, got, c.want)
		}
	}

}
