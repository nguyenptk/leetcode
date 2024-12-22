package medium

func RobII(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	sum1 := robIIFrom(nums, 0, n-2, make(map[int]int))
	sum2 := robIIFrom(nums, 1, n-1, make(map[int]int))

	return max(sum1, sum2)
}

func robIIFrom(nums []int, start, end int, memo map[int]int) int {
	if start > end {
		return 0
	}
	if _, ok := memo[start]; ok {
		return memo[start]
	}

	sum1 := nums[start] + robIIFrom(nums, start+2, end, memo)
	sum2 := robIIFrom(nums, start+1, end, memo)
	maxRob := max(sum1, sum2)

	memo[start] = maxRob

	return maxRob
}
