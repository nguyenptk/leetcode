package medium

import (
	"reflect"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	cases := []struct {
		rooms [][]int
		want  bool
	}{
		{
			[][]int{{1}, {2}, {3}, {}},
			true,
		},
		{
			[][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			false,
		},
	}
	for _, c := range cases {
		got := CanVisitAllRooms(c.rooms)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CanVisitAllRooms(%d) == %t, want %t", c.rooms, got, c.want)
		}
	}
}
