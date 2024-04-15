// https://leetcode.com/problems/permutations-ii/

package medium

func PermuteUnique(nums []int) [][]int {
	result := [][]int{}
	findPermutations(nums, 0, len(nums), &result)
	return result
}

func shouldSwap(nums []int, start int, curr int) bool {
	for i := start; i < curr; i++ {
		if nums[i] == nums[curr] {
			return false
		}
	}
	return true
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func findPermutations(nums []int, index int, n int, result *[][]int) {
	if index >= n {
		x := make([]int, len(nums))
		copy(x, nums)
		*result = append(*result, x)
		return
	}
	for i := index; i < n; i++ {
		check := shouldSwap(nums, index, i)
		if check {
			swap(nums, index, i)
			findPermutations(nums, index+1, n, result)
			swap(nums, index, i)
		}
	}
}
