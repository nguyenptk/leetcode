// https://leetcode.com/problems/uncrossed-lines/
package medium

func MaxUncrossedLines(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[n1][n2]
}
