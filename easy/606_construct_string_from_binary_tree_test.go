package easy

import (
	"fmt"
	"testing"
)

func TestTree2str(t *testing.T) {
	cases := []struct {
		in   *TreeNode
		want string
	}{
		// test case 1
		{
			&TreeNode{ // [1,2,3,4]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			"1(2(4))(3)",
		},
		// test case 2
		{
			&TreeNode{ // [1,2,3,null,4]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			"1(2()(4))(3)",
		},
	}
	for _, c := range cases {
		got := Tree2str(c.in)
		if got != c.want {
			t.Errorf("Tree2str(%q) == %s, want %s", fmt.Sprint(c.in), got, c.want)
		}
	}
}
