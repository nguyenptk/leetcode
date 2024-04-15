// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package medium

func SearchRange(nums []int, target int) []int {
	found := []int{}
	found = append(found, first(nums, target, 0, len(nums)-1))
	found = append(found, last(nums, target, 0, len(nums)-1))
	return found
}

func first(arr []int, target, l, r int) int {
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			res = mid
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

func last(arr []int, target, l, r int) int {
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			res = mid
			l = mid + 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
