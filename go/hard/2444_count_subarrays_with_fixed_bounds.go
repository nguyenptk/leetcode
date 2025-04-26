// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
package hard

func CountSubarrays(nums []int, minK int, maxK int) int64 {
	result := 0

	minIdx, maxIdx, l := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == minK {
			minIdx = i
		}
		if nums[i] == maxK {
			maxIdx = i
		}

		if nums[i] < minK || nums[i] > maxK {
			l = i
		}

		result += max(0, min(minIdx, maxIdx)-l)
	}

	return int64(result)
}
