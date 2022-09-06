package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPruneTree(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want *TreeNode
	}{
		// test case 1
		{
			&TreeNode{ // [1,null,0,0,1]
				Val: 1,
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			&TreeNode{ // [1,null,0,null,1]
				Val: 1,
				Right: &TreeNode{
					Val: 0,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		// test case 2
		{
			&TreeNode{ // [1,0,1,0,0,0,1]
				Val: 1,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			&TreeNode{ // [1,null,1,null,1]
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		// test case 3
		{
			&TreeNode{ // [1,1,0,1,1,0,1,0]
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			&TreeNode{ // [1,1,0,1,1,null,1]
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, c := range cases {
		got := PruneTree(c.root)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("PruneTree(%q) == %q, want %q", fmt.Sprint(c.root), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
