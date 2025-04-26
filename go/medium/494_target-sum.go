// https://leetcode.com/problems/target-sum/
package medium

func FindTargetSumWays(nums []int, target int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	memo := make(map[[2]int]int)

	var recMemo func(int, int) int
	recMemo = func(currIdx, currSum int) int {
		if currIdx == len(nums) {
			if currSum == target {
				return 1
			}
			return 0
		}

		key := [2]int{currIdx, currSum}
		if val, ok := memo[key]; ok {
			return val
		}

		add := recMemo(currIdx+1, currSum+nums[currIdx])
		subtract := recMemo(currIdx+1, currSum-nums[currIdx])
		memo[key] = add + subtract
		return memo[key]
	}

	return recMemo(0, 0)
}
