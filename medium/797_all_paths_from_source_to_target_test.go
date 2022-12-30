package medium

import (
	"reflect"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	cases := []struct {
		graph [][]int
		want  [][]int
	}{
		{
			[][]int{{1, 2}, {3}, {3}, {}},
			[][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			[][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
	}
	for _, c := range cases {
		got := AllPathsSourceTarget(c.graph)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("AllPathsSourceTarget(%d) == %d, want %d", c.graph, got, c.want)
		}
	}
}
