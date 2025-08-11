// https://leetcode.com/problems/4sum/
package medium

import "sort"

// Two pointers
// func fourSum(nums []int, target int) [][]int {
//     sort.Ints(nums)
//     return kSum(nums, target, 0, 4)
// }

// func kSum(nums []int, target int, start int, k int) [][]int {
//     result := make([][]int, 0)
//     if start == len(nums) {
//         return result
//     }

//     avg := target/k
//     if nums[start] > avg || avg > nums[len(nums)-1] {
//         return result
//     }

//     if k == 2 {
//         return twoSum(nums, target, start)
//     }

//     for i := start; i < len(nums); i++ {
//         if i == start || nums[i-1] != nums[i] {
//             subsets := kSum(nums, target-nums[i], i+1, k-1)
//             for _, subset := range subsets {
//                 result = append(result, append([]int{nums[i]}, subset...))
//             }
//         }
//     }

//     return result
// }

// func twoSum(nums []int, target int, start int) [][]int {
//     result := make([][]int, 0)
//     l := start
//     r := len(nums)-1
//     for l < r {
//         curr := nums[l] + nums[r]
//         if curr < target || (l > start && nums[l] == nums[l-1]) {
//             l++
//         } else if curr > target || (r < len(nums)-1 && nums[r] == nums[r+1]) {
//             r--
//         } else {
//             result = append(result, []int{nums[l], nums[r]})
//             l++
//             r--
//         }
//     }
//     return result
// }

// Time Complexity: O(n^(k−1)), or O(n^3) for 4Sum. We have k−2 loops, and twoSum is O(n).
// Note that for k>2, sorting the array does not change the overall time complexity.

// Space Complexity: O(n). We need O(k) space for the recursion. k can be the same as n in the worst case for the generalized algorithm.
// Note that, for the purpose of complexity analysis, we ignore the memory required for the output.

// Hashmap
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 0, 4)
}

func kSum(nums []int, target int, start int, k int) [][]int {
	result := make([][]int, 0)
	if start == len(nums) {
		return result
	}
	avg := target / k
	if nums[start] > avg || avg > nums[len(nums)-1] {
		return result
	}
	if k == 2 {
		return twoSum(nums, target, start)
	}
	for i := start; i < len(nums); i++ {
		if i == start || nums[i-1] != nums[i] {
			subsets := kSum(nums, target-nums[i], i+1, k-1)
			for _, subset := range subsets {
				result = append(result, append([]int{nums[i]}, subset...))
			}
		}
	}
	return result
}

func twoSum(nums []int, target int, start int) [][]int {
	result := make([][]int, 0)
	s := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if len(result) == 0 || result[len(result)-1][1] != nums[i] {
			if s[target-nums[i]] {
				result = append(result, []int{target - nums[i], nums[i]})
			}
		}
		s[nums[i]] = true
	}
	return result
}
