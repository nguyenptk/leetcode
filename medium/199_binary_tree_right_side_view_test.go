package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	cases := []struct {
		in   *TreeNode
		want []int
	}{
		// test case 1
		{
			&TreeNode{ // [1,2,3,null,5,null,4]
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			[]int{1, 3, 4},
		},
		// test case 2
		{
			&TreeNode{ // [1,null,3]
				Val: 1,
				Right: &TreeNode{
					Val: 3,
				},
			},
			[]int{1, 3},
		},
		// test case 3
		{
			nil, // []
			[]int{},
		},
	}
	for _, c := range cases {
		got := RightSideView(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RightSideView(%q) == %d, want %d", fmt.Sprint(c.in), got, c.want)
		}
	}
}
