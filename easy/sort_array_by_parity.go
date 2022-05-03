// https://leetcode.com/problems/sort-array-by-parity/
package easy

func SortArrayByParity(nums []int) []int {
	resultEven := []int{}
	resultOdd := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			resultEven = append(resultEven, nums[i])
		} else {
			resultOdd = append(resultOdd, nums[i])
		}
	}
	resultEven = append(resultEven, resultOdd...)

	return resultEven
}
