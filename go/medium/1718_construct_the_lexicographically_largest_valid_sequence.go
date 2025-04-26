// https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/
package medium

func ConstructDistancedSequence(n int) []int {
	if n == 1 {
		return []int{1}
	}

	size := 2*n - 1
	result := make([]int, size)
	used := make([]bool, n+1)

	var backtrack func(pos int) bool
	backtrack = func(pos int) bool {
		if pos == size {
			return true
		}

		if result[pos] != 0 {
			return backtrack(pos + 1)
		}

		for i := n; i >= 1; i-- {
			if used[i] {
				continue
			}

			if i == 1 {
				result[pos] = 1
				used[i] = true
				if backtrack(i + 1) {
					return true
				}
				result[pos] = 0
				used[i] = false
			} else {
				if pos+i < size && result[pos+i] == 0 {
					result[pos] = i
					result[pos+i] = i
					used[i] = true
					if backtrack(pos + 1) {
						return true
					}
					result[pos] = 0
					result[pos+i] = 0
					used[i] = false
				}
			}
		}

		return false
	}

	backtrack(0)
	return result
}
