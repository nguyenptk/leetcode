// https://leetcode.com/problems/maximum-length-of-repeated-subarray/
package medium

func FindLength(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	result := 0

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if result < dp[i][j] {
					result = dp[i][j]
				}
			}
		}
	}

	return result
}
