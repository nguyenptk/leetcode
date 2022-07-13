package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	cases := []struct {
		in   *TreeNode
		want [][]int
	}{
		// test case 1
		{
			&TreeNode{ // [3,9,20,null,null,15,7]
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		// test case 2
		{
			&TreeNode{ // [1]
				Val: 1,
			},
			[][]int{{1}},
		},
		// test case 3
		{
			nil,
			[][]int{},
		},
	}
	for _, c := range cases {
		got := LevelOrder(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("LevelOrder(%q) == %d, want %d", fmt.Sprint(c.in), got, c.want)
		}
	}
}
