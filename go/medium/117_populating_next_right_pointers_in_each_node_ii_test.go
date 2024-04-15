package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	cases := []struct {
		in   *Node
		want *Node
	}{
		// test case 1
		{
			&Node{ // [1,2,3,4,5,null,7]
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 3,
					Right: &Node{
						Val: 7,
					},
				},
			},
			&Node{ // [1,#,2,3,#,4,5,7,#]
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
						Next: &Node{
							Val: 5,
							Next: &Node{
								Val: 7,
							},
						},
					},
					Right: &Node{
						Val: 5,
						Next: &Node{
							Val: 7,
						},
					},
					Next: &Node{
						Val: 3,
						Right: &Node{
							Val: 7,
						},
					},
				},
				Right: &Node{
					Val: 3,
					Right: &Node{
						Val: 7,
					},
				},
			},
		},
		// test case 2
		{
			&Node{},
			&Node{},
		},
	}
	for _, c := range cases {
		got := Connect(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Connect(%q) == %q, want %q", fmt.Sprint(c.in), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
