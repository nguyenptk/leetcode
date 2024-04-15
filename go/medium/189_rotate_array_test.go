package medium

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, c := range cases {
		got := Rotate(c.in, c.k)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Rotate(%d, %d) == %d, want %d", c.in, c.k, got, c.want)
		}
	}
}
