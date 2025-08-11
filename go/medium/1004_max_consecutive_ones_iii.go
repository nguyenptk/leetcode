// https://leetcode.com/problems/max-consecutive-ones-iii
package medium

func LongestOnes(nums []int, k int) int {
	l := 0
	r := 0
	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			k--
		}
		if k < 0 {
			k += 1 - nums[l]
			l++
		}
	}
	return r - l
}
