package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddOneRow(t *testing.T) {
	cases := []struct {
		root  *TreeNode
		val   int
		depth int
		want  *TreeNode
	}{
		// test case 1
		{
			&TreeNode{ // [4,2,6,3,1,5]
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
			1,
			2,
			&TreeNode{ // [4,1,1,2,null,null,6,3,1,5]
				Val: 4,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
		// test case 2
		{
			&TreeNode{ // [4,2,null,3,1]
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			1,
			3,
			&TreeNode{ // [4,2,null,1,1,3,null,null,1]
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		got := AddOneRow(c.root, c.val, c.depth)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("AddOneRow(%q, %d, %d) == %q, want %q", fmt.Sprint(c.root), c.val, c.depth, fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
