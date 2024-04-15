package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	cases := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		// test case 1
		{
			&TreeNode{ // [3,5,1,6,2,0,8,null,null,7,4]
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
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			&TreeNode{ // [5]
				Val: 5,
			},
			&TreeNode{ // [1]
				Val: 1,
			},
			&TreeNode{ // [3,5,1,6,2,0,8,null,null,7,4]
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
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		},
		// test case 2
		{
			&TreeNode{ // [3,5,1,6,2,0,8,null,null,7,4]
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
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			&TreeNode{ // [5]
				Val: 5,
			},
			&TreeNode{ // [4]
				Val: 4,
			},
			&TreeNode{ // [5,6,2,null,null,7,4]
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
		},
		// test case 3
		{
			&TreeNode{ // [1,2]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			&TreeNode{ // [1]
				Val: 1,
			},
			&TreeNode{ // [2]
				Val: 2,
			},
			&TreeNode{ // [1,2]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
	}
	for _, c := range cases {
		got := LowestCommonAncestor(c.root, c.p, c.q)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("LowestCommonAncestor(%q, %q, %q) == %q, want %q", fmt.Sprint(c.root), fmt.Sprint(c.p), fmt.Sprint(c.q), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
