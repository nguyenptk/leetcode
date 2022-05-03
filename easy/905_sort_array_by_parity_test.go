package easy

import (
	"fmt"
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
		for i := 0; i < len(got); i++ {
			if got[i] != c.want[i] {
				t.Errorf("SortArrayByParity(%q) == %d, want %d", fmt.Sprint(c.in), got[i], c.want[i])
			}
		}
	}

}
