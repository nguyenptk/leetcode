package hard

import (
	"reflect"
	"testing"
)

func TestCriticalConnections(t *testing.T) {
	cases := []struct {
		n           int
		connections [][]int
		want        [][]int
	}{
		{4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}, [][]int{{1, 3}}},
		{2, [][]int{{0, 1}}, [][]int{{0, 1}}},
	}
	for _, c := range cases {
		got := CriticalConnections(c.n, c.connections)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CriticalConnections(%d, %d) == %d, want %d", c.n, c.connections, got, c.want)
		}
	}
}
