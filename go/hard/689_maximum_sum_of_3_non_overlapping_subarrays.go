// https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/
package hard

func MaxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	kSums := make([]int, 0)
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	kSums = append(kSums, sum)

	for i := k; i < n; i++ {
		sum += nums[i] - nums[i-k]
		kSums = append(kSums, sum)
	}

	memo := make(map[[2]int]int)

	var getMaxSum func(int, int) int
	getMaxSum = func(i, count int) int {
		if count == 3 || i > n-k {
			return 0
		}
		key := [2]int{i, count}
		if val, ok := memo[key]; ok {
			return val
		}

		include := kSums[i] + getMaxSum(i+k, count+1)
		skip := getMaxSum(i+1, count)

		memo[key] = max(include, skip)

		return memo[key]
	}

	var getIndicies func() []int
	getIndicies = func() []int {
		indicies := make([]int, 0)
		i := 0

		for i <= n-k && len(indicies) < 3 {
			include := kSums[i] + getMaxSum(i+k, len(indicies)+1)
			skip := getMaxSum(i+1, len(indicies))

			if include >= skip {
				indicies = append(indicies, i)
				i += k
			} else {
				i++
			}
		}

		return indicies
	}

	return getIndicies()
}
