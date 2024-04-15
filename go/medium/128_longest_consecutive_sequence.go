// https://leetcode.com/problems/longest-consecutive-sequence/
package medium

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Put nums to set
	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}

	// init the longest
	longest := 1

	for v := range set {
		// check the left side of the number is not existing
		if _, ok := set[v-1]; !ok {
			// count the right side of number consecutively
			count := 1
			for i := 1; ; i++ {
				if _, ok := set[v+i]; ok {
					count += 1
					if longest < count {
						longest = count
					}
				} else {
					break
				}
			}
		}
	}

	return longest
}
