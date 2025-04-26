// https://leetcode.com/problems/median-of-two-sorted-arrays/
package hard

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array for binary search
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	l, r := 0, m
	var aL, aR, bL, bR float64

	for l <= r {
		i := (l + r) >> 1   // Partition for nums1
		j := (m+n+1)>>1 - i // Partition for nums2

		aL = getBoundary(nums1, i-1)
		aR = getBoundary(nums1, i)
		bL = getBoundary(nums2, j-1)
		bR = getBoundary(nums2, j)

		// Check if partition is correct
		if aL <= bR && bL <= aR {
			// Odd total length
			if (m+n)%2 == 1 {
				return math.Max(aL, bL)
			}
			// Even total length
			return (math.Max(aL, bL) + min(aR, bR)) / 2
		} else if aL > bR {
			r = i - 1 // Move left
		} else {
			l = i + 1 // Move right
		}
	}

	return 0.0
}

// Helper function to get boundaries
func getBoundary(arr []int, idx int) float64 {
	if idx < 0 {
		return math.Inf(-1)
	}
	if idx >= len(arr) {
		return math.Inf(1)
	}
	return float64(arr[idx])
}
