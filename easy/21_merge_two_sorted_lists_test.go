package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		// test case 1
		{
			&ListNode{ // list1 [1,2,4]
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			&ListNode{ // list2 [1,3,4]
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			&ListNode{ // list1 [1,1,2,3,4,4]
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		got := MergeTwoLists(c.list1, c.list2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MergeTwoLists(%q, %q) == %q, want %q", fmt.Sprint(c.list1), fmt.Sprint(c.list2), fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
