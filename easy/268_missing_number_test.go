package easy

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}
	for _, c := range cases {
		got := MissingNumber(c.in)
		if got != c.want {
			t.Errorf("MissingNumber(%d) == %d, want %d", c.in, got, c.want)
		}
	}

}
