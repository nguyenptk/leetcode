package hard

import (
	"fmt"
	"testing"
)

func TestMinCameraCover(t *testing.T) {
	cases := []struct {
		in   *TreeNode
		want int
	}{
		// test case 1
		{
			&TreeNode{ // [0,0,nul,0,0]
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
			1,
		},
		// test case 2
		{
			&TreeNode{ // [0,0,null,0,null,0,null,null,0]
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
							Right: &TreeNode{
								Val: 0,
							},
							// left: null
						},
						// right: null
					},
					// right: null
				},
			},
			2,
		},
	}
	for _, c := range cases {
		got := MinCameraCover(c.in)
		if got != c.want {
			t.Errorf("MinCameraCover(%q) == %d, want %d", fmt.Sprint(c.in), got, c.want)
		}
	}
}
