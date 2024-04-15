package easy

import "testing"

func TestMaximum69Number(t *testing.T) {
	cases := []struct {
		num  int
		want int
	}{
		{9669, 9969},
		{9996, 9999},
		{9999, 9999},
	}
	for _, c := range cases {
		got := Maximum69Number(c.num)
		if got != c.want {
			t.Errorf("Maximum69Number(%d) == %d, want %d", c.num, got, c.want)
		}
	}
}
