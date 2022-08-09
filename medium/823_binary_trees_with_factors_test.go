package medium

import "testing"

func TestNumFactoredBinaryTrees(t *testing.T) {
	cases := []struct {
		arr  []int
		want int
	}{
		{[]int{2, 4}, 3},
		{[]int{2, 4, 5, 10}, 7},
	}
	for _, c := range cases {
		got := NumFactoredBinaryTrees(c.arr)
		if got != c.want {
			t.Errorf("NumFactoredBinaryTrees(%d) == %d, want %d", c.arr, got, c.want)
		}
	}
}
