package medium

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}
	for _, c := range cases {
		got := SearchRange(c.nums, c.target)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SearchRange(%d, %d) == %d, want %d", c.nums, c.target, got, c.want)
		}
	}
}
