package medium

import "testing"

func TestFind132pattern(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{3, 1, 4, 2}, true},
		{[]int{-3, 1, 2, 0}, true},
	}
	for _, c := range cases {
		got := Find132pattern(c.in)
		if got != c.want {
			t.Errorf("Find132pattern(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}
