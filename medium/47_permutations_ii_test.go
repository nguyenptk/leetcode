package medium

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	cases := []struct {
		k    []int
		want [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}
	for _, c := range cases {
		got := PermuteUnique(c.k)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("PermuteUnique(%d) == %d, want %d", c.k, got, c.want)
		}
	}
}
