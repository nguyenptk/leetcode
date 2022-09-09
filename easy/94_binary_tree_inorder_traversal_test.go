package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want []int
	}{
		// test case 1
		{
			&TreeNode{ // [1,null,2,3]
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			[]int{1, 3, 2},
		},
		// test case 2
		{
			nil, // []
			[]int{},
		},
		// test case 3
		{
			&TreeNode{Val: 1}, // [1]
			[]int{1},
		},
	}
	for _, c := range cases {
		got := InorderTraversal(c.root)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("InorderTraversal(%q) == %d, want %d", fmt.Sprint(c.root), got, c.want)
		}
	}
}
