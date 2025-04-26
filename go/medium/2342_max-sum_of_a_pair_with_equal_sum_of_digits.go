// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/

package medium

func MaximumSum(nums []int) int {
	result := -1
	m := make(map[int]int, 81)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if val, ok := m[sumDigits(num)]; ok {
			result = max(result, val+num)
		}
		m[sumDigits(num)] = max(m[sumDigits(num)], num)
	}

	return result
}

func sumDigits(num int) int {
	result := 0
	for num > 0 {
		result += num % 10
		num = num / 10
	}

	return result
}
