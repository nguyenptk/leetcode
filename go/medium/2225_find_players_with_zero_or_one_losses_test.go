package medium

import (
	"reflect"
	"testing"
)

func TestFindWinners(t *testing.T) {
	cases := []struct {
		matches [][]int
		want    [][]int
	}{
		{[][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}, [][]int{{1, 2, 10}, {4, 5, 7, 8}}},
		{[][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}, [][]int{{1, 2, 5, 6}, nil}},
	}
	for _, c := range cases {
		got := FindWinners(c.matches)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FindWinners(%d) == %d, want %d", c.matches, got, c.want)
		}
	}
}
