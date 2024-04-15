package medium

import (
	"fmt"
	"testing"
)

func TestMaxAncestorDiff(t *testing.T) {
	cases := []struct {
		in   *TreeNode
		want int
	}{
		// test case 1
		{
			&TreeNode{ // [8,3,10,1,6,null,14,null,null,4,7,13]
				Val: 8,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 10,
					Right: &TreeNode{
						Val: 14,
						Left: &TreeNode{
							Val: 13,
						},
					},
				},
			},
			7,
		},
		// test case 2
		{
			&TreeNode{ // [1,null,2,null,0,3]
				Val: 1,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			3,
		},
	}
	for _, c := range cases {
		got := MaxAncestorDiff(c.in)
		if got != c.want {
			t.Errorf("MaxAncestorDiff(%q) == %d, want %d", fmt.Sprint(c.in), got, c.want)
		}
	}
}
