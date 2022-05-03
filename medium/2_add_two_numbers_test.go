package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		// test case 1
		{
			&ListNode{ // l1 [2,4,3]
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			&ListNode{ // l2 [5,6,4]
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			&ListNode{ // want [7,0,8]
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		// test case 2
		{
			&ListNode{ // l1 [0]
				Val: 0,
			},
			&ListNode{ // l2 [0]
				Val: 0,
			},
			&ListNode{ // want [0]
				Val: 0,
			},
		},
		// test case 3
		{
			&ListNode{ // l1 [9,9,9,9,9,9,9]
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
									},
								},
							},
						},
					},
				},
			},
			&ListNode{ // l2 [9,9,9,9]
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			},
			&ListNode{ // want [8,9,9,9,0,0,0,1]
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		got := AddTwoNumbers(c.l1, c.l2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("AddTwoNumbers(%q, %q) == %q, want %q", fmt.Sprint(c.l1), fmt.Sprint(c.l2), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
