package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	cases := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		// test case 1
		{
			&ListNode{ // listA [4,1,8,4,5]
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 8,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			&ListNode{ // listB [5,6,1,8,4,5]
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			nil, // listA+listB == listB+listA
		},
		// test case 2
		{
			&ListNode{ // listA [1,9,1,2,4]
				Val: 1,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
			&ListNode{ // listB [3,2,4]
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			nil, // listA+listB == listB+listA
		},
		// test case 3
		{
			&ListNode{ // listA [2,6,4]
				Val: 2,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			&ListNode{ // listB [1,5]
				Val: 1,
				Next: &ListNode{
					Val: 5,
				},
			},
			nil, // No intersection
		},
	}
	for _, c := range cases {
		got := GetIntersectionNode(c.list1, c.list2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("GetIntersectionNode(%q, %q) == %q, want %q", fmt.Sprint(c.list1), fmt.Sprint(c.list2), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
