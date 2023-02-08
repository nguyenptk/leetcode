package medium

import (
	"reflect"
	"testing"
)

func TestJump(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	}
	for _, c := range cases {
		got := Jump(c.nums)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Jump(%d) == %d, want %d", c.nums, got, c.want)
		}
	}
}
