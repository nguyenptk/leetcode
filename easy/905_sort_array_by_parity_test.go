package easy

import (
	"fmt"
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
			t.Errorf("SortArrayByParity(%q) == %q, want %q", fmt.Sprint(c.in), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}

}
