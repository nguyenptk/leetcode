package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNAryLevelOrder(t *testing.T) {
	cases := []struct {
		root *Node
		want [][]int
	}{
		// test case 1
		{
			&Node{ // [1,null,3,2,4,null,5,6]
				Val: 1,
				Children: []*Node{
					{
						Val: 3,
						Children: []*Node{
							{Val: 5},
							{Val: 6},
						},
					},
					{Val: 2},
					{Val: 4},
				},
			},
			[][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		// test case 2
		{
			&Node{ // [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
				Val: 1,
				Children: []*Node{
					{Val: 2},
					{
						Val: 3,
						Children: []*Node{
							{Val: 6},
							{
								Val: 7,
								Children: []*Node{
									{
										Val: 11,
										Children: []*Node{
											{Val: 14},
										},
									},
								},
							},
						},
					},
					{
						Val: 4,
						Children: []*Node{
							{
								Val: 8,
								Children: []*Node{
									{Val: 12},
								},
							},
						},
					},
					{
						Val: 5,
						Children: []*Node{
							{
								Val: 9,
								Children: []*Node{
									{Val: 13},
								},
							},
							{Val: 10},
						},
					},
				},
			},
			[][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	}
	for _, c := range cases {
		got := NAryLevelOrder(c.root)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("NAryLevelOrder(%q) == %d, want %d", fmt.Sprint(c.root), got, c.want)
		}
	}
}
