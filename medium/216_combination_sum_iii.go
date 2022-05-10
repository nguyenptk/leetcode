// https://leetcode.com/problems/combination-sum-iii/

package medium

func CombinationSum3(k int, n int) [][]int {
	result := [][]int{}
	combine := []int{}
	combineSum3(k, n, 1, combine, &result)
	return result
}

func combineSum3(k, n, index int, combine []int, result *[][]int) {
	if n == 0 {
		if len(combine) == k {
			x := make([]int, len(combine))
			copy(x, combine)
			*result = append(*result, x)
		}
		return
	}
	for i := index; i < 10; i++ {
		if n >= i {
			combine = append(combine, i)
			combineSum3(k, n-i, i+1, combine, result)
			combine = combine[:len(combine)-1] // Pop
		}
	}
}
