// https://leetcode.com/problems/majority-element/
package easy

func MajorityElement(nums []int) int {
	majorityCount := len(nums) / 2

	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > majorityCount {
			return nums[i]
		}
	}

	return -1
}
