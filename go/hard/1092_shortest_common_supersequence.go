// https://leetcode.com/problems/shortest-common-supersequence/
package hard

func ShortestCommonSupersequence(str1 string, str2 string) string {
	lenStr1 := len(str1)
	lenStr2 := len(str2)

	dp := make([][]int, lenStr1+1)
	for i := range dp {
		dp[i] = make([]int, lenStr2+1)
	}

	for row := 0; row <= lenStr1; row++ {
		dp[row][0] = row
	}

	for col := 0; col <= lenStr2; col++ {
		dp[0][col] = col
	}

	for row := 1; row <= lenStr1; row++ {
		for col := 1; col <= lenStr2; col++ {
			if str1[row-1] == str2[col-1] {
				dp[row][col] = dp[row-1][col-1] + 1
			} else {
				dp[row][col] = min(dp[row-1][col], dp[row][col-1]) + 1
			}
		}
	}

	sequence := make([]byte, 0)
	row := lenStr1
	col := lenStr2

	for row > 0 && col > 0 {
		if str1[row-1] == str2[col-1] {
			sequence = append(sequence, str1[row-1])
			row--
			col--
		} else if dp[row-1][col] < dp[row][col-1] {
			sequence = append(sequence, str1[row-1])
			row--
		} else {
			sequence = append(sequence, str2[col-1])
			col--
		}
	}

	for row > 0 {
		sequence = append(sequence, str1[row-1])
		row--
	}
	for col > 0 {
		sequence = append(sequence, str2[col-1])
		col--
	}

	result := make([]byte, 0)
	for i := len(sequence) - 1; i >= 0; i-- {
		result = append(result, sequence[i])
	}

	return string(result)
}
