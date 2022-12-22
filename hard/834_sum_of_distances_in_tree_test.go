package hard

import (
	"reflect"
	"testing"
)

func TestSumOfDistancesInTree(t *testing.T) {
	cases := []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, []int{8, 12, 6, 10, 10, 10}},
		{1, [][]int{}, []int{0}},
		{2, [][]int{{1, 0}}, []int{1, 1}},
	}
	for _, c := range cases {
		got := SumOfDistancesInTree(c.n, c.edges)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SumOfDistancesInTree(%d, %d) == %d, want %d", c.n, c.edges, got, c.want)
		}
	}
}
