// https://leetcode.com/problems/count-number-of-bad-pairs/
package medium

func CountBadPairs(nums []int) int64 {
	result := int64(0)
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		diff := i - nums[i]
		good := m[diff]
		result += int64(i - good)
		m[diff] = good + 1
	}

	return result
}
