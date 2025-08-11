// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
package medium

import "sort"

func FindMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	result := 1
	firstEnd := points[0][1]
	for _, point := range points {
		if firstEnd < point[0] {
			result += 1
			firstEnd = point[1]
		}
	}

	return result
}
