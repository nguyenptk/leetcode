// https://leetcode.com/problems/find-the-duplicate-number
package medium

func FindDuplicate(nums []int) int {
	slow := 0
	fast := 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			return slow
		}
	}

	return slow
}
