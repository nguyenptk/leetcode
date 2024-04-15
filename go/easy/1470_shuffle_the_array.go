// https://leetcode.com/problems/shuffle-the-array/
package easy

func Shuffle(nums []int, n int) []int {
	result := make([]int, 2*n)
	for i := 0; i < n; i++ {
		result[2*i] = nums[i]
		result[2*i+1] = nums[n+i]
	}
	return result
}
