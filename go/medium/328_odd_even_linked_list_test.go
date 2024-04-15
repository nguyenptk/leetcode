package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	cases := []struct {
		head *ListNode
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
			&ListNode{ // [1,3,5,2,4]
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		},
		// test case 2
		{
			&ListNode{ // [2,1,3,5,6,4,7]
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 7,
									},
								},
							},
						},
					},
				},
			},
			&ListNode{ // [2,3,6,7,1,5,4]
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 4,
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
		got := OddEvenList(c.head)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("OddEvenList(%q) == %q, want %q", fmt.Sprint(c.head), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
