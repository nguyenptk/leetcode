package hard

import "testing"

func TestKInversePairs(t *testing.T) {
	cases := []struct {
		n    int
		k    int
		want int
	}{
		{3, 0, 1},
		{3, 1, 2},
	}
	for _, c := range cases {
		got := KInversePairs(c.n, c.k)
		if got != c.want {
			t.Errorf("KInversePairs(%d, %d) == %d, want %d", c.n, c.k, got, c.want)
		}
	}
}
