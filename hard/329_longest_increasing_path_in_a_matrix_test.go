package hard

import (
	"reflect"
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {
	cases := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4},
	}
	for _, c := range cases {
		got := LongestIncreasingPath(c.matrix)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("LongestIncreasingPath(%d) == %d, want %d", c.matrix, got, c.want)
		}
	}
}
