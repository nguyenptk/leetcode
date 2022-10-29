package hard

import (
	"testing"
)

func TestEarliestFullBloom(t *testing.T) {
	cases := []struct {
		plantTime []int
		growTime  []int
		want      int
	}{
		{[]int{1, 4, 3}, []int{2, 3, 1}, 9},
		{[]int{1, 2, 3, 2}, []int{2, 1, 2, 1}, 9},
		{[]int{1}, []int{1}, 2},
	}
	for _, c := range cases {
		got := EarliestFullBloom(c.plantTime, c.growTime)
		if got != c.want {
			t.Errorf("EarliestFullBloom(%d, %d) == %d, want %d", c.plantTime, c.growTime, got, c.want)
		}
	}
}
