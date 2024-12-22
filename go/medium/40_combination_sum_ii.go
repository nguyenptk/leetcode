// https://leetcode.com/problems/combination-sum-ii/
package medium

import "sort"

func CmbinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // O(nlogn)
	result := make([][]int, 0)
	curr := make([]int, 0)

	var backtrack func(index int, currSum int, curr []int)
	backtrack = func(index int, currSum int, curr []int) {
		if currSum == target {
			result = append(result, append([]int{}, curr...))
		}
		if currSum > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			curr = append(curr, candidates[i])
			backtrack(i+1, currSum+candidates[i], curr)
			curr = curr[:len(curr)-1] // POP
		}
	}

	backtrack(0, 0, curr) // O(n^2)
	return result
}
