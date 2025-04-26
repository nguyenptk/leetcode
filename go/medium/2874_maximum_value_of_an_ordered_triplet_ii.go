// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/
package medium

func MaximumTripletValue(nums []int) int64 {
	n := len(nums)
	var result, iMax, dMax int64 = 0, 0, 0
	for k := 0; k < n; k++ {
		result = max(result, dMax*int64(nums[k]))
		dMax = max(dMax, iMax-int64(nums[k]))
		iMax = max(iMax, int64(nums[k]))
	}
	return result
}
