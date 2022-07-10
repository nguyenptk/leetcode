package easy

import (
	"reflect"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, c := range cases {
		got := MinCostClimbingStairs(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MinCostClimbingStairs(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
