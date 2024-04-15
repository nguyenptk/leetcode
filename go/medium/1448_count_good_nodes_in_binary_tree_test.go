package medium

import (
	"fmt"
	"testing"
)

func TestGoodNodes(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want int
	}{
		// test case 1
		{
			&TreeNode{ // [3,1,4,3,null,1,5]
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			4,
		},
		// test case 2
		{
			&TreeNode{ // [3,3,null,4,2]
				Val: 3,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			3,
		},
		// test case 3
		{
			&TreeNode{ // [1]
				Val: 1,
			},
			1,
		},
	}
	for _, c := range cases {
		got := GoodNodes(c.root)
		if got != c.want {
			t.Errorf("GoodNodes(%q) == %d, want %d", fmt.Sprint(c.root), got, c.want)
		}
	}
}
