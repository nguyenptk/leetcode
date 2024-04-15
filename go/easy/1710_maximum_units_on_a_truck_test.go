package easy

import (
	"reflect"
	"testing"
)

func TestMaximumUnits(t *testing.T) {
	cases := []struct {
		boxTypes  [][]int
		truckSize int
		want      int
	}{
		{[][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8},
		{[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91},
	}
	for _, c := range cases {
		got := MaximumUnits(c.boxTypes, c.truckSize)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MaximumUnits(%d, %d) == %d, want %d", c.boxTypes, c.truckSize, got, c.want)
		}
	}
}
