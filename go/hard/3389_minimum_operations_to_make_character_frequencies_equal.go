// https://leetcode.com/problems/minimum-operations-to-make-character-frequencies-equal/
package hard

// BOTTOM-UP
// func makeStringGood(s string) int {
// 	ans := len(s)
// 	count := make([]int, 26)
// 	for _, c := range s {
// 		count[c-'a']++
// 	}

// 	maxCount := 0
// 	for _, c := range count {
// 		if c > maxCount {
// 			maxCount = c
// 		}
// 	}

// 	for target := 1; target <= maxCount; target++ {
// 		ans = min(ans, getMinOperations(count, target))
// 	}

// 	return ans
// }

// func getMinOperations(count []int, target int) int {
// 	dp := make([]int, 27)

// 	for i := 25; i >= 0; i-- {
// 		// Option 1: Delete all characters of this letter
// 		deleteAll := freq[i]

// 		// Option 2: Adjust count to target by insert/delete
// 		adjustToTarget := abs(targetFreq - freq[i])

// 		dp[i] = min(deleteAll, adjustToTarget) + dp[i+1]

// 		if i+1 < 26 && freq[i+1] < targetFreq {
// 			deficitNext := targetFreq - freq[i+1]

// 			// Number of current chars we can convert
// 			convertable := freq[i]
// 			if freq[i] > targetFreq {
// 				convertable = freq[i] - targetFreq
// 			}

// 			var changeCost int
// 			if deficitNext > convertable {
// 				// Convert all possible, then insert the rest
// 				changeCost = convertable + (deficitNext - convertable)
// 			} else {
// 				// Convert just enough, then delete the rest
// 				changeCost = deficitNext + (convertable - deficitNext)
// 			}

// 			dp[i] = min(dp[i], changeCost+dp[i+2])
// 		}
// 	}

// 	return dp[0]
// }

// TOP-DOWN
func makeStringGood(s string) int {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}

	ans := len(s)
	memo := map[[2]int]int{}

	var dfs func(i, target int) int
	dfs = func(i, target int) int {
		if i >= 26 {
			return 0
		}
		key := [2]int{i, target}
		if val, ok := memo[key]; ok {
			return val
		}

		// Option 1: delete all i-th characters
		deleteAll := freq[i] + dfs(i+1, target)

		// Option 2: insert/delete to reach target
		adjust := abs(freq[i]-target) + dfs(i+1, target)

		best := min(deleteAll, adjust)

		// Option 3: change some i-th to (i+1)-th if it helps
		if i < 25 && freq[i+1] < target {
			deficit := target - freq[i+1]
			var changeCost int
			available := freq[i]
			if freq[i] > target {
				available = freq[i] - target
			}

			if deficit > available {
				changeCost = available + (deficit - available)
			} else {
				changeCost = deficit + (available - deficit)
			}
			best = min(best, changeCost+dfs(i+2, target))
		}

		memo[key] = best
		return best
	}

	for target := 1; target <= maxFreq; target++ {
		memo = map[[2]int]int{}
		ans = min(ans, dfs(0, target))
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
