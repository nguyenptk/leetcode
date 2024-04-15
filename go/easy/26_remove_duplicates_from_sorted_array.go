// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package easy

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return i + 1
}
