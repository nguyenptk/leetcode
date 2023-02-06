package easy

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	cases := []struct {
		nums []int
		n    int
		want []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}
	for _, c := range cases {
		got := Shuffle(c.nums, c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Shuffle(%d, %d) == %d, want %d", c.nums, c.n, got, c.want)
		}
	}
}
