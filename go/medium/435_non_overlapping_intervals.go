package medium

import (
	"sort"
)

func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	result := 0
	prevEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		if start >= prevEnd {
			prevEnd = end
		} else {
			result++
		}
	}

	return result
}
