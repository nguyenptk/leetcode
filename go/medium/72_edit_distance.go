// https://leetcode.com/problems/edit-distance/
package medium

func EditDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1)+1)
	for i := range memo {
		memo[i] = make([]int, len(word2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(w1Idx, w2Idx int) int
	dfs = func(w1Idx, w2Idx int) int {
		if w1Idx == 0 {
			return w2Idx
		}
		if w2Idx == 0 {
			return w1Idx
		}
		if memo[w1Idx][w2Idx] != -1 {
			return memo[w1Idx][w2Idx]
		}
		minDis := 0
		if word1[w1Idx-1] == word2[w2Idx-1] {
			minDis = dfs(w1Idx-1, w2Idx-1)
		} else {
			insert := dfs(w1Idx, w2Idx-1)
			delete := dfs(w1Idx-1, w2Idx)
			replace := dfs(w1Idx-1, w2Idx-1)
			minDis = min(min(insert, delete), replace) + 1
		}
		memo[w1Idx][w2Idx] = minDis
		return minDis
	}

	return dfs(len(word1), len(word2))
}
