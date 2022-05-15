package medium

import (
	"reflect"
	"testing"
)

func TestNetworkDelayTime(t *testing.T) {
	cases := []struct {
		times [][]int
		n     int
		k     int
		want  int
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{[][]int{{1, 2, 1}}, 2, 1, 1},
		{[][]int{{1, 2, 1}}, 2, 2, -1},
	}
	for _, c := range cases {
		got := NetworkDelayTime(c.times, c.n, c.k)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("NetworkDelayTime(%d) == %d, want %d", c.k, got, c.want)
		}
	}
}
