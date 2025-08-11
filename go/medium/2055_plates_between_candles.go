// https://leetcode.com/problems/plates-between-candles/
package medium

func PlatesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	plates := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			count++
		}
		plates[i] = count
	}

	// prevCandle[i]: index of nearest '|' at or before i, or -1
	prevCandle := make([]int, n)
	prev := -1
	for i := 0; i < n; i++ {
		if s[i] == '|' {
			prev = i
		}
		prevCandle[i] = prev
	}

	// nextCandle[i]: index of nearest '|' at or after i, or -1
	nextCandle := make([]int, n)
	next := -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			next = i
		}
		nextCandle[i] = next
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		// find first candle >= l, last candle <= r
		left := nextCandle[l]
		right := prevCandle[r]
		if left == -1 || right == -1 || left >= right {
			result[i] = 0
		} else {
			result[i] = plates[right] - plates[left]
		}
	}

	return result
}
