package hard

import (
	"testing"
)

func TestOrderlyQueue(t *testing.T) {
	cases := []struct {
		s    string
		k    int
		want string
	}{
		// Test case 1
		{
			"cba",
			1,
			"acb",
		},
		// Test case 2
		{
			"baaca",
			3,
			"aaabc",
		},
	}
	for _, c := range cases {
		got := OrderlyQueue(c.s, c.k)
		if got != c.want {
			t.Errorf("OrderlyQueue(%s, %d) == %s, want %s", c.s, c.k, got, c.want)
		}
	}
}
