package easy

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	cases := []struct {
		in   [][]int
		want [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	}
	for _, c := range cases {
		got := Transpose(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Transpose(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
