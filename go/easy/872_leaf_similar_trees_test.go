package easy

import (
	"fmt"
	"testing"
)

func TestLeafSimilar(t *testing.T) {
	cases := []struct {
		root1 *TreeNode
		root2 *TreeNode
		want  bool
	}{
		// test case 1
		{
			&TreeNode{ // [3,5,1,6,2,9,8,null,null,7,4]
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			&TreeNode{ // [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			true,
		},
		// test case 2
		{
			&TreeNode{ // [1,2,3]
				Val: 1,
				Right: &TreeNode{
					Val: 3,
				},
				Left: &TreeNode{
					Val: 2,
				},
			},
			&TreeNode{ // [1,3,2]
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
				Left: &TreeNode{
					Val: 3,
				},
			},
			false,
		},
	}
	for _, c := range cases {
		got := LeafSimilar(c.root1, c.root2)
		if got != c.want {
			t.Errorf("LeafSimilar(%q, %q) == %t, want %t", fmt.Sprint(c.root1), fmt.Sprint(c.root2), got, c.want)
		}
	}
}
