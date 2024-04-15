// https://leetcode.com/problems/reduce-array-size-to-the-half/
package medium

import "sort"

type frequency struct {
	k int
	v int
}

func MinSetSize(arr []int) int {
	n := len(arr)
	mapN := map[int]int{}
	for _, v := range arr {
		mapN[v]++
	}

	sortedN := make([]frequency, 0, len(mapN))
	for k, v := range mapN {
		sortedN = append(sortedN, frequency{k: k, v: v})
	}
	sort.Slice(sortedN, func(i, j int) bool { return sortedN[i].v > sortedN[j].v })

	sum := 0
	for i := 0; i < n; i++ {
		sum += sortedN[i].v
		if sum >= n/2 {
			return i + 1
		}

	}

	return 1
}
