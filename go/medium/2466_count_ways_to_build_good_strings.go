// https://leetcode.com/problems/count-ways-to-build-good-strings/
package medium

func CountGoodStrings(low int, high int, zero int, one int) int {
	mod := int(1e9 + 7)

	memo := make(map[int]int)
	var dfs func(int) int
	dfs = func(length int) int {
		if length > high {
			return 0
		}
		if val, ok := memo[length]; ok {
			return val
		}

		var result int
		if length >= low {
			result = 1
		} else {
			result = 0
		}

		result += dfs(length+zero) + dfs(length+one)
		memo[length] = result

		return memo[length] % mod
	}

	return dfs(0)
}
