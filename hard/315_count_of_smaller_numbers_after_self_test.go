package hard

import (
	"reflect"
	"testing"
)

func TestCountSmaller(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{-1}, []int{0}},
		{[]int{-1, -1}, []int{0, 0}},
		{[]int{0, 2, 1}, []int{0, 1, 0}},
	}
	for _, c := range cases {
		got := CountSmaller(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CountSmaller(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
