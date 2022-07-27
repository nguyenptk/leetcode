package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want *TreeNode
	}{
		// test case 1
		{
			&TreeNode{ // [1,2,5,3,4,null,6]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			&TreeNode{ // [1,null,2,null,3,null,4,null,5,null,6]
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
		// test case 2
		{
			&TreeNode{}, // []
			&TreeNode{}, // []
		},
		// test case 3
		{
			&TreeNode{Val: 0}, // [0]
			&TreeNode{Val: 0}, // [0]
		},
	}
	for _, c := range cases {
		Flatten(c.root)
		got := c.root
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Flatten(%q) == %q, want %q", fmt.Sprint(c.root), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
