package easy

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, c := range cases {
		got := SearchInsert(c.in, c.target)
		if got != c.want {
			t.Errorf("SearchInsert(%d, %d), == %d, want %d", c.in, c.target, got, c.want)
		}
	}

}
