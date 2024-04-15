// https://leetcode.com/problems/max-number-of-k-sum-pairs/
package medium

func MaxOperations(nums []int, k int) int {
	// init map
	mapK := map[int]int{}
	// init counter
	count := 0
	for _, v := range nums {
		// check map has arr[i]
		if val, ok := mapK[v]; ok && val > 0 {
			mapK[v] = val - 1 // descrease the arr[i] by 1
			count++           // increase the counter
		} else {
			// check map has k - arr[i]
			if _, ok := mapK[k-v]; ok {
				mapK[k-v] += 1 // increase the k - arr[i] by 1
			} else {
				mapK[k-v] = 1 // decrease the k - arr[i] = 1
			}
		}
	}
	return count
}
