package easy

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, c := range cases {
		got := RemoveDuplicates(c.in)
		if got != c.want {
			t.Errorf("RemoveDuplicates(%d), == %d, want %d", c.in, got, c.want)
		}
	}

}
