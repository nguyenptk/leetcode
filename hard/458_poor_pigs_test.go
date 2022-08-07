package hard

import "testing"

func TestPoorPigs(t *testing.T) {
	cases := []struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
		want          int
	}{
		{1000, 15, 60, 5},
		{4, 15, 15, 2},
		{4, 15, 30, 2},
	}
	for _, c := range cases {
		got := PoorPigs(c.buckets, c.minutesToDie, c.minutesToTest)
		if got != c.want {
			t.Errorf("PoorPigs(%d, %d, %d) == %d, want %d", c.buckets, c.minutesToDie, c.minutesToTest, got, c.want)
		}
	}
}
