// https://leetcode.com/problems/count-vowels-permutation/
package hard

func CountVowelPermutation(n int) int {
	// To avoid the large output value
	kMod := 1e9 + 7

	// Initialize 2D dp array
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 5)
	}

	// Initialize dp[1][i] as 1 since
	// string of length 1 will consist
	// of only one vowel in the string
	for i := 0; i < 5; i++ {
		dp[1][i] = 1
	}

	// Directed graph using the
	// adjacency matrix
	relation := [][]int{
		{1}, {0, 2},
		{0, 1, 3, 4},
		{2, 4}, {0},
	}

	// Iterate over the range [1, N]
	for i := 1; i < n; i++ {

		// Traverse the directed graph
		for u := 0; u < 5; u++ {
			dp[i+1][u] = 0

			// Traversing the list
			for _, v := range relation[u] {

				// Update dp[i + 1][u]
				dp[i+1][u] += dp[i][v] % int(kMod)
			}
		}
	}

	// Stores total count of permutations
	result := 0

	for i := 0; i < 5; i++ {
		result = (result + dp[n][i]) % int(kMod)
	}

	// Return count of permutations
	return result
}
