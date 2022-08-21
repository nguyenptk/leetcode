package hard

import (
	"reflect"
	"testing"
)

func TestMovesToStamp(t *testing.T) {
	cases := []struct {
		stamp  string
		target string
		want   []int
	}{
		// Test case 1
		{
			"abc",
			"ababc",
			[]int{0, 2},
		},
		// Test case 2
		{
			"abca",
			"aabcaca",
			[]int{0, 3, 1},
		},
	}
	for _, c := range cases {
		got := MovesToStamp(c.stamp, c.target)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MovesToStamp(%s, %s) == %d, want %d", c.stamp, c.target, got, c.want)
		}
	}
}
