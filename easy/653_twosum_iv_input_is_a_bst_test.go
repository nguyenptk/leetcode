package easy

import (
	"fmt"
	"testing"
)

func TestFindTargetBST(t *testing.T) {
	cases := []struct {
		root *TreeNode
		k    int
		want bool
	}{
		// test case 1
		{
			&TreeNode{ // [5,3,6,2,4,null,7]
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			9,
			true,
		},
		// test case 2
		{
			&TreeNode{ // [5,3,6,2,4,null,7]
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			28,
			false,
		},
	}
	for _, c := range cases {
		got := FindTargetBST(c.root, c.k)
		if got != c.want {
			t.Errorf("FindTargetBST(%q, %d) == %t, want %t", fmt.Sprint(c.root), c.k, got, c.want)
		}
	}
}
