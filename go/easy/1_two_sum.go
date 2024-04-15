// https://leetcode.com/problems/two-sum/
package easy

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}

		seen[num] = i
	}

	return nil
}
