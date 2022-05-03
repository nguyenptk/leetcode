// https://leetcode.com/problems/binary-search/
package easy

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, l int, r int, target int) int {
	if r >= l {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			return binarySearch(nums, l, mid-1, target)
		}
		return binarySearch(nums, mid+1, r, target)
	}
	return -1
}
