package medium

import "testing"

func TestNumDecodings(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
	}
	for _, c := range cases {
		got := NumDecodings(c.s)
		if got != c.want {
			t.Errorf("NumDecodings(%s) == %d, want %d", c.s, got, c.want)
		}
	}
}
