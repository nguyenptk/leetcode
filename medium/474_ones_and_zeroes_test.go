package medium

import "testing"

func TestFindMaxForm(t *testing.T) {
	cases := []struct {
		strs []string
		m    int
		n    int
		want int
	}{
		{[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{[]string{"10", "0001", "111001", "1", "0"}, 1, 1, 2},
		{[]string{"10", "0", "1"}, 1, 1, 2},
	}
	for _, c := range cases {
		got := FindMaxForm(c.strs, c.m, c.n)
		if got != c.want {
			t.Errorf("FindMaxForm(%s, %d, %d) == %d, want %d", c.strs, c.m, c.n, got, c.want)
		}
	}
}
