// https://leetcode.com/problems/remove-element/
package easy

func RemoveElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		nums[i-count] = nums[i]
		if nums[i] == val {
			count += 1
		}
	}
	return len(nums) - count
}
