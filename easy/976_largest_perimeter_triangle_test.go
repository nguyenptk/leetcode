package easy

import (
	"testing"
)

func TestLargestPerimeter(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 1, 2}, 5},
		{[]int{1, 2, 1}, 0},
	}
	for _, c := range cases {
		got := LargestPerimeter(c.in)
		if got != c.want {
			t.Errorf("LargestPerimeter(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
