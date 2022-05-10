package medium

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	cases := []struct {
		k    int
		n    int
		want [][]int
	}{
		{3, 7, [][]int{{1, 2, 4}}},
		{3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}
	for _, c := range cases {
		got := CombinationSum3(c.k, c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CombinationSum3(%d, %d) == %d, want %d", c.k, c.n, got, c.want)
		}
	}
}
