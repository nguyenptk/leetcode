// https://leetcode.com/problems/top-k-frequent-elements/
package medium

func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n] = 1 + count[n]
	}

	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	result := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		result = append(result, freq[i]...)
		if len(result) == k {
			return result
		}
	}

	return result
}
