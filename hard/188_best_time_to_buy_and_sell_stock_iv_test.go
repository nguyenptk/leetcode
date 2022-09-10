package hard

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		k      int
		prices []int
		want   int
	}{
		{2, []int{2, 4, 1}, 2},
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
		{2, []int{}, 0},
	}
	for _, c := range cases {
		got := MaxProfit(c.k, c.prices)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MaxProfit(%d, %d) == %d, want %d", c.k, c.prices, got, c.want)
		}
	}
}
