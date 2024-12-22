// https://leetcode.com/problems/jump-game
package medium

func CanJump(nums []int) bool {
	lastPost := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPost {
			lastPost = i
		}
	}
	return lastPost == 0
}
