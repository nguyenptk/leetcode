package medium

import "testing"

func TestFindLength(t *testing.T) {
	cases := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	}
	for _, c := range cases {
		got := FindLength(c.nums1, c.nums2)
		if got != c.want {
			t.Errorf("FindLength(%d, %d) == %d, want %d", c.nums1, c.nums2, got, c.want)
		}
	}
}
