package hard

import (
	"testing"
)

func TestMaxPerformance(t *testing.T) {
	cases := []struct {
		n          int
		speed      []int
		efficiency []int
		k          int
		want       int
	}{
		{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2, 60},
		{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3, 68},
		{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4, 72},
	}
	for _, c := range cases {
		got := MaxPerformance(c.n, c.speed, c.efficiency, c.k)
		if got != c.want {
			t.Errorf("MaxPerformance(%d, %d, %d, %d) == %d, want %d", c.n, c.speed, c.efficiency, c.k, got, c.want)
		}
	}
}
