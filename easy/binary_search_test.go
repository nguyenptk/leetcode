package easy

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	}
	for _, c := range cases {
		got := Search(c.in, c.target)
		if got != c.want {
			t.Errorf("Search(%q, %d) == %d, want %d", fmt.Sprint(c.in), c.target, got, c.want)
		}
	}

}
