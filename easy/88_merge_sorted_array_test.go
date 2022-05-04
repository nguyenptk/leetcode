package easy

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		// test case 1
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		// test case 2
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, c := range cases {
		got := Merge(c.nums1, c.m, c.nums2, c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Merge(%d, %d, %d, %d), == %d, want %d", c.nums1, c.m, c.nums2, c.n, got, c.want)
		}
	}

}
