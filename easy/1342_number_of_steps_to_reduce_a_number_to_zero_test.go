package easy

import (
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{14, 6},
		{8, 4},
		{123, 12},
	}
	for _, c := range cases {
		got := NumberOfSteps(c.in)
		if got != c.want {
			t.Errorf("NumberOfSteps(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
