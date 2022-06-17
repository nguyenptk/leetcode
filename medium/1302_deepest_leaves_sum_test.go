package medium

import (
	"fmt"
	"testing"
)

func TestDeepestLeavesSum(t *testing.T) {
	cases := []struct {
		in   *TreeNode
		want int
	}{
		// test case 1
		{
			&TreeNode{ // [1,2,3,4,5,null,6,7,null,null,null,null,8]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 7,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			15,
		},
		{
			&TreeNode{ // [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
				Val: 6,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 9,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			19,
		},
	}
	for _, c := range cases {
		got := DeepestLeavesSum(c.in)
		if got != c.want {
			t.Errorf("DeepestLeavesSum(%q) == %d, want %d", fmt.Sprint(c.in), got, c.want)
		}
	}
}
