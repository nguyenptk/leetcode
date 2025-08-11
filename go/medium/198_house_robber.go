package medium

// TOP-DOWN
func Rob(nums []int) int {
	memo := map[int]int{}
	var rec func(idx int) int
	rec = func(idx int) int {
		if idx >= len(nums) {
			return 0
		}
		if val, ok := memo[idx]; ok {
			return val
		}
		first := nums[idx] + rec(idx+2)
		second := rec(idx + 1)
		memo[idx] = max(first, second)
		return memo[idx]
	}

	return rec(0)
}

// Time Complexity: O(N)
// Space Complexity: O(N)

// BOTTOM-UP
// func rob(nums []int) int {
//     n := len(nums)
//     dp := make([]int, n+1)
//     dp[0] = 0
//     dp[1] = nums[0]
//     for i := 2; i <= n; i++ {
//         take := dp[i-2] + nums[i-1]
//         skip := dp[i-1]
//         dp[i] = max(take, skip)
//     }

//     return dp[n]
// }
