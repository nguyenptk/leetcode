package medium

import (
	"fmt"
	"testing"
)

func TestPseudoPalindromicPaths(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want int
	}{
		// test case 1
		{
			&TreeNode{ // [2,3,1,3,1,null,1]
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			2,
		},
		// test case 2
		{
			&TreeNode{ // [2,1,1,1,3,null,null,null,null,null,1]
				Val: 2,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			1,
		},
		// test case 3
		{
			&TreeNode{ // [9]
				Val: 1,
			},
			1,
		},
	}
	for _, c := range cases {
		got := PseudoPalindromicPaths(c.root)
		if got != c.want {
			t.Errorf("PseudoPalindromicPaths(%q) == %d, want %d", fmt.Sprint(c.root), got, c.want)
		}
	}
}
