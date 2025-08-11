// https://leetcode.com/problems/longest-happy-string/
package medium

func LongestDiverseString(a int, b int, c int) string {
	result := ""
	countA := 0
	countB := 0
	countC := 0

	total := a + b + c

	for i := 0; i < total; i++ {
		if (a >= b && a >= c && countA != 2) ||
			(a > 0 && (countB == 2 || countC == 2)) {
			result += "a"
			a--
			countA++
			countB = 0
			countC = 0
		} else if (b >= a && b >= c && countB != 2) ||
			(b > 0 && (countA == 2 || countC == 2)) {
			result += "b"
			b--
			countB++
			countA = 0
			countC = 0
		} else if (c >= a && c >= b && countC != 2) ||
			(c > 0 && (countA == 2 || countB == 2)) {
			result += "c"
			c--
			countC++
			countA = 0
			countB = 0
		}
	}

	return result
}

// Time complexity: O(a+b+c)
// Space complexity: O(1)
