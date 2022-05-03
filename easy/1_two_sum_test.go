package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, c := range cases {
		got := TwoSum(c.in, c.target)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("TwoSum(%q, %d) == %q, want %q", fmt.Sprint(c.in), c.target, fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}

}
