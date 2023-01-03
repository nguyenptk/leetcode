package easy

import (
	"testing"
)

func TestMinDeletionSize(t *testing.T) {
	cases := []struct {
		strs []string
		want int
	}{
		{[]string{"cba", "daf", "ghi"}, 1},
		{[]string{"a", "b"}, 0},
		{[]string{"zyx", "wvu", "tsr"}, 3},
	}
	for _, c := range cases {
		got := MinDeletionSize(c.strs)
		if got != c.want {
			t.Errorf("MinDeletionSize(%s) == %d, want %d", c.strs, got, c.want)
		}
	}
}
