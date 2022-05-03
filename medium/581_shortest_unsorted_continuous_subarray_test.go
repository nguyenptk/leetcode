package medium

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{1}, 0},
		{[]int{1, 1, 3, 8, 8, 2, 2, 2}, 6},
	}
	for _, c := range cases {
		got := FindUnsortedSubarray(c.in)
		if got != c.want {
			t.Errorf("FindUnsortedSubarray(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
