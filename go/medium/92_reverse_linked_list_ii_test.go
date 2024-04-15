package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	cases := []struct {
		head  *ListNode
		left  int
		right int
		want  *ListNode
	}{
		// test case 1
		{
			&ListNode{ // [1,2,3,4,5]
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			2,
			4,
			&ListNode{ // [1,4,3,2,5]
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			&ListNode{ // [1]
				Val: 1,
			},
			1,
			1,
			&ListNode{ // [1]
				Val: 1,
			},
		},
	}
	for _, c := range cases {
		got := ReverseBetween(c.head, c.left, c.right)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ReverseBetween(%q, %d, %d) == %q, want %q", fmt.Sprint(c.head), c.left, c.right, fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
