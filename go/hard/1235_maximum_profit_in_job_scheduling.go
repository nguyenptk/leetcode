// https://leetcode.com/problems/maximum-profit-in-job-scheduling/
package hard

import (
	"sort"
)

func JobScheduling(startTime []int, endTime []int, profit []int) int {
	// Create a 2D array for jobs
	jobs := make([][]int, len(profit))
	for i := 0; i < len(profit); i++ {
		jobs[i] = []int{startTime[i], endTime[i], profit[i]}
	}

	// Sort jobs by start time
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	// Extract sorted start times for binary search
	start := make([]int, len(profit))
	for i := 0; i < len(profit); i++ {
		start[i] = jobs[i][0]
	}

	// Initialize memoization table
	memo := make([]int, len(profit))
	for i := range memo {
		memo[i] = -1
	}

	// Helper function to find the maximum profit
	var findMaxProfit func(position int) int
	findMaxProfit = func(position int) int {
		// Base case: No more jobs to process
		if position == len(jobs) {
			return 0
		}

		// Return memoized result if it exists
		if memo[position] != -1 {
			return memo[position]
		}

		// Find the next non-conflicting job using binary search
		nextIndex := sort.Search(len(start), func(i int) bool {
			return start[i] >= jobs[position][1]
		})

		// Calculate the maximum profit by either skipping or taking the current job
		maxProfit := max(
			findMaxProfit(position+1),
			jobs[position][2]+findMaxProfit(nextIndex),
		)

		// Store result in memo table
		memo[position] = maxProfit
		return maxProfit
	}

	return findMaxProfit(0)
}
