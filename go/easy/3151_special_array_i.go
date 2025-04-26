// https://leetcode.com/problems/special-array-i/
package easy

func IsArraySpecial(nums []int) bool {
	i := 0
	j := 1
	for i < len(nums)-1 {
		if nums[i]%2 == 0 {
			if nums[j]%2 == 0 {
				return false
			}
		} else {
			if nums[j]%2 != 0 {
				return false
			}
		}
		i++
		j++
	}
	return true
}
