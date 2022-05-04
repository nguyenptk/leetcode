// https://leetcode.com/problems/merge-sorted-array/
package easy

// Add additional return value for unit-test
func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 {
		for k := 0; k < n; k++ {
			nums1[k] = nums2[k]
		}
		return nil
	}
	nums1 = moveToEnd(nums1, m)
	i := n
	j := 0
	for k := 0; k < m+n; k++ {
		if i < m+n && j < n {
			if nums1[i] < nums2[j] {
				nums1[k] = nums1[i]
				i++
			} else {
				nums1[k] = nums2[j]
				j++
			}
		} else if j < n {
			nums1[k] = nums2[j]
			j++
		}
	}
	return nums1
}

func moveToEnd(nums []int, m int) []int {
	k := len(nums)
	for i := m - 1; i >= 0; i-- {
		nums[k-1] = nums[i]
		k--
	}
	return nums
}
