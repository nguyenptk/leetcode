// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
package medium

import "sort"

// import "github.com/emirpasic/gods/v2/queues/priorityqueue"

// type Element struct {
//     val int
//     row int
//     col int
// }

// func KthSmallest(matrix [][]int, k int) int {
//     n := len(matrix)

//     minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
//         return a.(Element).val - b.(Element).val
//     })

//     for i := 0; i < min(n, k); i++ {
//         minHeap.Enqueue(Element{
//             val: matrix[i][0],
//             row: i,
//             col: 0,
//         })
//     }

//     var curr Element
//     for k > 0 {
//         e, _ := minHeap.Dequeue()
//         curr = e.(Element)

//         if curr.col+1 < n {
//             minHeap.Enqueue(Element{
//                 val: matrix[curr.row][curr.col+1],
//                 row: curr.row,
//                 col: curr.col+1,
//             })
//         }
//         k--
//     }

//     return curr.val
// }

func KthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		m := l + (r-l)/2
		count := 0

		for _, row := range matrix {
			idx := sort.Search(n, func(i int) bool { return row[i] > m })
			count += idx
		}

		if count >= k {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
