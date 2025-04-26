// https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/
package easy

func MergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	n1 := len(nums1)
	n2 := len(nums2)
	i, j := 0, 0

	result := make([][]int, 0)
	for i < n1 && j < n2 {
		if nums1[i][0] == nums2[j][0] {
			result = append(result, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		} else if nums1[i][0] < nums2[j][0] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	for i < n1 {
		result = append(result, nums1[i])
		i++
	}

	for j < n2 {
		result = append(result, nums2[j])
		j++
	}

	return result
}
