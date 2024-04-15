package medium

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	cases := []struct {
		in   *TreeNode
		want bool
	}{
		// test case 1
		{
			&TreeNode{ // [2,1,3]
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			true,
		},
		// test case 2
		{
			&TreeNode{ // [5,1,4,nul,nul,3,6]
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			false,
		},
	}
	for _, c := range cases {
		got := IsValidBST(c.in)
		if got != c.want {
			t.Errorf("IsValidBST(%q) == %t, want %t", fmt.Sprint(c.in), got, c.want)
		}
	}
}
