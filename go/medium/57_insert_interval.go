package medium

func Insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	result := [][]int{}
	i := 0

	// Check overlapping
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Check overlapping and merge
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	result = append(result, newInterval)

	// Check non-overlapping after merge
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
