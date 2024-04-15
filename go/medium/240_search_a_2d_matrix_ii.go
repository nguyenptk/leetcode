// https://leetcode.com/problems/search-a-2d-matrix-ii/
package medium

func SearchMatrixII(matrix [][]int, target int) bool {
	n := len(matrix[0])

	for i := 0; i < len(matrix); i++ {
		if binarySearch(matrix[i], target, 0, n-1) {
			return true
		}
	}

	return false
}

func binarySearch(arr []int, target, l, r int) bool {
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
