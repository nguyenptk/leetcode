// https://leetcode.com/problems/search-insert-position/
package easy

func SearchInsert(nums []int, target int) int {
	n := len(nums)
	m := nums[n/2]
	pos := 0
	// from left to mid
	if target <= m {
		for i := n / 2; i >= 0; i-- {
			if target < nums[i] {
				pos = i - 1
			} else if target == nums[i] {
				pos = i
				break
			} else if target > nums[i] {
				pos = i + 1
				break
			}
		}
		// from mid to right
	} else {
		for i := n / 2; i < n; i++ {
			if target < nums[i] {
				if pos != 0 {
					break
				}
				pos = i - 1
				break
			} else if target == nums[i] {
				pos = i
				break
			} else if target > nums[i] {
				pos = i + 1
			}
		}
	}
	// add to the first position
	if pos < 0 {
		return 0
	}
	return pos
}
