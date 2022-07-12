package medium

import "testing"

func TestMakesquare(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 1, 2, 2, 2}, true},
		{[]int{3, 3, 3, 3, 4}, false},
		{[]int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2}, true},
		{[]int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2}, true},
	}
	for _, c := range cases {
		got := Makesquare(c.in)
		if got != c.want {
			t.Errorf("Makesquare(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}
