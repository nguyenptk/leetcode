// https://leetcode.com/problems/tuple-with-same-product/
package medium

func TupleSameProduct(nums []int) int {
	n := len(nums)

	freq := map[int]int{}
	total := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			freq[nums[i]*nums[j]]++
		}
	}

	for _, v := range freq {
		pairs := v * (v - 1) / 2
		total += 8 * pairs
	}

	return total
}
