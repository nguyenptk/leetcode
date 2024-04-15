package medium

import (
	"testing"
)

func TestMinimumRounds(t *testing.T) {
	cases := []struct {
		tasks []int
		want  int
	}{
		{[]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, 4},
		{[]int{2, 3, 3}, -1},
	}
	for _, c := range cases {
		got := MinimumRounds(c.tasks)
		if got != c.want {
			t.Errorf("MinimumRounds(%d) == %d, want %d", c.tasks, got, c.want)
		}
	}
}
