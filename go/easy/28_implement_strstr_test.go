package easy

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	cases := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	}
	for _, c := range cases {
		got := StrStr(c.haystack, c.needle)
		if got != c.want {
			t.Errorf("StrStr(%s, %s), == %d, want %d", c.haystack, c.needle, got, c.want)
		}
	}
}
