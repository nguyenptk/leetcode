package easy

import (
	"testing"
)

func TestIsToeplitzMatrix(t *testing.T) {
	cases := []struct {
		matrix [][]int
		want   bool
	}{
		{[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true},
		{[][]int{{1, 2}, {2, 2}}, false},
	}
	for _, c := range cases {
		got := IsToeplitzMatrix(c.matrix)
		if got != c.want {
			t.Errorf("IsToeplitzMatrix(%d) == %t, want %t", c.matrix, got, c.want)
		}
	}
}
