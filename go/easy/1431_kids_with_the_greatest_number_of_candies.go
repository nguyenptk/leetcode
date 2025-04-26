// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
package easy

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxC := 0
	for _, c := range candies {
		maxC = max(maxC, c)
	}

	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxC {
			result[i] = true
		}
	}

	return result
}
