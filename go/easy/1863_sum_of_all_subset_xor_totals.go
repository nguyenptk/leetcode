// https://leetcode.com/problems/sum-of-all-subset-xor-totals/
package easy

func SubsetXORSum(nums []int) int {
	var xorSum func(idx, curr int) int
	xorSum = func(idx, curr int) int {
		if idx == len(nums) {
			return curr
		}

		withElement := xorSum(idx+1, curr^nums[idx])
		withoutElement := xorSum(idx+1, curr)

		return withElement + withoutElement
	}
	return xorSum(0, 0)
}
