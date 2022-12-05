package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	cases := []struct {
		list *ListNode
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
			&ListNode{ // [3,4,5]
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
		// test case 2
		{
			&ListNode{ // list [1,2,3,4,5,6]
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
			&ListNode{ // list [4,5,6]
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
	}
	for _, c := range cases {
		got := MiddleNode(c.list)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MiddleNode(%q) == %q, want %q", fmt.Sprint(c.list), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
