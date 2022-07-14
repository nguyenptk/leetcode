package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	cases := []struct {
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		// test case 1
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			&TreeNode{ // [3,9,20,null,null,15,7]
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		// test case 2
		{
			[]int{-1},
			[]int{-1},
			&TreeNode{ // [-1]
				Val: -1,
			},
		},
	}
	for _, c := range cases {
		got := BuildTree(c.preorder, c.inorder)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("BuildTree(%d, %d) == %q, want %q", c.preorder, c.inorder, fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
