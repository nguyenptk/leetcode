package easy

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
	}
	for _, c := range cases {
		got := RunningSum(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RunningSum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
