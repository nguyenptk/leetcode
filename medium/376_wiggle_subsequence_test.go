package medium

import "testing"

func TestWiggleMaxLength(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
	}
	for _, c := range cases {
		got := WiggleMaxLength(c.in)
		if got != c.want {
			t.Errorf("WiggleMaxLength(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
