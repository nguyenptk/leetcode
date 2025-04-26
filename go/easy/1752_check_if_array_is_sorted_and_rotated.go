// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/
package easy

func Check(nums []int) bool {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			count++
		}
		if count > 1 {
			return false
		}
	}

	return true
}
