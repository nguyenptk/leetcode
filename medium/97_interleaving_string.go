// https://leetcode.com/problems/97/
package medium

func IsInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1) + 2
	m := len(s2) + 2
	if n+m-4 != len(s3) {
		return false
	}
	var dp [102]bool
	dp[1] = true
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			up := dp[j] && (i < 2 || s1[i-2] == s3[j+i-3])
			left := dp[j-1] && (j < 2 || s2[j-2] == s3[j+i-3])
			dp[j] = up || left
		}
	}
	return dp[m-1]
}
