// https://leetcode.com/problems/candy/
package hard

func Candy(ratings []int) int {
	count := 0
	n := len(ratings)

	// Initialize left array
	l := make([]int, n)
	for i := range l {
		l[i] = 1
	}

	// Initialize right array
	r := make([]int, n)
	for i := range r {
		r[i] = 1
	}

	// Iterate from left
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			l[i] = l[i-1] + 1
		}
	}

	// Iterate from right
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			r[i] = r[i+1] + 1
		}
	}

	// Merge left and right
	for i := 0; i < n; i++ {
		if l[i] >= r[i] {
			count += l[i]
		} else {
			count += r[i]
		}
	}

	return count
}
