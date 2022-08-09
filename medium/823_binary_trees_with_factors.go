// https://leetcode.com/problems/binary-trees-with-factors/
package medium

import "sort"

func NumFactoredBinaryTrees(arr []int) int {
	// init
	mod := (int)(1e9 + 7)
	n := len(arr)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	// sort the arr
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// setup mapNums to store the nums of arr
	mapNums := map[int]int{}
	for i := 0; i < n; i++ {
		mapNums[arr[i]] = i
	}

	for i := 0; i < n; i++ {
		// arr[i] is root
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 { // arr[j] is left subtree
				right := arr[i] / arr[j]
				if _, ok := mapNums[right]; ok {
					dp[i] += dp[j] * dp[mapNums[right]]
					dp[i] %= mod
				}
			}
		}
	}

	result := 0
	for i := 0; i < len(dp); i++ {
		result += dp[i]
	}

	return result % mod
}
