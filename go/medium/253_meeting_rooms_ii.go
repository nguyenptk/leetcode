// https://leetcode.com/problems/meeting-rooms-ii/
package medium

import (
	"sort"

	"github.com/emirpasic/gods/v2/queues/priorityqueue"
)

func MinMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	allocator := priorityqueue.NewWith(func(i, j int) int {
		return i - j
	})
	allocator.Enqueue(intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		top, _ := allocator.Peek()
		if intervals[i][0] >= top {
			allocator.Dequeue()
		}
		allocator.Enqueue(intervals[i][1])
	}

	return allocator.Size()
}

// Time Complexity: O(NlogN)
// Space Complexity: O(N)
