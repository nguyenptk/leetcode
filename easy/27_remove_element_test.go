package easy

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		in   []int
		val  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for _, c := range cases {
		got := RemoveElement(c.in, c.val)
		if got != c.want {
			t.Errorf("RemoveElement(%d, %d), == %d, want %d", c.in, c.val, got, c.want)
		}
	}

}
