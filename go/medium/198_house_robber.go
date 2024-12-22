package medium

func Rob(nums []int) int {
	return robFrom(nums, 0, make(map[int]int))
}

func robFrom(nums []int, idx int, memo map[int]int) int {
	if idx >= len(nums) {
		return 0
	}

	if _, ok := memo[idx]; ok {
		return memo[idx]
	}

	sum1 := nums[idx] + robFrom(nums, idx+2, memo)
	sum2 := robFrom(nums, idx+1, memo)

	result := max(sum1, sum2)

	memo[idx] = result

	return result
}
