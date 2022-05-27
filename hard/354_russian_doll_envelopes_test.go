package hard

import (
	"reflect"
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	cases := []struct {
		n    [][]int
		want int
	}{
		{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{[][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}, 4},
	}
	for _, c := range cases {
		got := MaxEnvelopes(c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MaxEnvelopes(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}
