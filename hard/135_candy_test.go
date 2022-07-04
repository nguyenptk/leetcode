package hard

import "testing"

func TestCandy(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
	}
	for _, c := range cases {
		got := Candy(c.in)
		if got != c.want {
			t.Errorf("Candy(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
