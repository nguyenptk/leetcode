// https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
package medium

import "sort"

func MaxPiece(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Slice(horizontalCuts, func(i, j int) bool {
		return horizontalCuts[i] < horizontalCuts[j]
	})

	sort.Slice(verticalCuts, func(i, j int) bool {
		return verticalCuts[i] < verticalCuts[j]
	})

	var maxH int
	if horizontalCuts[0] > h-horizontalCuts[len(horizontalCuts)-1] {
		maxH = horizontalCuts[0]
	} else {
		maxH = h - horizontalCuts[len(horizontalCuts)-1]
	}

	var maxV int
	if verticalCuts[0] > w-verticalCuts[len(verticalCuts)-1] {
		maxV = verticalCuts[0]
	} else {
		maxV = w - verticalCuts[len(verticalCuts)-1]
	}

	for i := 1; i < len(horizontalCuts); i++ {
		if maxH < horizontalCuts[i]-horizontalCuts[i-1] {
			maxH = horizontalCuts[i] - horizontalCuts[i-1]
		}
	}

	for i := 1; i < len(verticalCuts); i++ {
		if maxV < verticalCuts[i]-verticalCuts[i-1] {
			maxV = verticalCuts[i] - verticalCuts[i-1]
		}
	}

	return maxH * maxV % 1000000007
}
