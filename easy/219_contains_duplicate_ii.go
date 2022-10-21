// https://leetcode.com/problems/contains-duplicate-ii/
package easy

func ContainsNearbyDuplicate(nums []int, k int) bool {
	mapD := map[int]int{}
	for i, v := range nums {
		if key, ok := mapD[v]; ok && i-key <= k {
			return true
		}
		mapD[v] = i
	}
	return false
}
