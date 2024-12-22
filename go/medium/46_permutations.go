// https://leetcode.com/problems/permutations/
package medium

func Permute(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	curr := make([]int, 0)
	visit := make(map[int]int)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(curr) == n {
			result = append(result, append([]int{}, curr...))
			return
		}

		for i := 0; i < n; i++ {
			if visit[i] == 0 {
				visit[i]++
				curr = append(curr, nums[i])
				backtrack(i + 1)
				curr = curr[:len(curr)-1]
				visit[i]--
			}
		}
	}

	backtrack(0)
	return result
}
