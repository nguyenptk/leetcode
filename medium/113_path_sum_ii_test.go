package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	cases := []struct {
		root      *TreeNode
		targetSum int
		want      [][]int
	}{
		// test case 1
		{
			&TreeNode{ // [5,4,8,11,null,13,4,7,2,null,null,5,1]
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
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			22,
			[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
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
			[][]int{},
		},
		// test case 3
		{
			&TreeNode{ // [1,2]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			0,
			[][]int{},
		},
	}
	for _, c := range cases {
		got := PathSum(c.root, c.targetSum)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("PathSum(%q, %d) == %d, want %d", fmt.Sprint(c.root), c.targetSum, got, c.want)
		}
	}
}
