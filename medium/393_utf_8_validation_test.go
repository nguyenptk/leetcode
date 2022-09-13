package medium

import "testing"

func TestValidUtf8(t *testing.T) {
	cases := []struct {
		data []int
		want bool
	}{
		{[]int{197, 130, 1}, true},
		{[]int{235, 140, 4}, false},
	}
	for _, c := range cases {
		got := ValidUtf8(c.data)
		if got != c.want {
			t.Errorf("ValidUtf8(%d) == %t, want %t", c.data, got, c.want)
		}
	}
}
