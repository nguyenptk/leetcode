package hard

import (
	"testing"
)

func TestMinDifficulty(t *testing.T) {
	cases := []struct {
		jobDifficulty []int
		d             int
		want          int
	}{
		{[]int{6, 5, 4, 3, 2, 1}, 2, 7},
		{[]int{9, 9, 9}, 4, -1},
		{[]int{1, 1, 1}, 3, 3},
	}
	for _, c := range cases {
		got := MinDifficulty(c.jobDifficulty, c.d)
		if got != c.want {
			t.Errorf("MinDifficulty(%d, %d) == %d, want %d", c.jobDifficulty, c.d, got, c.want)
		}
	}
}
