package easy

import (
	"fmt"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	cases := []struct {
		root      *TreeNode
		targetSum int
		want      bool
	}{
		// test case 1
		{
			&TreeNode{ // [5,4,8,11,null,13,4,7,2,null,null,null,1]
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			22,
			true,
		},
		// test case 2
		{
			&TreeNode{ // [1,2,3]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			5,
			false,
		},
		// test case 3
		{
			nil, // []
			0,
			false,
		},
	}
	for _, c := range cases {
		got := HasPathSum(c.root, c.targetSum)
		if got != c.want {
			t.Errorf("HasPathSum(%q, %d) == %t, want %t", fmt.Sprint(c.root), c.targetSum, got, c.want)
		}
	}
}
