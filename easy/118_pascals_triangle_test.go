package easy

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		in   int
		want [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}
	for _, c := range cases {
		got := Generate(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Generate(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
