// https://leetcode.com/problems/maximum-frequency-after-subarray-operation/
package medium

func MaxFrequency(nums []int, k int) int {
	freqK := 0
	for _, num := range nums {
		if num == k {
			freqK++
		}
	}

	result := freqK
	for i := 1; i <= 50; i++ {
		if i == k {
			continue
		}
		curr, maxSub := 0, 0
		for _, num := range nums {
			delta := 0
			if num == i {
				delta = 1
			} else if num == k {
				delta = -1
			}
			curr += delta
			if curr < 0 {
				curr = 0
			}
			maxSub = max(maxSub, curr)
		}
		result = max(result, freqK+maxSub)
	}

	return result
}

// Time complexity: O(N)
// Space complexity: O(1)
