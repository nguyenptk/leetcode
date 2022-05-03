// https://leetcode.com/problems/two-sum/
package easy

func TwoSum(nums []int, target int) []int {
	mapComplement := map[int]int{}
	for k, v := range nums {
		mapComplement[target-v] = k
	}
	for k, v := range nums {
		if val, ok := mapComplement[v]; ok && k != val {
			return []int{k, val}
		}
	}
	return nil
}
