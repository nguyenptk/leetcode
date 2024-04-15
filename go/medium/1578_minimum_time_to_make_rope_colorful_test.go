package medium

import "testing"

func TestMinCost(t *testing.T) {
	cases := []struct {
		colors     string
		neededTime []int
		want       int
	}{
		{"abaac", []int{1, 2, 3, 4, 5}, 3},
		{"abc", []int{1, 2, 3}, 0},
		{"aabaa", []int{1, 2, 3, 4, 1}, 2},
	}
	for _, c := range cases {
		got := MinCost(c.colors, c.neededTime)
		if got != c.want {
			t.Errorf("MinCost(%s, %d) == %d, want %d", c.colors, c.neededTime, got, c.want)
		}
	}
}
