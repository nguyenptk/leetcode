// https://leetcode.com/problems/longest-increasing-subsequence/
package medium

func LengthOfLIS(nums []int) int {
	list := []int{}

	for _, num := range nums {
		if len(list) == 0 || num > list[len(list)-1] {
			list = append(list, num)
		} else {
			i := 0
			j := len(list) - 1

			for i < j {
				mid := (i + j) / 2
				if list[mid] < num {
					i = mid + 1
				} else {
					j = mid
				}
			}

			list[j] = num
		}
	}

	return len(list)
}
