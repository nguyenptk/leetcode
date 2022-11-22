// https://leetcode.com/problems/perfect-squares/
package hard

func NumSquares(n int) int {
	result := make([]int, n+1)
	for i := range result {
		result[i] = n
	}

	result[0] = 0
	result[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if result[i] > result[i-j*j]+1 {
				result[i] = result[i-j*j] + 1
			}
		}
	}

	return result[n]
}
