package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	cases := []struct {
		in   *TreeNode
		want []float64
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
			[]float64{3.00000, 14.50000, 11.00000},
		},
		// test case 2
		{
			&TreeNode{ // [3,9,20,15,7]
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
			[]float64{3.00000, 14.50000, 11.00000},
		},
	}
	for _, c := range cases {
		got := AverageOfLevels(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("AverageOfLevels(%q) == %f, want %f", fmt.Sprint(c.in), got, c.want)
		}
	}
}
