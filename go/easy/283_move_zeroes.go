// https://leetcode.com/problems/move-zeroes/
package easy

func MoveZeroes(nums []int) {
	nonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZero] = nums[nonZero], nums[i]
			nonZero++
		}
	}
}
