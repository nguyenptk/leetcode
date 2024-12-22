// https://leetcode.com/problems/missing-number/
package easy

func MissingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2 // https://brilliant.org/wiki/sum-of-n-n2-or-n3/
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return total - sum
}
