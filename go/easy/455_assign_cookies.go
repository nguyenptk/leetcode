// https://leetcode.com/problems/assign-cookies/
package easy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	count := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
			count++
		}
		j++
	}

	return count
}
