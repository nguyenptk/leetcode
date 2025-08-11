// https://leetcode.com/problems/meeting-rooms-iii/
package hard

import (
	"sort"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func MostBooked(n int, meetings [][]int) int {
	meetingCount := make([]int, n)

	// used rooms: {endTime, roomId}
	usedRooms := priorityqueue.NewWith(func(a, b interface{}) int {
		x := a.([]int)
		y := b.([]int)
		if x[0] == y[0] {
			return x[1] - y[1]
		}
		return x[0] - y[0]
	})
	unusedRooms := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	for i := 0; i < n; i++ {
		unusedRooms.Enqueue(i)
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]

		for !usedRooms.Empty() {
			val, _ := usedRooms.Peek()
			room := val.([]int)
			if room[0] > start {
				break
			}
			usedRooms.Dequeue()
			unusedRooms.Enqueue(room[1])
		}
		if !unusedRooms.Empty() {
			val, _ := unusedRooms.Dequeue()
			roomId := val.(int)
			meetingCount[roomId]++
			usedRooms.Enqueue([]int{end, roomId})
		} else {
			val, _ := usedRooms.Dequeue()
			earliest := val.([]int)
			newEnd := earliest[0] + (end - start)
			roomId := earliest[1]
			meetingCount[roomId]++
			usedRooms.Enqueue([]int{newEnd, roomId})
		}
	}

	maxCount := 0
	resultRoom := 0
	for i := 0; i < n; i++ {
		if meetingCount[i] > maxCount {
			maxCount = meetingCount[i]
			resultRoom = i
		}
	}

	return resultRoom
}
