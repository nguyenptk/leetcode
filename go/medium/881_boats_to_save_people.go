// https://leetcode.com/problems/boats-to-save-people/
package medium

import "sort"

func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	result := 0

	l := 0
	r := len(people) - 1
	for l <= r {
		result++
		if people[l]+people[r] <= limit {
			l++
		}
		r--
	}

	return result
}
