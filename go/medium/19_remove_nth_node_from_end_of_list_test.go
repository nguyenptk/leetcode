package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		head *ListNode
		n    int
		want *ListNode
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
			&ListNode{ // [1,2,3,5]
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
		// test case 2
		{
			&ListNode{ // [1]
				Val: 1,
			},
			1,
			nil, // []
		},
		// test case 3
		{
			&ListNode{ // [1,2]
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			1,
			&ListNode{ // [1]
				Val: 1,
			},
		},
	}
	for _, c := range cases {
		got := RemoveNthFromEnd(c.head, c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RemoveNthFromEnd(%q, %d) == %q, want %q", fmt.Sprint(c.head), c.n, fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
