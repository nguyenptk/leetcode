package medium

import "testing"

func TestNumberOfWeakCharacters(t *testing.T) {
	cases := []struct {
		properties [][]int
		want       int
	}{
		{[][]int{{5, 5}, {6, 3}, {3, 6}}, 0},
		{[][]int{{2, 2}, {3, 3}}, 1},
		{[][]int{{1, 5}, {10, 4}, {4, 3}}, 1},
	}
	for _, c := range cases {
		got := NumberOfWeakCharacters(c.properties)
		if got != c.want {
			t.Errorf("NumberOfWeakCharacters(%d) == %d, want %d", c.properties, got, c.want)
		}
	}
}
