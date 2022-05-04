package easy

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
	}
	for _, c := range cases {
		got := SortArrayByParity(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SortArrayByParity(%d) == %d, want %d", c.in, got, c.want)
		}
	}

}
