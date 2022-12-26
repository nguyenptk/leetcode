package medium

import (
	"reflect"
	"testing"
)

func TestCanJump(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}
	for _, c := range cases {
		got := CanJump(c.nums)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CanJump(%d) == %t, want %t", c.nums, got, c.want)
		}
	}
}
