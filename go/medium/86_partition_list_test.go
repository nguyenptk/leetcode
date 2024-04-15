package medium

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		head *ListNode
		x    int
		want *ListNode
	}{
		// test case 1
		{
			&ListNode{ // [1,4,3,2,5,2]
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 2,
								},
							},
						},
					},
				},
			},
			3,
			&ListNode{ // [1,2,2,4,3,5]
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
		{
			&ListNode{ // [2,1]
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
			2,
			&ListNode{ // [1,2]
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}
	for _, c := range cases {
		got := Partition(c.head, c.x)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Partition(%q, %d) == %q, want %q", fmt.Sprint(c.head), c.x, fmt.Sprint(got), fmt.Sprint(c.want))
		}
	}
}
