// https://leetcode.com/problems/valid-mountain-array/
package easy

func ValidMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	i := 0
	for i < len(arr)-1 && arr[i] < arr[i+1] {
		i++
	}
	if i == 0 || i == len(arr)-1 {
		return false
	}
	for i < len(arr)-1 && arr[i] > arr[i+1] {
		i++
	}
	if i == len(arr)-1 {
		return true
	}
	return false
}
