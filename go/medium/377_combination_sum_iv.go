package medium

func CombinationSum4(nums []int, target int) int {
	// return topDownCombinationSum4(nums, target, make(map[int]int))
	return bottomUpCombinationSum4(nums, target)
}

func topDownCombinationSum4(nums []int, target int, memo map[int]int) int {
	if target == 0 {
		return 1
	}
	if _, ok := memo[target]; ok {
		return memo[target]
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		subAmount := target - nums[i]
		if subAmount >= 0 {
			count += topDownCombinationSum4(nums, subAmount, memo)
		}
	}
	memo[target] = count

	return count
}

func bottomUpCombinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for qty := 1; qty <= target; qty++ {
		for _, num := range nums {
			if qty-num >= 0 {
				dp[qty] += dp[qty-num]
			}
		}
	}

	return dp[target]
}
