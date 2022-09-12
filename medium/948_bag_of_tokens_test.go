package medium

import "testing"

func TestBagOfTokensScore(t *testing.T) {
	cases := []struct {
		tokens []int
		power  int
		want   int
	}{
		{[]int{100}, 50, 0},
		{[]int{100, 200}, 150, 1},
		{[]int{100, 200, 300, 400}, 200, 2},
	}
	for _, c := range cases {
		got := BagOfTokensScore(c.tokens, c.power)
		if got != c.want {
			t.Errorf("BagOfTokensScore(%d, %d) == %d, want %d", c.tokens, c.power, got, c.want)
		}
	}
}
