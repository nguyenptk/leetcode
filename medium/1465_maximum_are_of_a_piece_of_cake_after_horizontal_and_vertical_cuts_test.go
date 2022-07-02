package medium

import "testing"

func TestMaxPiece(t *testing.T) {
	cases := []struct {
		h              int
		w              int
		horizontalCuts []int
		verticalCuts   []int
		want           int
	}{
		{5, 4, []int{1, 2, 4}, []int{1, 3}, 4},
		{5, 4, []int{3, 1}, []int{1}, 6},
		{5, 4, []int{3}, []int{3}, 9},
	}
	for _, c := range cases {
		got := MaxPiece(c.h, c.w, c.horizontalCuts, c.verticalCuts)
		if got != c.want {
			t.Errorf("MaxPiece(%d, %d, %d, %d) == %d, want %d", c.h, c.w, c.horizontalCuts, c.verticalCuts, got, c.want)
		}
	}
}
