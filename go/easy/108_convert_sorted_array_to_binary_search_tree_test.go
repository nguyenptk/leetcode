package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	cases := []struct {
		in   []int
		want *TreeNode
	}{
		// test case 1
		{
			[]int{-10, -3, 0, 5, 9},
			&TreeNode{ // [0,-10,5,null,-3,null,9]
				Val: 0,
				Left: &TreeNode{
					Val: -10,
					Right: &TreeNode{
						Val: -3,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
		},
		// test case 2
		{
			[]int{1, 3},
			&TreeNode{ // [1,null,3]
				Val: 1,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}
	for _, c := range cases {
		got := SortedArrayToBST(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SortedArrayToBST(%d) == %q, want %q", c.in, fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
