// https://leetcode.com/problems/subsets/

package medium

func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	curr := make([]int, 0)

	var backtrack func(n int)
	backtrack = func(n int) {
		result = append(result, append([]int{}, curr...))
		for i := n; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)

	return result
}
