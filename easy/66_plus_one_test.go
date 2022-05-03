package easy

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{9, 8, 9}, []int{9, 9, 0}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{5, 2, 2, 6, 5, 7, 1, 9, 0, 3, 8, 6, 8, 6, 5, 2, 1, 8, 7, 9, 8, 3, 8, 4, 7, 2, 5, 8, 9}, []int{5, 2, 2, 6, 5, 7, 1, 9, 0, 3, 8, 6, 8, 6, 5, 2, 1, 8, 7, 9, 8, 3, 8, 4, 7, 2, 5, 9, 0}},
	}
	for _, c := range cases {
		got := PlusOne(c.in)
		for i := 0; i < len(got); i++ {
			if got[i] != c.want[i] {
				t.Errorf("PlusOne(%q) == %d, want %d", fmt.Sprint(c.in), got[i], c.want[i])
			}
		}

	}

}
