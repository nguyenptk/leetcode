// https://leetcode.com/problems/combination-sum/
package medium

func CombinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	curr := make([]int, 0)

	var backtrack func(remain int, curr []int, start int)
	backtrack = func(remain int, curr []int, start int) {
		if remain == 0 {
			result = append(result, append([]int{}, curr...))
			return
		} else if remain < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			curr = append(curr, candidates[i])
			backtrack(remain-candidates[i], curr, i)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(target, curr, 0)
	return result
}
