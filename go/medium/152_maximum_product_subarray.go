package medium

func MaxProductSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	currMax := nums[0]
	currMin := nums[0]
	result := currMax

	for i := 1; i < n; i++ {
		curr := nums[i]
		tmpMax := max(curr, max(currMax*curr, currMin*curr))
		currMin = min(curr, min(currMax*curr, currMin*curr))
		currMax = tmpMax

		result = max(result, currMax)
	}

	return result
}
