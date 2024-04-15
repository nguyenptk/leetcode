// https://leetcode.com/problems/contains-duplicate/description/
package easy

func ContainsDuplicate(nums []int) bool {
	exist := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := exist[nums[i]]; ok {
			return true
		}
		exist[nums[i]] = true
	}
	return false
}
