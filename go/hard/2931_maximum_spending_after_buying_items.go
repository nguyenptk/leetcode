// https://leetcode.com/problems/maximum-spending-after-buying-items/
package hard

import "github.com/emirpasic/gods/queues/priorityqueue"

type Item struct {
	val  int
	shop int
	idx  int
}

func MaxSpending(values [][]int) int64 {
	m, n := len(values), len(values[0])

	minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(*Item).val - b.(*Item).val
	})

	for i := 0; i < m; i++ {
		idx := n - 1
		minHeap.Enqueue(&Item{
			val:  values[i][idx],
			shop: i,
			idx:  idx,
		})
	}

	total := int64(0)
	for day := 1; day <= m*n; day++ {
		node, _ := minHeap.Dequeue()
		curr := node.(*Item)
		total += int64(curr.val * day)

		if curr.idx-1 >= 0 {
			nxtIdx := curr.idx - 1
			minHeap.Enqueue(&Item{
				val:  values[curr.shop][nxtIdx],
				shop: curr.shop,
				idx:  nxtIdx,
			})
		}
	}

	return total
}

// Time Complexity: O(m * n * log m)
// Space Complexity: O(m)
