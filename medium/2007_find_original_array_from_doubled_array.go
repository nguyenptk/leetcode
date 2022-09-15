// https://leetcode.com/problems/find-original-array-from-doubled-array/
package medium

import "sort"

func FindOriginalArray(changed []int) []int {
	result := []int{}
	queue := []int{}

	sort.Slice(changed, func(i, j int) bool {
		return changed[i] < changed[j]
	})

	for _, v := range changed {
		if len(queue) > 0 && v == queue[0] {
			queue = queue[1:] // POLL
		} else {
			queue = append(queue, v*2)
			result = append(result, v)
		}
	}

	if len(queue) == 0 {
		return result
	}

	return []int{}
}
