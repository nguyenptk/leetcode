package medium

import (
	"reflect"
	"testing"
)

func TestFindOriginalArray(t *testing.T) {
	cases := []struct {
		changed []int
		want    []int
	}{
		{[]int{1, 3, 4, 2, 6, 8}, []int{1, 3, 4}},
		{[]int{6, 3, 0, 1}, []int{}},
		{[]int{1}, []int{}},
	}
	for _, c := range cases {
		got := FindOriginalArray(c.changed)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FindOriginalArray(%d) == %d, want %d", c.changed, got, c.want)
		}
	}
}
