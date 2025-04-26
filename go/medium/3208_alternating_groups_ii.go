// https://leetcode.com/problems/alternating-groups-ii/
package medium

func NumberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	result := 0

	count := 1
	lastColor := colors[0]

	for i := 0; i < n+k-1; i++ {
		idx := i % n
		if colors[idx] == lastColor {
			count = 1
			continue
		}

		count++
		if count >= k {
			result++
		}
		lastColor = colors[idx]
	}

	return result
}
