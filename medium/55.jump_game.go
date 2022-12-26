// https://leetcode.com/problems/jump-game
package medium

func CanJump(nums []int) bool {
	i := 0

	for reach := 0; i < len(nums) && i <= reach; i++ {
		if reach < i+nums[i] {
			reach = i + nums[i]
		}
	}

	return i == len(nums)
}
