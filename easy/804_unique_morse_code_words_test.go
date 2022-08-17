package easy

import (
	"reflect"
	"testing"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
		{[]string{"a"}, 1},
	}
	for _, c := range cases {
		got := UniqueMorseRepresentations(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("UniqueMorseRepresentations(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}
