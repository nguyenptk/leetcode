// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
package medium

func NumPairsDivisibleBy60(time []int) int {
	result := 0
	m := make(map[int]int)

	for _, t := range time {
		remainder := t % 60
		amount := (60 - remainder) % 60
		result += m[amount]
		m[remainder]++
	}

	return result
}
