// https://leetcode.com/problems/regular-expression-matching/
package hard

func IsMatch(s string, p string) bool {
	memo := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		key := [2]int{i, j}
		if v, ok := memo[key]; ok {
			return v
		}
		var result bool
		if j == len(p) {
			result = (i == len(s))
		} else {
			firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				result = dfs(i, j+2) || (firstMatch && dfs(i+1, j))
			} else {
				result = firstMatch && dfs(i+1, j+1)
			}
		}
		memo[key] = result
		return result
	}
	return dfs(0, 0)
}
